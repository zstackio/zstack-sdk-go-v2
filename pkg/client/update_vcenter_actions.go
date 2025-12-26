// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateVCenter updates VCenter
func (cli *ZSClient) UpdateVCenter(uuid string, params param.UpdateVCenterParam) (*view.UpdateVCenterEventView, error) {
	resp := view.UpdateVCenterEventView{}
	if err := cli.Put("v1/vcenters/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
