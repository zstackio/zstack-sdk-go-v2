// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateEcsSecurityGroup updates EcsSecurityGroup
func (cli *ZSClient) UpdateEcsSecurityGroup(uuid string, params param.UpdateEcsSecurityGroupParam) (*view.EcsSecurityGroupInventoryView, error) {
	var resp view.UpdateEcsSecurityGroupEventView
	if err := cli.Put("v1/hybrid/aliyun/security-group/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
