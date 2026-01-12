// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryImageGroup(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryImageGroup(&queryParam)
	if err != nil {
		t.Errorf("TestQueryImageGroup error: %v", err)
		return
	}
	golog.Infof("QueryImageGroup result count: %d", len(result))
}

func TestExpungeImageGroup(t *testing.T) {
	// Expunge operation - permanently deletes
	t.Skip("TestExpungeImageGroup is dangerous - permanently deletes resource")

}
