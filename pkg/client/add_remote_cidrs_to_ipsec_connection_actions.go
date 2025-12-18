// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddRemoteCidrsToIPsecConnection 操作AddRemoteCidrsToIPsecConnection
func (cli *ZSClient) AddRemoteCidrsToIPsecConnection(params param.AddRemoteCidrsToIPsecConnectionParam) (*view.AddRemoteCidrsToIPsecConnectionEventView, error) {
	resp := view.AddRemoteCidrsToIPsecConnectionEventView{}
	if err := cli.Post("v1/ipsec/{uuid}/remote-cidrs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

