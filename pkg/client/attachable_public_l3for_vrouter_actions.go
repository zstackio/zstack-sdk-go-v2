// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetAttachablePublicL3ForVRouter 获取AttachablePublicL3ForVRouter详情
func (cli *ZSClient) GetAttachablePublicL3ForVRouter(uuid string) (*view.GetAttachablePublicL3ForVRouterView, error) {
	var resp view.GetAttachablePublicL3ForVRouterView
	if err := cli.Get("v1/vm-instances/appliances/virtual-routers/{vmInstanceUuid}/attachable-public-l3s", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

