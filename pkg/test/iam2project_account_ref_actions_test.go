// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryIAM2ProjectAccountRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryIAM2ProjectAccountRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryIAM2ProjectAccountRef error: %v", err)
		return
	}
	golog.Infof("QueryIAM2ProjectAccountRef result count: %d", len(result))
}
func TestGetIAM2ProjectAccountRef(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryIAM2ProjectAccountRef(&queryParam)
	if err != nil {
		t.Errorf("TestGetIAM2ProjectAccountRef Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No IAM2ProjectAccountRef found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetIAM2ProjectAccountRef(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetIAM2ProjectAccountRef error: %v", err)
		return
	}
	golog.Infof("GetIAM2ProjectAccountRef result: %s", result.UUID)
}
