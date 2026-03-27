// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddSanSecSecurityMachine adds SanSecSecurityMachine
func (cli *ZSClient) AddSanSecSecurityMachine(ctx context.Context, params param.AddSanSecSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	resp := view.SecurityMachineInventoryView{}
	if err := cli.Post(ctx, "v1/security-machine/sanSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateSanSecSecurityMachine updates SanSecSecurityMachine
func (cli *ZSClient) UpdateSanSecSecurityMachine(ctx context.Context, uuid string, params param.UpdateSanSecSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	resp := view.SecurityMachineInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/security-machines/sanSec", uuid, "", map[string]interface{}{
		"updateSanSecSecurityMachine": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
