// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateFlkSecSecurityMachine updates FlkSecSecurityMachine
func (cli *ZSClient) UpdateFlkSecSecurityMachine(ctx context.Context, uuid string, params param.UpdateFlkSecSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	resp := view.SecurityMachineInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/security-machines/flkSec", uuid, "", map[string]interface{}{
		"updateFlkSecSecurityMachine": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddFlkSecSecurityMachine adds FlkSecSecurityMachine
func (cli *ZSClient) AddFlkSecSecurityMachine(ctx context.Context, params param.AddFlkSecSecurityMachineParam) (*view.SecurityMachineInventoryView, error) {
	resp := view.SecurityMachineInventoryView{}
	if err := cli.Post(ctx, "v1/security-machine/flkSec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
