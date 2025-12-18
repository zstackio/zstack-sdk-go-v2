// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryLogServer queries LogServer list
func (cli *ZSClient) QueryLogServer(params param.QueryParam) ([]view.LogServerInventoryView, error) {
	var resp []view.LogServerInventoryView
	return resp, cli.List("v1/log/servers", &params, &resp)
}
