// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVmInstanceMdevDeviceSpecRef queries VmInstanceMdevDeviceSpecRef list
func (cli *ZSClient) QueryVmInstanceMdevDeviceSpecRef(params *param.QueryParam) ([]view.VmInstanceMdevDeviceSpecRefInventoryView, error) {
	var resp []view.VmInstanceMdevDeviceSpecRefInventoryView
	return resp, cli.List("v1/vm-instances/{vmInstanceUuid}/mdev-device-specs", params, &resp)
}

// PageVmInstanceMdevDeviceSpecRef Pagination
func (cli *ZSClient) PageVmInstanceMdevDeviceSpecRef(params *param.QueryParam) ([]view.VmInstanceMdevDeviceSpecRefInventoryView, int, error) {
	var vmInstanceMdevDeviceSpecRefs []view.VmInstanceMdevDeviceSpecRefInventoryView
	total, err := cli.Page("v1/vm-instances/{vmInstanceUuid}/mdev-device-specs", params, &vmInstanceMdevDeviceSpecRefs)
	return vmInstanceMdevDeviceSpecRefs, total, err
}
