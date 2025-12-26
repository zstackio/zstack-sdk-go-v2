// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachHybridKey operates on HybridKey
func (cli *ZSClient) AttachHybridKey(uuid string, params param.AttachHybridKeyParam) (*view.AttachHybridKeyEventView, error) {
	resp := view.AttachHybridKeyEventView{}
	if err := cli.Put("v1/hybrid/hybrid/key/{uuid}/attach", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
