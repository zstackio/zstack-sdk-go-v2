package client

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestVMAsyncSingleRequests(t *testing.T) {
	const id = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	calls := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		calls++
		if r.Header.Get("Authorization") == "" {
			t.Error("unsigned request")
		}
		if r.Method == http.MethodPost {
			w.WriteHeader(202)
			fmt.Fprint(w, `{"location":"http://internal:8080/zstack/v1/api-jobs/`+id+`"}`)
			return
		}
		if r.URL.Path != "/zstack/v1/api-jobs/"+id {
			t.Error("wrong query target")
		}
		w.WriteHeader(202)
		fmt.Fprint(w, `{}`)
	}))
	defer server.Close()
	host, rawPort, _ := net.SplitHostPort(strings.TrimPrefix(server.URL, "http://"))
	port, _ := strconv.Atoi(rawPort)
	c := &ZSClient{ZSHttpClient: &ZSHttpClient{ZSConfig: NewZSConfig(host, port, "zstack").AccessKey("ak", "sk").RetryTimes(100), httpClient: server.Client()}}
	out, err := c.CreateVmInstanceAsync(context.Background(), param.CreateVmInstanceParam{})
	if err != nil || out.State != "accepted" || out.JobID != id || calls != 1 {
		t.Fatal("submission must not wait", out, err, calls)
	}
	out, err = c.QueryVmCreation(context.Background(), id)
	if err != nil || out.State != "running" || calls != 2 {
		t.Fatal("query must not loop", out, err, calls)
	}
	c.readOnly = true
	if _, err = c.CreateVmInstanceAsync(context.Background(), param.CreateVmInstanceParam{}); err == nil || calls != 2 {
		t.Fatal("read-only mutation")
	}
	for _, location := range []string{"http://evil/zstack/v1/api-jobs/" + id + "?secret=x", "http://user:password@evil/zstack/v1/api-jobs/" + id, "/zstack/v1/api-jobs/../licenses", "http://evil/zstack/v1/api-jobs/%61" + id[1:], "//evil/zstack/v1/api-jobs/" + id} {
		if _, err = c.asyncJobID(location); err == nil {
			t.Fatal("unsafe correlation accepted")
		}
	}
}

func TestVMCreationAcknowledgementCloudDates(t *testing.T) {
	const id = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
	for _, submit := range []bool{false, true} {
		t.Run(fmt.Sprint(submit), func(t *testing.T) {
			calls := 0
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				calls++
				fmt.Fprint(w, `{"inventory":{"uuid":"`+id+`","name":"synthetic","state":"Running","createDate":"2026-10-18 08:23:29","lastOpDate":"2026-10-18 08:23:31","allVolumes":[{"uuid":"bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb","createDate":"2026-10-18 08:23:29"}]}}`)
			}))
			defer server.Close()
			host, rawPort, _ := net.SplitHostPort(strings.TrimPrefix(server.URL, "http://"))
			port, _ := strconv.Atoi(rawPort)
			c := &ZSClient{ZSHttpClient: &ZSHttpClient{ZSConfig: NewZSConfig(host, port, "zstack").AccessKey("ak", "sk").RetryTimes(100), httpClient: server.Client()}}
			var out VMAsyncResult
			var err error
			if submit {
				out, err = c.CreateVmInstanceAsync(context.Background(), param.CreateVmInstanceParam{})
			} else {
				out, err = c.QueryVmCreation(context.Background(), id)
			}
			if err != nil || out.State != "succeeded" || out.Inventory == nil || out.Inventory.UUID != id || out.Inventory.CreateDate.IsZero() || len(out.Inventory.AllVolumes) != 1 || out.Inventory.AllVolumes[0].CreateDate.IsZero() || calls != 1 {
				t.Fatalf("Cloud date acknowledgement: state=%s err=%v calls=%d", out.State, err, calls)
			}
		})
	}
}
