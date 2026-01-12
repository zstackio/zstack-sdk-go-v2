// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSEndpointThirdpartyAlertHistory(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySNSEndpointThirdpartyAlertHistory(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSEndpointThirdpartyAlertHistory error: %v", err)
		return
	}
	golog.Infof("QuerySNSEndpointThirdpartyAlertHistory result count: %d", len(result))
}
func TestGetSNSEndpointThirdpartyAlertHistory(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySNSEndpointThirdpartyAlertHistory(&queryParam)
	if err != nil {
		t.Errorf("TestGetSNSEndpointThirdpartyAlertHistory Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSEndpointThirdpartyAlertHistory found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetSNSEndpointThirdpartyAlertHistory(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSNSEndpointThirdpartyAlertHistory error: %v", err)
		return
	}
	golog.Infof("GetSNSEndpointThirdpartyAlertHistory result: %s", result.UUID)
}
