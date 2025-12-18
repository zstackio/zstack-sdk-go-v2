// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetVmInstanceHaLevel operates on SetVmInstanceHaLevel
func (cli *ZSClient) SetVmInstanceHaLevel(params param.SetVmInstanceHaLevelParam) (*view.SetVmInstanceHaLevelEventView, error) {
	resp := view.SetVmInstanceHaLevelEventView{}
	if err := cli.Post("v1/vm-instances/{uuid}/ha-levels", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
