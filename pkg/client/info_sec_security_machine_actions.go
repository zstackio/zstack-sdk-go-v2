// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddInfoSecSecurityMachine adds InfoSecSecurityMachine
func (cli *ZSClient) AddInfoSecSecurityMachine(ctx context.Context, params param.AddInfoSecSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	resp := view.SecurityMachineInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/security-machine/infoSec", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateInfoSecSecurityMachine updates InfoSecSecurityMachine
func (cli *ZSClient) UpdateInfoSecSecurityMachine(ctx context.Context, uuid string, params param.UpdateInfoSecSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	resp := view.SecurityMachineInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/security-machines/infoSec", uuid, "", map[string]interface{}{
		"updateInfoSecSecurityMachine": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
