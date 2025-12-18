// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachAliyunKey operates on AliyunKey
func (cli *ZSClient) DetachAliyunKey(uuid string, params param.DetachAliyunKeyParam) (*view.DetachAliyunKeyEventView, error) {
	resp := view.DetachAliyunKeyEventView{}
	if err := cli.Put("v1/hybrid/aliyun/key/{uuid}/detach", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
