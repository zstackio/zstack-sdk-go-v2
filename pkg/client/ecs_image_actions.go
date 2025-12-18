// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateEcsImage 更新EcsImage
func (cli *ZSClient) UpdateEcsImage(uuid string, params param.UpdateEcsImageParam) (*view.UpdateEcsImageEventView, error) {
	resp := view.UpdateEcsImageEventView{}
	if err := cli.Put("v1/hybrid/aliyun/image/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

