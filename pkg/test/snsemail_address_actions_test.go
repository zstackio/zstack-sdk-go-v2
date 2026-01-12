// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSEmailAddress(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySNSEmailAddress(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSEmailAddress error: %v", err)
		return
	}
	golog.Infof("QuerySNSEmailAddress result count: %d", len(result))
}
func TestGetSNSEmailAddress(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySNSEmailAddress(&queryParam)
	if err != nil {
		t.Errorf("TestGetSNSEmailAddress Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSEmailAddress found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetSNSEmailAddress(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSNSEmailAddress error: %v", err)
		return
	}
	golog.Infof("GetSNSEmailAddress result: %s", result.UUID)
}
