// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryBuildApp queries BuildApp list
func (cli *ZSClient) QueryBuildApp(params *param.QueryParam) ([]view.BuildApplicationInventoryView, error) {
	var resp []view.BuildApplicationInventoryView
	return resp, cli.List("v1/appcenter/buildapp", params, &resp)
}
