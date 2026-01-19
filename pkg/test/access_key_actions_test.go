// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/util/ptr"
)

func TestQueryAccessKey(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryAccessKey(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAccessKey error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryAccessKey result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.AccessKeyID, r.State)
	}
	golog.Infof("======================================")
}

func TestPageAccessKey(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageAccessKey(&queryParam)
	if err != nil {
		t.Errorf("TestPageAccessKey error: %v", err)
		return
	}
	golog.Infof("PageAccessKey result: total=%d, returned=%d", total, len(result))
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.AccessKeyID, r.State)
	}
}

func TestGetAccessKey(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryAccessKey(&queryParam)
	if err != nil {
		t.Errorf("TestGetAccessKey Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AccessKey found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetAccessKey(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetAccessKey error: %v", err)
		return
	}
	golog.Infof("GetAccessKey result: %s, AccessKeyId: %s", result.UUID, result.AccessKeyID)
}

func TestCreateAccessKey(t *testing.T) {
	if loginSession == nil {
		t.Skip("Skipping CreateAccessKey: loginSession is nil")
		return
	}

	// 1. Create AccessKey using current login session info
	result, err := accountLoginCli.CreateAccessKey(param.CreateAccessKeyParam{
		Params: param.CreateAccessKeyParamDetail{
			UserUuid:    loginSession.UserUuid,
			AccountUuid: loginSession.AccountUuid,
			Description: ptr.Of("Created by SDK Test"),
		},
	})
	if err != nil {
		t.Errorf("CreateAccessKey error: %v", err)
		return
	}
	golog.Infof("Created AccessKey: UUID=%s, ID=%s", result.UUID, result.AccessKeyID)

	// Clean up
	defer func() {
		err := accountLoginCli.DeleteAccessKey(result.UUID, param.DeleteModePermissive)
		if err != nil {
			golog.Errorf("Failed to delete AccessKey %s: %v", result.UUID, err)
		} else {
			golog.Infof("Deleted AccessKey %s", result.UUID)
		}
	}()
}

func TestCreateGetDeleteAccessKeyFlow(t *testing.T) {
	if loginSession == nil {
		t.Skip("Skipping Flow: loginSession is nil")
		return
	}

	// 1. Create
	createResp, err := accountLoginCli.CreateAccessKey(param.CreateAccessKeyParam{
		Params: param.CreateAccessKeyParamDetail{
			UserUuid:    loginSession.UserUuid,
			AccountUuid: loginSession.AccountUuid,
			Description: ptr.Of("Flow Test Key"),
		},
	})
	if err != nil {
		t.Fatalf("CreateAccessKey error: %v", err)
	}
	golog.Infof("Step 1: Created AccessKey %s", createResp.UUID)

	// 2. Get
	getResp, err := accountLoginCli.GetAccessKey(createResp.UUID)
	if err != nil {
		t.Errorf("GetAccessKey error: %v", err)
	} else {
		golog.Infof("Step 2: Get AccessKey success, ID=%s", getResp.AccessKeyID)
		if getResp.UUID != createResp.UUID {
			t.Errorf("Get UUID mismatch: expected %s, got %s", createResp.UUID, getResp.UUID)
		}
	}

	// 3. Delete
	err = accountLoginCli.DeleteAccessKey(createResp.UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("DeleteAccessKey error: %v", err)
	} else {
		golog.Infof("Step 3: Deleted AccessKey %s", createResp.UUID)
	}

	// 4. Verify Delete (Get should fail or return error)
	// Note: GetAccessKey might return 404 error here, which is expected
	_, err = accountLoginCli.GetAccessKey(createResp.UUID)
	if err == nil {
		t.Errorf("Expected error when getting deleted AccessKey, but got nil")
	} else {
		golog.Infof("Step 4: Verify delete success (got expected error: %v)", err)
	}
}
