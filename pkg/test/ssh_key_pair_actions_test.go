// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySshKeyPair(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accessKeyAuthCli.QuerySshKeyPair(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySshKeyPair error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QuerySshKeyPair result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s", r.UUID, r.Name)
	}
	golog.Infof("======================================")
}

func TestPageSshKeyPair(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accessKeyAuthCli.PageSshKeyPair(&queryParam)
	if err != nil {
		t.Errorf("TestPageSshKeyPair error: %v", err)
		return
	}
	golog.Infof("PageSshKeyPair result: total=%d, returned=%d", total, len(result))
	golog.Infof("======================================")
	for _, r := range result {
		golog.Infof("%s\t%s", r.UUID, r.Name)
	}
}

func TestGetSshKeyPair(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accessKeyAuthCli.QuerySshKeyPair(&queryParam)
	if err != nil {
		t.Errorf("TestGetSshKeyPair Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SshKeyPair found to test Get")
		return
	}

	result, err := accessKeyAuthCli.GetSshKeyPair(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSshKeyPair error: %v", err)
		return
	}
	golog.Infof("GetSshKeyPair result: %s, Name: %s", result.UUID, result.Name)
}
