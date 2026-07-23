package client

import (
	"net/http"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestDetachDataVolumeFromVm_RequestContract(t *testing.T) {
	vmUUID := "vm-1"
	cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
		assertDataVolumeDetachRequest(t, r, "/zstack/v1/volumes/volume-1/vm-instances")

		query := r.URL.Query()
		if got := query.Get("vmUuid"); got != vmUUID {
			t.Fatalf("unexpected vmUuid: got %q, want %q", got, vmUUID)
		}

		body := decodeRequestBody(t, r)
		assertJSONField(t, body, "vmUuid", vmUUID)
		assertJSONStringArrayField(t, body, "systemTags", "system-tag-1", "system-tag-2")
		assertJSONStringArrayField(t, body, "userTags", "user-tag")
		assertJSONField(t, body, "requestIp", "172.25.16.10")
		if _, found := body["detachDataVolumeFromVm"]; found {
			t.Fatalf("request must not contain a detachDataVolumeFromVm wrapper: %#v", body)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"inventory":{"uuid":"volume-1","vmInstanceUuid":""}}`))
	})

	volume, err := cli.DetachDataVolumeFromVm("volume-1", param.DetachDataVolumeFromVmParam{
		BaseParam: param.BaseParam{
			SystemTags: []string{"system-tag-1", "system-tag-2"},
			UserTags:   []string{"user-tag"},
			RequestIp:  "172.25.16.10",
		},
		Params: param.DetachDataVolumeFromVmParamDetail{
			VmUuid: &vmUUID,
		},
	})
	if err != nil {
		t.Fatalf("DetachDataVolumeFromVm returned error: %v", err)
	}
	if volume.UUID != "volume-1" || volume.VmInstanceUuid != "" {
		t.Fatalf("unexpected detached volume: %+v", volume)
	}
}

func TestDetachDataVolumeFromVm_VmUuidOptional(t *testing.T) {
	cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
		assertDataVolumeDetachRequest(t, r, "/zstack/v1/volumes/volume-1/vm-instances")
		if _, found := r.URL.Query()["vmUuid"]; found {
			t.Fatalf("vmUuid must be omitted when unset: %s", r.URL.RawQuery)
		}
		body := decodeRequestBody(t, r)
		if _, found := body["vmUuid"]; found {
			t.Fatalf("vmUuid must be omitted from the body when unset: %#v", body)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"inventory":{"uuid":"volume-1"}}`))
	})

	_, err := cli.DetachDataVolumeFromVm("volume-1", param.DetachDataVolumeFromVmParam{})
	if err != nil {
		t.Fatalf("DetachDataVolumeFromVm returned error: %v", err)
	}
}

func TestDetachDataVolumeFromHost_RequestContract(t *testing.T) {
	hostUUID := "host-1"
	cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
		assertDataVolumeDetachRequest(t, r, "/zstack/v1/volumes/volume-1/hosts")
		if got := r.URL.Query().Get("hostUuid"); got != "" {
			t.Fatalf("hostUuid must not be sent as a query parameter, got %q", got)
		}
		body := decodeRequestBody(t, r)
		assertJSONField(t, body, "hostUuid", hostUUID)
		if _, found := body["detachDataVolumeFromHost"]; found {
			t.Fatalf("request must not contain a detachDataVolumeFromHost wrapper: %#v", body)
		}

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"success":true}`))
	})

	err := cli.DetachDataVolumeFromHost("volume-1", param.DetachDataVolumeFromHostParam{
		Params: param.DetachDataVolumeFromHostParamDetail{
			HostUuid: &hostUUID,
		},
	})
	if err != nil {
		t.Fatalf("DetachDataVolumeFromHost returned error: %v", err)
	}
}

func assertDataVolumeDetachRequest(t *testing.T, r *http.Request, wantPath string) {
	t.Helper()
	if r.Method != http.MethodDelete {
		t.Fatalf("expected DELETE request, got %s", r.Method)
	}
	if r.URL.Path != wantPath {
		t.Fatalf("unexpected request path: got %s, want %s", r.URL.Path, wantPath)
	}
	if _, found := r.URL.Query()["deleteMode"]; found {
		t.Fatalf("detach request must not contain deleteMode: %s", r.URL.RawQuery)
	}
}
