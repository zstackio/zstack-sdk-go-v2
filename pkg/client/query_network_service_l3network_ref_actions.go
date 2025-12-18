// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryNetworkServiceL3NetworkRef queries NetworkServiceL3NetworkRef list
func (cli *ZSClient) QueryNetworkServiceL3NetworkRef(params param.QueryParam) ([]view.NetworkServiceL3NetworkRefInventoryView, error) {
	var resp []view.NetworkServiceL3NetworkRefInventoryView
	return resp, cli.List("v1/l3-networks/network-services/refs", &params, &resp)
}
