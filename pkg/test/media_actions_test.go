// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryMedia(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryMedia(&queryParam)
	if err != nil {
		t.Errorf("TestQueryMedia error: %v", err)
		return
	}
	golog.Infof("QueryMedia result count: %d", len(result))
}
func TestGetMedia(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMedia(&queryParam)
	if err != nil {
		t.Errorf("TestGetMedia Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Media found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetMedia(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetMedia error: %v", err)
		return
	}
	golog.Infof("GetMedia result: %s", result.UUID)
}

func TestDeleteMedia(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteMedia is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryMedia(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteMedia Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No Media found to test Delete")
		return
	}

	err = accountLoginCli.DeleteMedia(list[0].UUID, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteMedia error: %v", err)
		return
	}
	golog.Infof("DeleteMedia succeeded for UUID: %s", list[0].UUID)
}
