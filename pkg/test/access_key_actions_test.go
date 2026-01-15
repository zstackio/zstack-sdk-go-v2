// Copyright (c) ZStack.io, Inc.

package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryAccessKey(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryAccessKey(&queryParam)
	if err != nil {
		t.Errorf("TestQueryAccessKey error: %v", err)
		return
	}
	golog.Infof("QueryAccessKey result count: %d", len(result))
}
func TestGetAccessKey(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAccessKey(&queryParam)
	if err != nil {
		t.Errorf("TestGetAccessKey Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AccessKey found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetAccessKey(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetAccessKey error: %v", err)
		return
	}
	golog.Infof("GetAccessKey result: %s", result.UUID)
}

func TestDeleteAccessKey(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	//t.Skip("TestDeleteAccessKey is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryAccessKey(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteAccessKey Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No AccessKey found to test Delete")
		return
	}

	err = accountLoginCli.DeleteAccessKey(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteAccessKey error: %v", err)
		return
	}
	golog.Infof("DeleteAccessKey succeeded for UUID: %s", list[0].UUID)
}

func TestCreateAccessKey(t *testing.T) {
	// WARNING: This test will create a real resource!
	//t.Skip("TestCreateAccessKey is skipped by default. Implement with valid params to test creation.")

	sess, err := accountLoginCli.Login()
	if err != nil {
		t.Fatalf("Login error: %v", err)
	}
	defer accountLoginCli.Logout()
	// sess contains AccountUuid and UserUuid for the currently logged-in session

	// Build a unique AccessKeyID so that repeated test runs don't collide
	// Max length 20 chars. Current: 4 + 10 = 14 chars
	//accessKeyID := fmt.Sprintf("sdk-%d", time.Now().Unix())

	createParam := param.CreateAccessKeyParam{
		BaseParam: param.BaseParam{},
		CreateAccessKey: param.CreateAccessKeyParamDetail{
			AccountUuid: sess.AccountUuid,
			UserUuid:    sess.UserUuid,
			Description: "Chi-test",
			//AccessKeyID: accessKeyID,
		},
	}
	result, err := accountLoginCli.CreateAccessKey(createParam)
	if err != nil {
		t.Errorf("TestCreateAccessKey error: %v", err)
		return
	}
	golog.Infof("CreateAccessKey result: %s", result.UUID)

}

func TestCreateGetDeleteAccessKeyFlow(t *testing.T) {
	// Ensure we have a logged-in client and obtain account & user UUIDs
	sess, err := accountLoginCli.Login()
	if err != nil {
		t.Fatalf("Login error: %v", err)
	}
	defer accountLoginCli.Logout()
	// sess contains AccountUuid and UserUuid for the currently logged-in session

	// Build a unique AccessKeyID so that repeated test runs don't collide
	// Max length 20 chars. Current: 4 + 10 = 14 chars
	// Build a unique AccessKeyID so that repeated test runs don't collide
	// Max length 20 chars. Current: 4 + 10 = 14 chars
	accessKeyID := fmt.Sprintf("sdk-%d", time.Now().Unix())
	accessKeySecret := fmt.Sprintf("secret-%d", time.Now().Unix())

	createParam := param.CreateAccessKeyParam{
		BaseParam: param.BaseParam{},
		CreateAccessKey: param.CreateAccessKeyParamDetail{
			AccountUuid:     sess.AccountUuid,
			UserUuid:        sess.UserUuid,
			Description:     "Chi-test",
			AccessKeyID:     accessKeyID,
			AccessKeySecret: accessKeySecret,
		},
	}

	// Create the access key
	resp, err := accountLoginCli.CreateAccessKey(createParam)
	if err != nil {
		t.Fatalf("CreateAccessKey error: %v", err)
	}
	if resp == nil {
		t.Fatalf("CreateAccessKey returned nil response")
	}
	golog.Infof("Created access key: %v", resp)

	// Try to get the access key by UUID
	getResp, err := accountLoginCli.GetAccessKey(resp.UUID)
	if err != nil {
		t.Fatalf("GetAccessKey error: %v", err)
	}
	golog.Infof("Get access key: %v", getResp)

	// Delete the access key
	err = accountLoginCli.DeleteAccessKey(resp.UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("DeleteAccessKey error: %v", err)
		// don't return right away; we may still want to check deletion or proceed to cleanup in other tests
	}

	// Optionally verify that querying with the accessKeyID returns none or not; skipping strict assertion for environment variance
}
