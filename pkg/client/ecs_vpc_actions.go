// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateEcsVpc 更新EcsVpc
func (cli *ZSClient) UpdateEcsVpc(uuid string, params param.UpdateEcsVpcParam) (*view.UpdateEcsVpcEventView, error) {
	resp := view.UpdateEcsVpcEventView{}
	if err := cli.Put("v1/hybrid/aliyun/vpc/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

