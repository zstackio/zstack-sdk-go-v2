// Copyright (c) ZStack.io, Inc.

package test

import (
	"testing"

	"github.com/kataras/golog"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

func TestGetTwoFactorAuthenticationSecret(t *testing.T) {
	// First query to get a valid UUID
	queryParam := param.NewQueryParam()
	queryParam.Limit(1)
	list, err := accountLoginCli.QueryTwoFactorAuthentication(&queryParam)
	if err != nil {
		t.Errorf("TestGetTwoFactorAuthenticationSecret Query error: %v", err)
		return
	}
	if len(list) == 0 {
		t.Skip("No TwoFactorAuthenticationSecret found to test Get")
		return
	}

	// Get by UUID
	result, err := accountLoginCli.GetTwoFactorAuthenticationSecret()
	if err != nil {
		t.Errorf("TestGetTwoFactorAuthenticationSecret error: %v", err)
		return
	}
	golog.Infof("GetTwoFactorAuthenticationSecret result: %s", result.UUID)
}

func TestResetTwoFactorAuthenticationSecret(t *testing.T) {
	// Reset operation
	t.Skip("TestResetTwoFactorAuthenticationSecret may affect resource state")

}
