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

func TestRemoveSNSDingTalkAtPerson(t *testing.T) {
	// RemoveSNSDingTalkAtPerson operation
	t.Skip("TestRemoveSNSDingTalkAtPerson requires manual implementation")

}

func TestAddSNSDingTalkAtPerson(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddSNSDingTalkAtPerson requires valid creation parameters")

}
