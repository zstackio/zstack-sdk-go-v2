// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryLicenseAuthorizedNode(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryLicenseAuthorizedNode(&queryParam)
	if err != nil {
		t.Errorf("TestQueryLicenseAuthorizedNode error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryLicenseAuthorizedNode result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s", r.UUID, r.AppId)
	}
	golog.Infof("======================================")
}

func TestPageLicenseAuthorizedNode(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageLicenseAuthorizedNode(&queryParam)
	if err != nil {
		t.Errorf("TestPageLicenseAuthorizedNode error: %v", err)
		return
	}
	golog.Infof("PageLicenseAuthorizedNode result: total=%d, returned=%d", total, len(result))
}

func TestGetLicenseAuthorizedNode(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryLicenseAuthorizedNode(&queryParam)
	if err != nil {
		t.Errorf("TestGetLicenseAuthorizedNode Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No LicenseAuthorizedNode found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetLicenseAuthorizedNode(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetLicenseAuthorizedNode error: %v", err)
		return
	}
	golog.Infof("GetLicenseAuthorizedNode result: %s", result.UUID)
}
