// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateHybridEip updates HybridEip
func (cli *ZSClient) UpdateHybridEip(uuid string, params param.UpdateHybridEipParam) (*view.UpdateHybridEipEventView, error) {
	resp := view.UpdateHybridEipEventView{}
	if err := cli.Put("v1/hybrid/eip/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
