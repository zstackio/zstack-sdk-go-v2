// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSFeiShuAtPerson(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySNSFeiShuAtPerson(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSFeiShuAtPerson error: %v", err)
		return
	}
	golog.Infof("QuerySNSFeiShuAtPerson result count: %d", len(result))
}
func TestGetSNSFeiShuAtPerson(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySNSFeiShuAtPerson(&queryParam)
	if err != nil {
		t.Errorf("TestGetSNSFeiShuAtPerson Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSFeiShuAtPerson found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetSNSFeiShuAtPerson(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSNSFeiShuAtPerson error: %v", err)
		return
	}
	golog.Infof("GetSNSFeiShuAtPerson result: %s", result.UUID)
}

func TestRemoveSNSFeiShuAtPerson(t *testing.T) {
	// RemoveSNSFeiShuAtPerson operation
	t.Skip("TestRemoveSNSFeiShuAtPerson requires manual implementation")

}

func TestAddSNSFeiShuAtPerson(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddSNSFeiShuAtPerson requires valid creation parameters")

}
