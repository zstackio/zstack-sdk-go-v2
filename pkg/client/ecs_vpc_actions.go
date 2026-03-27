// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateEcsVpc updates EcsVpc
func (cli *ZSClient) UpdateEcsVpc(ctx context.Context, uuid string, params param.UpdateEcsVpcParam) (*view.EcsVpcInventoryView, error) {
	resp := view.EcsVpcInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/hybrid/aliyun/vpc", uuid, "", map[string]interface{}{
		"updateEcsVpc": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
