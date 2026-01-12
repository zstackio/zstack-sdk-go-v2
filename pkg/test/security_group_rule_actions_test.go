// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestQuerySecurityGroupRule(t *testing.T) {
	queryParam := param.NewQueryParam()
	result, err := accountLoginCli.QuerySecurityGroupRule(&queryParam)
	if err != nil {
		t.Errorf("TestQuerySecurityGroupRule error: %v", err)
		return
	}
	golog.Infof("QuerySecurityGroupRule result count: %d", len(result))
}

func TestDeleteSecurityGroupRule(t *testing.T) {
	// WARNING: This test will actually delete a resource!
	// Query first to get UUID (but skip by default to avoid accidental deletion)
	t.Skip("TestDeleteSecurityGroupRule is skipped by default to prevent accidental deletion. Remove this line to enable.")

	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QuerySecurityGroupRule(&queryParam)
	if err != nil {
		t.Errorf("TestDeleteSecurityGroupRule Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No SecurityGroupRule found to test Delete")
		return
	}

	err = accountLoginCli.DeleteSecurityGroupRule(list[0].Uuid, param.DeleteModePermissive)
	if err != nil {
		t.Errorf("TestDeleteSecurityGroupRule error: %v", err)
		return
	}
	golog.Infof("DeleteSecurityGroupRule succeeded for UUID: %s", list[0].Uuid)
}

func TestChangeSecurityGroupRule(t *testing.T) {
	// Change operation
	t.Skip("TestChangeSecurityGroupRule requires specific parameters")

}

func TestValidateSecurityGroupRule(t *testing.T) {
	// ValidateSecurityGroupRule operation
	t.Skip("TestValidateSecurityGroupRule requires manual implementation")

}

func TestAddSecurityGroupRule(t *testing.T) {
	// Add operation - similar to Create
	t.Skip("TestAddSecurityGroupRule requires valid creation parameters")

}
