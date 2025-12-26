// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryIPSecConnection queries IPSecConnection list
func (cli *ZSClient) QueryIPSecConnection(params *param.QueryParam) ([]view.IPsecConnectionInventoryView, error) {
	var resp []view.IPsecConnectionInventoryView
	return resp, cli.List("v1/ipsec", params, &resp)
}
