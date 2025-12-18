// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAccountResourceRef queries AccountResourceRef list
func (cli *ZSClient) QueryAccountResourceRef(params param.QueryParam) ([]view.AccountResourceRefInventoryView, error) {
	var resp []view.AccountResourceRefInventoryView
	return resp, cli.List("v1/accounts/resources/refs", &params, &resp)
}
