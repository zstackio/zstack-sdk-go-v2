// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachAliyunKey operates on AliyunKey
func (cli *ZSClient) AttachAliyunKey(uuid string, params param.AttachAliyunKeyParam) (*view.AttachAliyunKeyEventView, error) {
	resp := view.AttachAliyunKeyEventView{}
	if err := cli.Put("v1/hybrid/aliyun/key/{uuid}/attach", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
