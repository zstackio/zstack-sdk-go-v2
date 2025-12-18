// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachHybridKey 操作HybridKey
func (cli *ZSClient) DetachHybridKey(uuid string, params param.DetachHybridKeyParam) (*view.DetachHybridKeyEventView, error) {
	resp := view.DetachHybridKeyEventView{}
	if err := cli.Put("v1/hybrid/hybrid/key/{uuid}/detach", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

