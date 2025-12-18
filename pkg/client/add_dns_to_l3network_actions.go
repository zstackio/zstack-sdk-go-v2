// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddDnsToL3Network adds DnsToL3Network
func (cli *ZSClient) AddDnsToL3Network(params param.AddDnsToL3NetworkParam) (*view.AddDnsToL3NetworkEventView, error) {
	resp := view.AddDnsToL3NetworkEventView{}
	if err := cli.Post("v1/l3-networks/{l3NetworkUuid}/dns", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
