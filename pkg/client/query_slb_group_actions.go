// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QuerySlbGroup queries SlbGroup list
func (cli *ZSClient) QuerySlbGroup(params param.QueryParam) ([]view.SlbGroupInventoryView, error) {
	var resp []view.SlbGroupInventoryView
	return resp, cli.List("v1/load-balancers/slb/groups", &params, &resp)
}
