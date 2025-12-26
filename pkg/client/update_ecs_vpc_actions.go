// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateEcsVpc updates EcsVpc
func (cli *ZSClient) UpdateEcsVpc(uuid string, params param.UpdateEcsVpcParam) (*view.UpdateEcsVpcEventView, error) {
	resp := view.UpdateEcsVpcEventView{}
	if err := cli.Put("v1/hybrid/aliyun/vpc/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
