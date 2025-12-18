// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryVmNic queries VmNic list
func (cli *ZSClient) QueryVmNic(params param.QueryParam) ([]view.VmNicInventoryView, error) {
	var resp []view.VmNicInventoryView
	return resp, cli.List("v1/vm-instances/nics", &params, &resp)
}
