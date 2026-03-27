// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateFiSecSecurityMachine updates FiSecSecurityMachine
func (cli *ZSClient) UpdateFiSecSecurityMachine(ctx context.Context, uuid string, params param.UpdateFiSecSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	resp := view.SecurityMachineInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/security-machines/fiSec", uuid, "", map[string]interface{}{
		"updateFiSecSecurityMachine": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddFiSecSecurityMachine adds FiSecSecurityMachine
func (cli *ZSClient) AddFiSecSecurityMachine(ctx context.Context, params param.AddFiSecSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	resp := view.SecurityMachineInventoryView{}
	if err := cli.Post(ctx, "v1/security-machine/fiSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
