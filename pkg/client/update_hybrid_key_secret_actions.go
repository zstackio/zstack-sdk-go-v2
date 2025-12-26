// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateHybridKeySecret updates HybridKeySecret
func (cli *ZSClient) UpdateHybridKeySecret(uuid string, params param.UpdateHybridKeySecretParam) (*view.UpdateHybridKeySecretEventView, error) {
	resp := view.UpdateHybridKeySecretEventView{}
	if err := cli.Put("v1/hybrid/hybrid/{uuid}/key", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
