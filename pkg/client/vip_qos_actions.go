// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteVipQos deletes VipQos
func (cli *ZSClient) DeleteVipQos(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vips", uuid, string(deleteMode))
}
// SetVipQos operates on VipQos
func (cli *ZSClient) SetVipQos(uuid string, params param.SetVipQosParam) (*view.VipQosInventoryView, error) {
	resp := view.VipQosInventoryView{}
	if err := cli.Put("v1/vips", uuid, map[string]interface{}{
		"setVipQos": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// GetVipQos gets VipQos by uuid
func (cli *ZSClient) GetVipQos(uuid string) (*view.VipQosInventoryView, error) {
	var resp view.VipQosInventoryView
	if err := cli.Get("v1/vip", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
