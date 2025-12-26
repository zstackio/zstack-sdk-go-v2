// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ValidateSession operates on ValidateSession
func (cli *ZSClient) ValidateSession(params param.ValidateSessionParam) (*view.ValidateSessionView, error) {
	var resp view.ValidateSessionView
	if err := cli.Get("v1/accounts/sessions/{sessionUuid}/valid", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
