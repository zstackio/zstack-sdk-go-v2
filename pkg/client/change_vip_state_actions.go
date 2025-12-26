// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeVipState changes VipState
func (cli *ZSClient) ChangeVipState(uuid string, params param.ChangeVipStateParam) (*view.ChangeVipStateEventView, error) {
	resp := view.ChangeVipStateEventView{}
	if err := cli.Put("v1/vips/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
