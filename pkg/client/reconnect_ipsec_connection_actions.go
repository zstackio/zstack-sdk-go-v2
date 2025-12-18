// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ReconnectIPsecConnection 操作ReconnectIPsecConnection
func (cli *ZSClient) ReconnectIPsecConnection(uuid string, params param.ReconnectIPsecConnectionParam) (*view.ReconnectIPsecConnectionEventView, error) {
	resp := view.ReconnectIPsecConnectionEventView{}
	if err := cli.Put("v1/ipsec/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

