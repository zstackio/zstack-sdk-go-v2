// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryCertificate(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryCertificate(&queryParam)
	if err != nil {
		t.Errorf("TestQueryCertificate error: %v", err)
		return
	}
	golog.Infof("QueryCertificate result count: %d", len(result))
}
func TestGetCertificate(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryCertificate(&queryParam)
	if err != nil {
		t.Errorf("TestGetCertificate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Certificate found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetCertificate(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetCertificate error: %v", err)
		return
	}
	golog.Infof("GetCertificate result: %s", result.UUID)
}

func TestUpdateCertificate(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryCertificate(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateCertificate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Certificate found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateCertificateParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateCertificateParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateCertificate(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateCertificate error: %v", err)
		return
	}
	golog.Infof("UpdateCertificate result: %s", result.UUID)
}

func TestDeleteCertificate(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteCertificate is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryCertificate(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteCertificate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Certificate found to test Delete")
		return
	}

	err = accountLoginCli.DeleteCertificate(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteCertificate error: %v", err)
		return
	}
	golog.Infof("DeleteCertificate succeeded for UUID: %s", list[0].UUID)
}

func TestCreateCertificate(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateCertificate is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateCertificateParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateCertificateParamDetail{
	// 		Name: "test-certificate",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateCertificate(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateCertificate error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateCertificate result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteCertificate(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteCertificate error: %v", err)
	// }
}
