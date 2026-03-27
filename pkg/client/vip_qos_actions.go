// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteVipQos deletes VipQos
func (cli *ZSClient) DeleteVipQos(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/vips", uuid, string(deleteMode))
}
// SetVipQos operates on VipQos
func (cli *ZSClient) SetVipQos(ctx context.Context, uuid string, params param.SetVipQosParam) (*view.VipQosInventoryView, error) {
	resp := view.VipQosInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/vips", uuid, "", map[string]interface{}{
		"setVipQos": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// GetVipQos gets VipQos by uuid
func (cli *ZSClient) GetVipQos(ctx context.Context, uuid string) (*view.GetVipQosView, error) {
	var resp view.GetVipQosView
	if err := cli.GetWithRespKey(ctx, "v1/vip", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
