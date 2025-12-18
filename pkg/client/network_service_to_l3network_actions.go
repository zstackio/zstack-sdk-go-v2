// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachNetworkServiceToL3Network 操作NetworkServiceToL3Network
func (cli *ZSClient) AttachNetworkServiceToL3Network(params param.AttachNetworkServiceToL3NetworkParam) (*view.AttachNetworkServiceToL3NetworkEventView, error) {
	resp := view.AttachNetworkServiceToL3NetworkEventView{}
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/network-services", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

