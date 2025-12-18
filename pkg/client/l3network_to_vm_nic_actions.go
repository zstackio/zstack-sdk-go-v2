// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachL3NetworkToVmNic 操作L3NetworkToVmNic
func (cli *ZSClient) AttachL3NetworkToVmNic(params param.AttachL3NetworkToVmNicParam) (*view.AttachL3NetworkToVmNicEventView, error) {
	resp := view.AttachL3NetworkToVmNicEventView{}
	if err := cli.Post("v1/nics/{vmNicUuid}/l3-networks/{l3NetworkUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

