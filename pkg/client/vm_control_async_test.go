package client

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

func TestVMControlSingleSubmission(t *testing.T) {
	const vm = "11111111111111111111111111111111"
	const job = "22222222222222222222222222222222"
	for _, tc := range []struct {
		action, method, path, body string
	}{
		{"stop", "PUT", "/zstack/v1/vm-instances/" + vm + "/actions", `{"stopVmInstance":{"type":"grace"}}`},
		{"force-stop", "PUT", "/zstack/v1/vm-instances/" + vm + "/actions", `{"stopVmInstance":{"type":"cold"}}`},
		{"destroy", "DELETE", "/zstack/v1/vm-instances/" + vm + "?deleteMode=Enforcing", ""},
		{"expunge", "PUT", "/zstack/v1/vm-instances/" + vm + "/actions", `{"expungeVmInstance":{}}`},
	} {
		t.Run(tc.action, func(t *testing.T) {
			calls := 0
			status, response := 202, `{"location":"http://internal:8080/zstack/v1/api-jobs/`+job+`"}`
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				calls++
				body, err := io.ReadAll(r.Body)
				if err != nil || string(body) != tc.body || r.Method != tc.method || r.URL.RequestURI() != tc.path || r.Header.Get("X-Job-UUID") != job || r.Header.Get("Authorization") == "" {
					t.Error("incorrect signed control request")
				}
				w.Header().Set("Location", "/must-not-follow")
				w.WriteHeader(status)
				fmt.Fprint(w, response)
			}))
			defer server.Close()
			host, rawPort, _ := net.SplitHostPort(strings.TrimPrefix(server.URL, "http://"))
			port, _ := strconv.Atoi(rawPort)
			c := &ZSClient{ZSHttpClient: &ZSHttpClient{ZSConfig: NewZSConfig(host, port, "zstack").AccessKey("ak", "sk").RetryTimes(100), httpClient: server.Client()}}
			out, err := c.ControlVmInstanceAsync(context.Background(), job, vm, tc.action)
			if err != nil || out.State != "accepted" || out.JobID != job || calls != 1 {
				t.Fatal("submission must neither retry nor poll", out, err, calls)
			}
			status, response = 200, `{}`
			out, err = c.ControlVmInstanceAsync(context.Background(), job, vm, tc.action)
			wantVoid := tc.action == "destroy" || tc.action == "expunge"
			if (err == nil) != wantVoid || wantVoid && out.State != "succeeded" || calls != 2 {
				t.Fatal("action-scoped synchronous void receipt", out, err, calls)
			}
			status, response = 503, `{"error":{"code":"TEST.FAILURE"}}`
			out, err = c.ControlVmInstanceAsync(context.Background(), job, vm, tc.action)
			if err != nil || out.State != "failed" || calls != 3 {
				t.Fatal("Cloud failure must not retry", out, err, calls)
			}
			status, response = 307, ""
			if _, err = c.ControlVmInstanceAsync(context.Background(), job, vm, tc.action); err == nil || calls != 4 {
				t.Fatal("redirect followed or accepted", err, calls)
			}
			c.readOnly = true
			if _, err = c.ControlVmInstanceAsync(context.Background(), job, vm, tc.action); err == nil || calls != 4 {
				t.Fatal("read-only mutation", err, calls)
			}
		})
	}
}

func TestVMControlJobStates(t *testing.T) {
	const vm = "11111111111111111111111111111111"
	const job = "22222222222222222222222222222222"
	for _, tc := range []struct {
		status      int
		body, state string
		wantErr     error
	}{
		{202, "", "running", nil},
		{202, `{}`, "running", nil},
		{200, `{"inventory":{"uuid":"` + vm + `"}}`, "succeeded", nil},
		{200, `{"inventory":{"uuid":"` + job + `"}}`, "", ErrAsyncResponse},
		{200, `{"success":false}`, "", ErrAsyncResponse},
		{503, `{"error":{"code":"TEST.FAILURE"}}`, "failed", nil},
		{503, `{}`, "", ErrAsyncResponse},
		{404, "", "", ErrAsyncNotFound},
		{202, `{"location":"http://internal:8080/zstack/v1/api-jobs/` + vm + `"}`, "", ErrAsyncResponse},
	} {
		t.Run(fmt.Sprint(tc.status)+"/"+tc.body, func(t *testing.T) {
			calls := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				calls++
				w.WriteHeader(tc.status)
				fmt.Fprint(w, tc.body)
			}))
			defer server.Close()
			host, rawPort, _ := net.SplitHostPort(strings.TrimPrefix(server.URL, "http://"))
			port, _ := strconv.Atoi(rawPort)
			c := &ZSClient{ZSHttpClient: &ZSHttpClient{ZSConfig: NewZSConfig(host, port, "zstack").AccessKey("ak", "sk").RetryTimes(100), httpClient: server.Client()}}
			out, err := c.QueryVmControlForAction(context.Background(), job, vm, "stop")
			if err != tc.wantErr || out.State != tc.state || calls != 1 {
				t.Fatal("unexpected job observation", out, err, calls)
			}
		})
	}
}

func TestVMControlVoidReceiptsRequireOriginalAction(t *testing.T) {
	const vm = "11111111111111111111111111111111"
	const job = "22222222222222222222222222222222"
	for _, body := range []string{`{}`, ``, `null`, `{"unrelated":true}`, `{"error":{}}`} {
		for _, action := range []string{"stop", "force-stop", "destroy", "expunge", "invalid"} {
			t.Run(action+"/"+body, func(t *testing.T) {
				calls := 0
				server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					calls++
					if r.Method != "GET" || r.URL.Path != "/zstack/v1/api-jobs/"+job || r.Header.Get("Authorization") == "" {
						t.Error("not one signed original job query")
					}
					fmt.Fprint(w, body)
				}))
				defer server.Close()
				host, rawPort, _ := net.SplitHostPort(strings.TrimPrefix(server.URL, "http://"))
				port, _ := strconv.Atoi(rawPort)
				c := &ZSClient{ZSHttpClient: &ZSHttpClient{ZSConfig: NewZSConfig(host, port, "zstack").AccessKey("ak", "sk").RetryTimes(100), httpClient: server.Client()}}
				out, e := c.QueryVmControlForAction(context.Background(), job, vm, action)
				want := body == `{}` && (action == "destroy" || action == "expunge")
				if (e == nil) != want || want && out.State != "succeeded" {
					t.Fatal("action-scoped acknowledgement", out, e)
				}
				n := 1
				if action == "invalid" {
					n = 0
				}
				if calls != n {
					t.Fatal("unexpected retries", calls)
				}
				if _, e = c.QueryVmControl(context.Background(), job, vm); e == nil {
					t.Fatal("unscoped job query accepted void receipt")
				}
			})
		}
	}
}
