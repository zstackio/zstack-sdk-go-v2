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
func TestGetCCSCertificate(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryCCSCertificate(&queryParam)
	if err != nil {
		t.Errorf("TestGetCCSCertificate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No CCSCertificate found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetCCSCertificate(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetCCSCertificate error: %v", err)
		return
	}
	golog.Infof("GetCCSCertificate result: %s", result.UUID)
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

	err = accountLoginCli.DeleteCCSCertificate(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteCCSCertificate error: %v", err)
		return
	}
	golog.Infof("DeleteCCSCertificate succeeded for UUID: %s", list[0].UUID)
}

func TestAddCCSCertificate(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddCCSCertificate requires valid creation parameters")

}
