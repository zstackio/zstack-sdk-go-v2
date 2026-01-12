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
func TestGetApplicationDevelopmentService(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryApplicationDevelopmentService(&queryParam)
	if err != nil {
		t.Errorf("TestGetApplicationDevelopmentService Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No ApplicationDevelopmentService found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetApplicationDevelopmentService(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetApplicationDevelopmentService error: %v", err)
		return
	}
	golog.Infof("GetApplicationDevelopmentService result: %s", result.UUID)
}
