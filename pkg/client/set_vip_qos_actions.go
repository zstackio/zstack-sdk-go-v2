// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetVipQos operates on SetVipQos
func (cli *ZSClient) SetVipQos(uuid string, params param.SetVipQosParam) (*view.SetVipQosEventView, error) {
	resp := view.SetVipQosEventView{}
	if err := cli.Put("v1/vips/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
