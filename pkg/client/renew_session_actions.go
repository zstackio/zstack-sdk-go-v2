// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RenewSession operates on RenewSession
func (cli *ZSClient) RenewSession(uuid string, params param.RenewSessionParam) (*view.RenewSessionEventView, error) {
	resp := view.RenewSessionEventView{}
	if err := cli.Put("v1/accounts/sessions/{sessionUuid}/renew", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
