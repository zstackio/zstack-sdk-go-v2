// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBareMetal2ProvisionNetwork(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryBareMetal2ProvisionNetwork(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBareMetal2ProvisionNetwork error: %v", err)
		return
	}
	golog.Infof("QueryBareMetal2ProvisionNetwork result count: %d", len(result))
}
func TestGetBareMetal2ProvisionNetwork(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBareMetal2ProvisionNetwork(&queryParam)
	if err != nil {
		t.Errorf("TestGetBareMetal2ProvisionNetwork Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BareMetal2ProvisionNetwork found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetBareMetal2ProvisionNetwork(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetBareMetal2ProvisionNetwork error: %v", err)
		return
	}
	golog.Infof("GetBareMetal2ProvisionNetwork result: %s", result.UUID)
}

func TestUpdateBareMetal2ProvisionNetwork(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBareMetal2ProvisionNetwork(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateBareMetal2ProvisionNetwork Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BareMetal2ProvisionNetwork found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateBareMetal2ProvisionNetworkParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateBareMetal2ProvisionNetworkParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateBareMetal2ProvisionNetwork(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateBareMetal2ProvisionNetwork error: %v", err)
		return
	}
	golog.Infof("UpdateBareMetal2ProvisionNetwork result: %s", result.UUID)
}

func TestDeleteBareMetal2ProvisionNetwork(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteBareMetal2ProvisionNetwork is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBareMetal2ProvisionNetwork(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteBareMetal2ProvisionNetwork Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BareMetal2ProvisionNetwork found to test Delete")
		return
	}

	err = accountLoginCli.DeleteBareMetal2ProvisionNetwork(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteBareMetal2ProvisionNetwork error: %v", err)
		return
	}
	golog.Infof("DeleteBareMetal2ProvisionNetwork succeeded for UUID: %s", list[0].UUID)
}

func TestCreateBareMetal2ProvisionNetwork(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateBareMetal2ProvisionNetwork is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateBareMetal2ProvisionNetworkParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateBareMetal2ProvisionNetworkParamDetail{
	// 		Name: "test-baremetal2provisionnetwork",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateBareMetal2ProvisionNetwork(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateBareMetal2ProvisionNetwork error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateBareMetal2ProvisionNetwork result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteBareMetal2ProvisionNetwork(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteBareMetal2ProvisionNetwork error: %v", err)
	// }
}
