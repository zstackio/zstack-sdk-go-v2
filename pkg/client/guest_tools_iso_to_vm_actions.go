// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachGuestToolsIsoToVm 操作GuestToolsIsoToVm
func (cli *ZSClient) AttachGuestToolsIsoToVm(uuid string, params param.AttachGuestToolsIsoToVmParam) (*view.AttachGuestToolsIsoToVmEventView, error) {
	resp := view.AttachGuestToolsIsoToVmEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

