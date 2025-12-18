// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddAliyunKeySecret adds AliyunKeySecret
func (cli *ZSClient) AddAliyunKeySecret(params param.AddAliyunKeySecretParam) (*view.AddAliyunKeySecretEventView, error) {
	resp := view.AddAliyunKeySecretEventView{}
	if err := cli.Post("v1/hybrid/aliyun/key", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
