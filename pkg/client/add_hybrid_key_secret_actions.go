// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddHybridKeySecret 操作AddHybridKeySecret
func (cli *ZSClient) AddHybridKeySecret(params param.AddHybridKeySecretParam) (*view.AddHybridKeySecretEventView, error) {
	resp := view.AddHybridKeySecretEventView{}
	if err := cli.Post("v1/hybrid/hybrid/key", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

