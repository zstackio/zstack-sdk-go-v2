// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCreateEcsImageProgress gets CreateEcsImageProgress by uuid
func (cli *ZSClient) GetCreateEcsImageProgress(uuid string) (*view.GetCreateEcsImageProgressView, error) {
	var resp view.GetCreateEcsImageProgressView
	if err := cli.Get("v1/hybrid/aliyun/image/{dataCenterUuid}/{imageUuid}/progress", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
