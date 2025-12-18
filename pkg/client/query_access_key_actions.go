// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryAccessKey queries AccessKey list
func (cli *ZSClient) QueryAccessKey(params param.QueryParam) ([]view.AccessKeyInventoryView, error) {
	var resp []view.AccessKeyInventoryView
	return resp, cli.List("v1/accesskeys", &params, &resp)
}
