// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeAccessKeyState changes AccessKeyState
func (cli *ZSClient) ChangeAccessKeyState(uuid string, params param.ChangeAccessKeyStateParam) (*view.ChangeAccessKeyStateEventView, error) {
	resp := view.ChangeAccessKeyStateEventView{}
	if err := cli.Put("v1/accesskeys/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
