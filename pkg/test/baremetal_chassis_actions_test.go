// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryBaremetalChassis(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryBaremetalChassis(&queryParam)
	if err != nil {
		t.Errorf("TestQueryBaremetalChassis error: %v", err)
		return
	}
	golog.Infof("QueryBaremetalChassis result count: %d", len(result))
}

func TestUpdateBaremetalChassis(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBaremetalChassis(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateBaremetalChassis Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BaremetalChassis found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateBaremetalChassisParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateBaremetalChassisParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateBaremetalChassis(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateBaremetalChassis error: %v", err)
		return
	}
	golog.Infof("UpdateBaremetalChassis result: %s", result.UUID)
}

func TestDeleteBaremetalChassis(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteBaremetalChassis is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryBaremetalChassis(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteBaremetalChassis Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No BaremetalChassis found to test Delete")
		return
	}

	err = accountLoginCli.DeleteBaremetalChassis(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteBaremetalChassis error: %v", err)
		return
	}
	golog.Infof("DeleteBaremetalChassis succeeded for UUID: %s", list[0].UUID)
}

func TestCreateBaremetalChassis(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateBaremetalChassis is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateBaremetalChassisParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateBaremetalChassisParamDetail{
	// 		Name: "test-baremetalchassis",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateBaremetalChassis(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateBaremetalChassis error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateBaremetalChassis result: %s", result.Uuid)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteBaremetalChassis(result.Uuid, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteBaremetalChassis error: %v", err)
	// }
}

func TestInspectBaremetalChassis(t *testing.T) {
	// InspectBaremetalChassis operation
	t.Skip("TestInspectBaremetalChassis requires manual implementation")

}
