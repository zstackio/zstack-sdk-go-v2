package client

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

func TestAuditQueriesDecodeErrorSummary(t *testing.T) {
	for _, paged := range []bool{false, true} {
		t.Run(fmt.Sprintf("paged=%t", paged), func(t *testing.T) {
			cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodGet || r.URL.Path != "/zstack/v1/zwatch/audit-records" {
					t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
				}
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprint(w, `{"total":3,"inventories":[
					{"id":1,"success":false,"error":"VM operation failed","responseDump":"{\"error\":{\"code\":\"SYS.1006\"}}"},
					{"id":2,"success":true},
					{"id":3,"success":true,"error":""}
				]}`)
			})
			q := param.NewQueryParam()
			var rows []view.AuditsInventoryView
			var err error
			if paged {
				var total int
				rows, total, err = cli.PageAudit(&q)
				if err == nil && total != 3 {
					t.Fatalf("expected total=3, got %d", total)
				}
			} else {
				rows, err = cli.QueryAudit(&q)
			}
			if err != nil || len(rows) != 3 {
				t.Fatalf("audit query: rows=%+v err=%v", rows, err)
			}
			if rows[0].Error != "VM operation failed" || rows[0].Success || rows[1].Error != "" || rows[2].Error != "" || !rows[1].Success || !rows[2].Success {
				t.Fatalf("unexpected audit fields: %+v", rows)
			}
			if rows[0].ResponseDump != `{"error":{"code":"SYS.1006"}}` {
				t.Fatalf("responseDump changed: %s", rows[0].ResponseDump)
			}

		})
	}
}
