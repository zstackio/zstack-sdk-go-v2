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
	err := cli.PutWithSpec("v1/hybrid/aliyun/security-group", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
