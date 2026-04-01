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
	resp := view.EcsImageInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/image", uuid, "", map[string]interface{}{
		"updateEcsImage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
