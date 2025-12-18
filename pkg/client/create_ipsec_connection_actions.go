// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateIPsecConnection creates IPsecConnection
func (cli *ZSClient) CreateIPsecConnection(params param.CreateIPsecConnectionParam) (*view.CreateIPsecConnectionEventView, error) {
	resp := view.CreateIPsecConnectionEventView{}
	if err := cli.Post("v1/ipsec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
