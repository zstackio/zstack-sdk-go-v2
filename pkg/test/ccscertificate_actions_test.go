// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryCCSCertificate(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryCCSCertificate(&queryParam)
	if err != nil {
		t.Errorf("TestQueryCCSCertificate error: %v", err)
		return
	}
	golog.Infof("QueryCCSCertificate result count: %d", len(result))
}

func TestDeleteCCSCertificate(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteCCSCertificate is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryCCSCertificate(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteCCSCertificate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No CCSCertificate found to test Delete")
		return
	}

	err = accountLoginCli.DeleteCCSCertificate(list[0].Uuid, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteCCSCertificate error: %v", err)
		return
	}
	golog.Infof("DeleteCCSCertificate succeeded for UUID: %s", list[0].Uuid)
}

func TestAddCCSCertificate(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddCCSCertificate requires valid creation parameters")

}
