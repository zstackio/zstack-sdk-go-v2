// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// DetachHybridKey operates on HybridKey
func (cli *ZSClient) DetachHybridKey(uuid string, params param.DetachHybridKeyParam) (*view.DetachHybridKeyEventView, error) {
	resp := view.DetachHybridKeyEventView{}
	if err := cli.Put("v1/hybrid/hybrid/key/{uuid}/detach", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
