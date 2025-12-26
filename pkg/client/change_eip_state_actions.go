// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeEipState changes EipState
func (cli *ZSClient) ChangeEipState(uuid string, params param.ChangeEipStateParam) (*view.ChangeEipStateEventView, error) {
	resp := view.ChangeEipStateEventView{}
	if err := cli.Put("v1/eips/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
