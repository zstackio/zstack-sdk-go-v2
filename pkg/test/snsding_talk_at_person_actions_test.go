// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySNSDingTalkAtPerson(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySNSDingTalkAtPerson(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySNSDingTalkAtPerson error: %v", err)
		return
	}
	golog.Infof("QuerySNSDingTalkAtPerson result count: %d", len(result))
}
func TestGetSNSDingTalkAtPerson(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySNSDingTalkAtPerson(&queryParam)
	if err != nil {
		t.Errorf("TestGetSNSDingTalkAtPerson Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SNSDingTalkAtPerson found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetSNSDingTalkAtPerson(list[0].UUID)
	if err != nil {
		t.Errorf("TestGetSNSDingTalkAtPerson error: %v", err)
		return
	}
	golog.Infof("GetSNSDingTalkAtPerson result: %s", result.UUID)
}

func TestRemoveSNSDingTalkAtPerson(t *testing.T) {
	// RemoveSNSDingTalkAtPerson operation
	t.Skip("TestRemoveSNSDingTalkAtPerson requires manual implementation")

}

func TestAddSNSDingTalkAtPerson(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddSNSDingTalkAtPerson requires valid creation parameters")

}
