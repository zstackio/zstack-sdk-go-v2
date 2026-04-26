// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryCCSCertificate(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryCCSCertificate(&queryParam)
	if err != nil {
		t.Errorf("TestQueryCCSCertificate error: %v", err)
		return
	}
	golog.Infof("QueryCCSCertificate result count: %d", len(result))
}

