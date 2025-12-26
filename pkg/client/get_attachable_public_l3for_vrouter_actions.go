// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetAttachablePublicL3ForVRouter gets AttachablePublicL3ForVRouter by uuid
func (cli *ZSClient) GetAttachablePublicL3ForVRouter(uuid string) (*view.GetAttachablePublicL3ForVRouterView, error) {
	var resp view.GetAttachablePublicL3ForVRouterView
	if err := cli.Get("v1/vm-instances/appliances/virtual-routers/{vmInstanceUuid}/attachable-public-l3s", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
