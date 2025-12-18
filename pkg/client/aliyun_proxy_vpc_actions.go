// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateAliyunProxyVpc 更新AliyunProxyVpc
func (cli *ZSClient) UpdateAliyunProxyVpc(uuid string, params param.UpdateAliyunProxyVpcParam) (*view.UpdateAliyunProxyVpcEventView, error) {
	resp := view.UpdateAliyunProxyVpcEventView{}
	if err := cli.Put("v1/aliyun-proxy/vpcs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

