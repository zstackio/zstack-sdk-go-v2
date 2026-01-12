// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryZBox(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryZBox(&queryParam)
	if err != nil {
		t.Errorf("TestQueryZBox error: %v", err)
		return
	}
	golog.Infof("QueryZBox result count: %d", len(result))
}
func TestGetZBox(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryZBox(&queryParam)
	if err != nil {
		t.Errorf("TestGetZBox Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ZBox found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetZBox(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetZBox error: %v", err)
		return
	}
	golog.Infof("GetZBox result: %s", result.UUID)
}

func TestAddZBox(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddZBox requires valid creation parameters")

}
