// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryCertificate(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryCertificate(&queryParam)
	if err != nil {
		t.Errorf("TestQueryCertificate error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryCertificate result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s", r.UUID, r.Name)
	}
	golog.Infof("======================================")
}

func TestPageCertificate(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageCertificate(&queryParam)
	if err != nil {
		t.Errorf("TestPageCertificate error: %v", err)
		return
	}
	golog.Infof("PageCertificate result: total=%d, returned=%d", total, len(result))
}

func TestGetCertificate(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryCertificate(&queryParam)
	if err != nil {
		t.Errorf("TestGetCertificate Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Certificate found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetCertificate(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetCertificate error: %v", err)
		return
	}
	golog.Infof("GetCertificate result: %s, Name: %s", result.UUID, result.Name)
}
