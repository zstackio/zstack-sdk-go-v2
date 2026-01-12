// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySharedBlockGroupPrimaryStorage(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySharedBlockGroupPrimaryStorage(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySharedBlockGroupPrimaryStorage error: %v", err)
		return
	}
	golog.Infof("QuerySharedBlockGroupPrimaryStorage result count: %d", len(result))
}
func TestGetSharedBlockGroupPrimaryStorage(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySharedBlockGroupPrimaryStorage(&queryParam)
	if err != nil {
		t.Errorf("TestGetSharedBlockGroupPrimaryStorage Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SharedBlockGroupPrimaryStorage found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetSharedBlockGroupPrimaryStorage(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSharedBlockGroupPrimaryStorage error: %v", err)
		return
	}
	golog.Infof("GetSharedBlockGroupPrimaryStorage result: %s", result.UUID)
}

func TestAddSharedBlockGroupPrimaryStorage(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddSharedBlockGroupPrimaryStorage requires valid creation parameters")

}
