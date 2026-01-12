// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateEcsSecurityGroup updates EcsSecurityGroup
func (cli *ZSClient) UpdateEcsSecurityGroup(uuid string, params param.UpdateEcsSecurityGroupParam) (*view.EcsSecurityGroupInventoryView, error) {
	var resp view.UpdateEcsSecurityGroupEventView
	if err := cli.Put("v1/hybrid/aliyun/security-group", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
