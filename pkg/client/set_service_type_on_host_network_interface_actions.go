// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetServiceTypeOnHostNetworkInterface 操作SetServiceTypeOnHostNetworkInterface
func (cli *ZSClient) SetServiceTypeOnHostNetworkInterface(params param.SetServiceTypeOnHostNetworkInterfaceParam) (*view.SetServiceTypeOnHostNetworkInterfaceEventView, error) {
	resp := view.SetServiceTypeOnHostNetworkInterfaceEventView{}
	if err := cli.Post("v1/hosts/nics/service-types", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

