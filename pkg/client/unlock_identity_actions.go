// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UnlockIdentity operates on UnlockIdentity
func (cli *ZSClient) UnlockIdentity(params param.UnlockIdentityParam) (*view.UnlockIdentityView, error) {
	var resp view.UnlockIdentityView
	if err := cli.Get("v1/login/control/unlock", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
