// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateVpcHaGroup updates VpcHaGroup
func (cli *ZSClient) UpdateVpcHaGroup(uuid string, params param.UpdateVpcHaGroupParam) (*view.UpdateVpcHaGroupEventView, error) {
	resp := view.UpdateVpcHaGroupEventView{}
	if err := cli.Put("v1/vpc/hagroups/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
