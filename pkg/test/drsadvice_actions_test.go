// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryDRSAdvice(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryDRSAdvice(&queryParam)
	if err != nil {
		t.Errorf("TestQueryDRSAdvice error: %v", err)
		return
	}
	golog.Infof("QueryDRSAdvice result count: %d", len(result))
}
func TestGetDRSAdvice(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryDRSAdvice(&queryParam)
	if err != nil {
		t.Errorf("TestGetDRSAdvice Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No DRSAdvice found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetDRSAdvice(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetDRSAdvice error: %v", err)
		return
	}
	golog.Infof("GetDRSAdvice result: %s", result.UUID)
}

func TestApplyDRSAdvice(t *testing.T) {
	// ApplyDRSAdvice operation
	t.Skip("TestApplyDRSAdvice requires manual implementation")

}
