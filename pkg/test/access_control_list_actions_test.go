// Copyright (c) ZStack.io, Inc.

package test

import (
    "testing"

    "github.com/kataras/golog"
    "github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAccessControlList(t *testing.T) {
    queryParam := param.NewQueryParam()
    result, err := accountLoginCli.QueryAccessControlList(&queryParam)
    if err != nil {
        t.Errorf("TestQueryAccessControlList error: %v", err)
        return
    }
    golog.Infof("QueryAccessControlList result count: %d", len(result))
}

func TestGetAccessControlList(t *testing.T) {
    // First query to get a valid UUID
    queryParam := param.NewQueryParam()
    queryParam.Limit(1)
    list, err := accountLoginCli.QueryAccessControlList(&queryParam)
    if err != nil {
        t.Errorf("TestGetAccessControlList Query error: %v", err)
        return
    }
    if len(list) == 0 {
        t.Skip("No AccessControlList found to test Get")
        return
    }

    // Get by UUID
    result, err := accountLoginCli.GetAccessControlList(list[0].UUID)
    if err != nil {
        t.Errorf("TestGetAccessControlList error: %v", err)
        return
    }
    golog.Infof("GetAccessControlList result: %s", result.UUID)
}

func TestPageAccessControlList(t *testing.T) {
    queryParam := param.NewQueryParam()
    queryParam.Limit(10)
    result, total, err := accountLoginCli.PageAccessControlList(&queryParam)
    if err != nil {
        t.Errorf("TestPageAccessControlList error: %v", err)
        return
    }
    golog.Infof("PageAccessControlList result count: %d, total: %d", len(result), total)
}

func TestUpdateAccessControlList(t *testing.T) {
    // First query to get a valid UUID
    queryParam := param.NewQueryParam()
    queryParam.Limit(1)
    list, err := accountLoginCli.QueryAccessControlList(&queryParam)
    if err != nil {
        t.Errorf("TestUpdateAccessControlList Query error: %v", err)
        return
    }
    if len(list) == 0 {
        t.Skip("No AccessControlList found to test Update")
        return
    }

    // Update with minimal params - 字段名改为操作名称
    updateParam := param.UpdateAccessControlListParam{
        BaseParam: param.BaseParam{},
        UpdateAccessControlList: param.UpdateAccessControlListParamDetail{
            // Add fields you want to update here
            // Name: "updated-name",
            // Description: "updated description",
        },
    }
    result, err := accountLoginCli.UpdateAccessControlList(list[0].UUID, updateParam)
    if err != nil {
        t.Errorf("TestUpdateAccessControlList error: %v", err)
        return
    }
    golog.Infof("UpdateAccessControlList result: %s", result.UUID)
}

func TestDeleteAccessControlList(t *testing.T) {
    // WARNING: This test will actually delete a resource!
    // Query first to get UUID (but skip by default to avoid accidental deletion)
    t.Skip("TestDeleteAccessControlList is skipped by default to prevent accidental deletion. Remove this line to enable.")

    queryParam := param.NewQueryParam()
    queryParam.Limit(1)
    list, err := accountLoginCli.QueryAccessControlList(&queryParam)
    if err != nil {
        t.Errorf("TestDeleteAccessControlList Query error: %v", err)
        return
    }
    if len(list) == 0 {
        t.Skip("No AccessControlList found to test Delete")
        return
    }

    err = accountLoginCli.DeleteAccessControlList(list[0].UUID, param.DeleteModePermissive)
    if err != nil {
        t.Errorf("TestDeleteAccessControlList error: %v", err)
        return
    }
    golog.Infof("DeleteAccessControlList succeeded for UUID: %s", list[0].UUID)
}

func TestCreateAccessControlList(t *testing.T) {
    // WARNING: This test will create a real resource!
    t.Skip("TestCreateAccessControlList is skipped by default. Implement with valid params to test creation.")

    // createParam := param.CreateAccessControlListParam{
    // 	BaseParam: param.BaseParam{},
    // 	CreateAccessControlList: param.CreateAccessControlListParamDetail{  // 字段名改为 CreateAccessControlList
    // 		Name: "test-accesscontrollist",
    // 		// Add other required fields
    // 	},
    // }
    // result, err := accountLoginCli.CreateAccessControlList(createParam)
    // if err != nil {
    // 	t.Errorf("TestCreateAccessControlList error: %v", err)
    // 	return
    // }
    // golog.Infof("CreateAccessControlList result: %s", result.UUID)
    //
    // // Cleanup: delete the created resource
    // err = accountLoginCli.DeleteAccessControlList(result.UUID, param.DeleteModePermissive)
    // if err != nil {
    // 	t.Logf("Cleanup DeleteAccessControlList error: %v", err)
    // }
}
