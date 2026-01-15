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
	resp := view.MdevDeviceSpecInventoryView{}
	if err := cli.Put("v1/mdev-device-specs", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryMdevDeviceSpec queries MdevDeviceSpec list
func (cli *ZSClient) QueryMdevDeviceSpec(params *param.QueryParam) ([]view.MdevDeviceSpecInventoryView, error) {
	var resp []view.MdevDeviceSpecInventoryView
	return resp, cli.List("v1/mdev-device-specs", params, &resp)
}

// PageMdevDeviceSpec Pagination
func (cli *ZSClient) PageMdevDeviceSpec(params *param.QueryParam) ([]view.MdevDeviceSpecInventoryView, int, error) {
	var mdevDeviceSpecs []view.MdevDeviceSpecInventoryView
	total, err := cli.Page("v1/mdev-device-specs", params, &mdevDeviceSpecs)
	return mdevDeviceSpecs, total, err
}
