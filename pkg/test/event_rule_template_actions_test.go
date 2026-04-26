// Copyright (c) ZStack.io, Inc.
// Auto-generated integration tests. DO NOT EDIT.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQueryEventRuleTemplate(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QueryEventRuleTemplate(&queryParam)
	if err != nil {
		t.Errorf("TestQueryEventRuleTemplate error: %v", err)
		return
	}
	golog.Infof("QueryEventRuleTemplate result count: %d", len(result))
}

