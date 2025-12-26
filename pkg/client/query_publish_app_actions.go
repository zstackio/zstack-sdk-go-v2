// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryPublishApp queries PublishApp list
func (cli *ZSClient) QueryPublishApp(params *param.QueryParam) ([]view.PublishAppInventoryView, error) {
	var resp []view.PublishAppInventoryView
	return resp, cli.List("v1/appcenter/app", params, &resp)
}
