// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateMdevDeviceSpec updates MdevDeviceSpec
func (cli *ZSClient) UpdateMdevDeviceSpec(uuid string, params param.UpdateMdevDeviceSpecParam) (*view.MdevDeviceSpecInventoryView, error) {
	var resp view.UpdateMdevDeviceSpecEventView
	if err := cli.Put("v1/mdev-device-specs/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryMdevDeviceSpec queries MdevDeviceSpec list
func (cli *ZSClient) QueryMdevDeviceSpec(params *param.QueryParam) ([]view.MdevDeviceSpecInventoryView, error) {
	var resp []view.MdevDeviceSpecInventoryView
	return resp, cli.List("v1/mdev-device-specs", params, &resp)
}
