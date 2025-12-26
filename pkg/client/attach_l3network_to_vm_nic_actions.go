// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachL3NetworkToVmNic operates on L3NetworkToVmNic
func (cli *ZSClient) AttachL3NetworkToVmNic(params param.AttachL3NetworkToVmNicParam) (*view.AttachL3NetworkToVmNicEventView, error) {
	resp := view.AttachL3NetworkToVmNicEventView{}
	if err := cli.Post("v1/nics/{vmNicUuid}/l3-networks/{l3NetworkUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
