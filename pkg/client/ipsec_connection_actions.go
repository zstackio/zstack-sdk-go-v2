// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateIPsecConnection 更新IPsecConnection
func (cli *ZSClient) UpdateIPsecConnection(uuid string, params param.UpdateIPsecConnectionParam) (*view.UpdateIPsecConnectionEventView, error) {
	resp := view.UpdateIPsecConnectionEventView{}
	if err := cli.Put("v1/ipsec/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

