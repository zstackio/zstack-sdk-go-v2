// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSWeComAtPerson(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySNSWeComAtPerson(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSWeComAtPerson error: %v", err)
		return
	}
	golog.Infof("QuerySNSWeComAtPerson result count: %d", len(result))
}
func TestGetSNSWeComAtPerson(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySNSWeComAtPerson(&queryParam)
	if err != nil {
		t.Errorf("TestGetSNSWeComAtPerson Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSWeComAtPerson found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetSNSWeComAtPerson(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSNSWeComAtPerson error: %v", err)
		return
	}
	golog.Infof("GetSNSWeComAtPerson result: %s", result.UUID)
}

func TestAddSNSWeComAtPerson(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddSNSWeComAtPerson requires valid creation parameters")

}

func TestRemoveSNSWeComAtPerson(t *testing.T) {
	// RemoveSNSWeComAtPerson operation
	t.Skip("TestRemoveSNSWeComAtPerson requires manual implementation")

}
