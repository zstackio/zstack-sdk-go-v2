// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryLdapServer(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QueryLdapServer(&queryParam)
	if err != nil {
		t.Errorf("TestQueryLdapServer error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryLdapServer result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.UUID, r.Name, r.Url)
	}
	golog.Infof("======================================")
}

func TestPageLdapServer(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageLdapServer(&queryParam)
	if err != nil {
		t.Errorf("TestPageLdapServer error: %v", err)
		return
	}
	golog.Infof("PageLdapServer result: total=%d, returned=%d", total, len(result))
}

func TestGetLdapServer(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QueryLdapServer(&queryParam)
	if err != nil {
		t.Errorf("TestGetLdapServer Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No LdapServer found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetLdapServer(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetLdapServer error: %v", err)
		return
	}
	golog.Infof("GetLdapServer result: %s, Name: %s", result.UUID, result.Name)
}
