// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/util/ptr"
)

func TestQueryGlobalConfig(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10)
	result, err := accountLoginCli.QueryGlobalConfig(&queryParam)
	if err != nil {
		t.Errorf("TestQueryGlobalConfig error: %v", err)
		return
	}
	golog.Infof("======================================")
	golog.Infof("QueryGlobalConfig result count: %d", len(result))
	for _, r := range result {
		golog.Infof("%s\t%s\t%s", r.Category, r.Name, r.Value)
	}
	golog.Infof("======================================")
}

func TestPageGlobalConfig(t *testing.T) {
	queryParam := param.NewQueryParam()
	queryParam.Limit(10).Start(0)
	result, total, err := accountLoginCli.PageGlobalConfig(&queryParam)
	if err != nil {
		t.Errorf("TestPageGlobalConfig error: %v", err)
		return
	}
	golog.Infof("PageGlobalConfig result: total=%d, returned=%d", total, len(result))
}

func TestUpdateGlobalConfig(t *testing.T) {
	enableRemoteWrite := param.UpdateGlobalConfigParam{
		Params: param.UpdateGlobalConfigParamDetail{
			Value: ptr.Of("true"),
		},
	}
	if _, err := accountLoginCli.UpdateGlobalConfig("prometheus", "enable.remote.write", enableRemoteWrite); err != nil {
		t.Errorf("failed to update global config: %v", err)
	}
}
