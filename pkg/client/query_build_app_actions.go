// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryBuildApp queries BuildApp list
func (cli *ZSClient) QueryBuildApp(params param.QueryParam) ([]view.BuildApplicationInventoryView, error) {
	var resp []view.BuildApplicationInventoryView
	return resp, cli.List("v1/appcenter/buildapp", &params, &resp)
}
