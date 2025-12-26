// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryAppBuildSystem queries AppBuildSystem list
func (cli *ZSClient) QueryAppBuildSystem(params *param.QueryParam) ([]view.AppBuildSystemInventoryView, error) {
	var resp []view.AppBuildSystemInventoryView
	return resp, cli.List("v1/appcenter/buildsystem", params, &resp)
}
