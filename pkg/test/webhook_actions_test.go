// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryWebhook(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryWebhook(&queryParam)
	if err != nil {
		t.Errorf("TestQueryWebhook error: %v", err)
		return
	}
	golog.Infof("QueryWebhook result count: %d", len(result))
}
func TestGetWebhook(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryWebhook(&queryParam)
	if err != nil {
		t.Errorf("TestGetWebhook Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Webhook found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetWebhook(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetWebhook error: %v", err)
		return
	}
	golog.Infof("GetWebhook result: %s", result.UUID)
}

func TestUpdateWebhook(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryWebhook(&queryParam)
	if err != nil {
		t.Errorf("TestUpdateWebhook Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Webhook found to test Update")
		return
	}

	// Update with minimal params
	updateParam := param.UpdateWebhookParam{
		BaseParam: param.BaseParam{},
		Params:    param.UpdateWebhookParamDetail{
			// Keep original values - just testing the API works
		},
	}
	result, err := accountLoginCli.UpdateWebhook(list[0].UUID, updateParam)
	if err != nil {
		t.Errorf("TestUpdateWebhook error: %v", err)
		return
	}
	golog.Infof("UpdateWebhook result: %s", result.UUID)
}

func TestDeleteWebhook(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteWebhook is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryWebhook(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteWebhook Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Webhook found to test Delete")
		return
	}

	err = accountLoginCli.DeleteWebhook(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteWebhook error: %v", err)
		return
	}
	golog.Infof("DeleteWebhook succeeded for UUID: %s", list[0].UUID)
}

func TestCreateWebhook(t *testing.T) {
	// WARNING: This test will create a real resource!
	t.Skip("TestCreateWebhook is skipped by default. Implement with valid params to test creation.")

	// createParam := param.CreateWebhookParam{
	// 	BaseParam: param.BaseParam{},
	// 	Params: param.CreateWebhookParamDetail{
	// 		Name: "test-webhook",
	// 		// Add other required fields
	// 	},
	// }
	// result, err := accountLoginCli.CreateWebhook(createParam)
	// if err != nil {
	// 	t.Errorf("TestCreateWebhook error: %v", err)
	// 	return
	// }
	// golog.Infof("CreateWebhook result: %s", result.UUID)
	//
	// // Cleanup: delete the created resource
	// err = accountLoginCli.DeleteWebhook(result.UUID, param.DeleteModePermissive)
	// if err != nil {
	// 	t.Logf("Cleanup DeleteWebhook error: %v", err)
	// }
}
