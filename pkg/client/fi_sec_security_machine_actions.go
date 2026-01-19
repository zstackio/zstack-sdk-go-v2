// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateFiSecSecurityMachine updates FiSecSecurityMachine
func (cli *ZSClient) UpdateFiSecSecurityMachine(uuid string, params param.UpdateFiSecSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	resp := view.SecurityMachineInventoryView{}
	if err := cli.Put("v1/security-machines/fiSec", uuid, map[string]interface{}{
		"updateFiSecSecurityMachine": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddFiSecSecurityMachine adds FiSecSecurityMachine
func (cli *ZSClient) AddFiSecSecurityMachine(params param.AddFiSecSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	resp := view.SecurityMachineInventoryView{}
	if err := cli.Post("v1/security-machine/fiSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
