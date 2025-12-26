// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateEcsImage updates EcsImage
func (cli *ZSClient) UpdateEcsImage(uuid string, params param.UpdateEcsImageParam) (*view.UpdateEcsImageEventView, error) {
	resp := view.UpdateEcsImageEventView{}
	if err := cli.Put("v1/hybrid/aliyun/image/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
