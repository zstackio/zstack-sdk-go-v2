// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateVip updates Vip
func (cli *ZSClient) UpdateVip(uuid string, params param.UpdateVipParam) (*view.UpdateVipEventView, error) {
	resp := view.UpdateVipEventView{}
	if err := cli.Put("v1/vips/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
