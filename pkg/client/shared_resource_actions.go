// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySharedResource queries SharedResource list
func (cli *ZSClient) QuerySharedResource(params *param.QueryParam) ([]view.SharedResourceInventoryView, error) {
	var resp []view.SharedResourceInventoryView
	return resp, cli.List("v1/accounts/resources", params, &resp)
}
