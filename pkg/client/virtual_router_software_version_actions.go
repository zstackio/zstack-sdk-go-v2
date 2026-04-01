// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// GetVirtualRouterSoftwareVersion gets VirtualRouterSoftwareVersion by uuid
func (cli *ZSClient) GetVirtualRouterSoftwareVersion() (*view.VirtualRouterSoftwareVersionInventoryView, error) {
	var resp view.VirtualRouterSoftwareVersionInventoryView
	if err := cli.GetWithRespKey("v1/vpc/virtual-routers/softwareversion", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateVirtualRouterSoftwareVersion updates VirtualRouterSoftwareVersion
func (cli *ZSClient) UpdateVirtualRouterSoftwareVersion(uuid string, params param.UpdateVirtualRouterSoftwareVersionParam) (*view.VirtualRouterSoftwareVersionInventoryView, error) {
	resp := view.VirtualRouterSoftwareVersionInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/vpc/virtual-routers/%s/softwareversion", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
