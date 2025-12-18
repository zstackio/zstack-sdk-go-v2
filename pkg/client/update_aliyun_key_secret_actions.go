// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateAliyunKeySecret updates AliyunKeySecret
func (cli *ZSClient) UpdateAliyunKeySecret(uuid string, params param.UpdateAliyunKeySecretParam) (*view.UpdateAliyunKeySecretEventView, error) {
	resp := view.UpdateAliyunKeySecretEventView{}
	if err := cli.Put("v1/hybrid/aliyun/{uuid}/key", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
