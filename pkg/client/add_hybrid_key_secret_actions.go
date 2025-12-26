// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddHybridKeySecret adds HybridKeySecret
func (cli *ZSClient) AddHybridKeySecret(params param.AddHybridKeySecretParam) (*view.AddHybridKeySecretEventView, error) {
	resp := view.AddHybridKeySecretEventView{}
	if err := cli.Post("v1/hybrid/hybrid/key", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
