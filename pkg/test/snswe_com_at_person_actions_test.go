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

func TestAddSNSWeComAtPerson(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddSNSWeComAtPerson requires valid creation parameters")

}

func TestRemoveSNSWeComAtPerson(t *testing.T) {
	// RemoveSNSWeComAtPerson operation
	t.Skip("TestRemoveSNSWeComAtPerson requires manual implementation")

}
