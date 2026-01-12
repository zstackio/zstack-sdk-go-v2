// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySharedResource(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySharedResource(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySharedResource error: %v", err)
		return
	}
	golog.Infof("QuerySharedResource result count: %d", len(result))
}
func TestGetSharedResource(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySharedResource(&queryParam)
	if err != nil {
		t.Errorf("TestGetSharedResource Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SharedResource found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetSharedResource(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSharedResource error: %v", err)
		return
	}
	golog.Infof("GetSharedResource result: %s", result.UUID)
}
