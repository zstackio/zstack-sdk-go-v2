// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateEcsImage updates EcsImage
func (cli *ZSClient) UpdateEcsImage(uuid string, params param.UpdateEcsImageParam) (*view.EcsImageInventoryView, error) {
	var resp view.UpdateEcsImageEventView
	err := cli.PutWithSpec("v1/hybrid/aliyun/image", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
