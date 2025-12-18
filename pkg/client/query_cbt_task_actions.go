// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryCbtTask queries CbtTask list
func (cli *ZSClient) QueryCbtTask(params param.QueryParam) ([]view.CbtTaskInventoryView, error) {
	var resp []view.CbtTaskInventoryView
	return resp, cli.List("v1/cbt-task", &params, &resp)
}
