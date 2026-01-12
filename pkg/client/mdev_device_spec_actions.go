// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateMdevDeviceSpec updates MdevDeviceSpec
func (cli *ZSClient) UpdateMdevDeviceSpec(uuid string, params param.UpdateMdevDeviceSpecParam) (*view.MdevDeviceSpecInventoryView, error) {
	var resp view.UpdateMdevDeviceSpecEventView
	err := cli.PutWithSpec("v1/mdev-device-specs", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryMdevDeviceSpec queries MdevDeviceSpec list
func (cli *ZSClient) QueryMdevDeviceSpec(params *param.QueryParam) ([]view.MdevDeviceSpecInventoryView, error) {
	var resp []view.MdevDeviceSpecInventoryView
	return resp, cli.List("v1/mdev-device-specs", params, &resp)
}

func (cli *ZSClient) GetMdevDeviceSpec(uuid string) (*view.MdevDeviceSpecInventoryView, error) {
	var resp view.MdevDeviceSpecInventoryView
	if err := cli.Get("v1/mdev-device-specs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
