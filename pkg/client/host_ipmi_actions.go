// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateHostIpmi updates HostIpmi
func (cli *ZSClient) UpdateHostIpmi(uuid string, params param.UpdateHostIpmiParam) (*view.HostIpmiInventoryView, error) {
	resp := view.HostIpmiInventoryView{}
	if err := cli.Put("v1/hosts/ipmi/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
