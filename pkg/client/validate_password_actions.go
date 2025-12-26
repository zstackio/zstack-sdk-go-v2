// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ValidatePassword operates on ValidatePassword
func (cli *ZSClient) ValidatePassword(uuid string, params param.ValidatePasswordParam) (*view.ValidatePasswordView, error) {
	resp := view.ValidatePasswordView{}
	if err := cli.Put("v1/password/verify", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
