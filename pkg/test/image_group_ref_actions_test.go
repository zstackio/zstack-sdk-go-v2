// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryImageGroupRef(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryImageGroupRef(&queryParam)
	if err != nil {
		t.Errorf("TestQueryImageGroupRef error: %v", err)
		return
	}
	golog.Infof("QueryImageGroupRef result count: %d", len(result))
}
func TestGetImageGroupRef(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryImageGroupRef(&queryParam)
	if err != nil {
		t.Errorf("TestGetImageGroupRef Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ImageGroupRef found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetImageGroupRef(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetImageGroupRef error: %v", err)
		return
	}
	golog.Infof("GetImageGroupRef result: %s", result.UUID)
}
