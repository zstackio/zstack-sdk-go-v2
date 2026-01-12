// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryApplicationDevelopmentService(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryApplicationDevelopmentService(&queryParam)
	if err != nil {
		t.Errorf("TestQueryApplicationDevelopmentService error: %v", err)
		return
	}
	golog.Infof("QueryApplicationDevelopmentService result count: %d", len(result))
}
