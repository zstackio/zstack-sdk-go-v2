// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryApplicationDevelopmentService queries ApplicationDevelopmentService list
func (cli *ZSClient) QueryApplicationDevelopmentService(params *param.QueryParam) ([]view.ApplicationDevelopmentServiceInventoryView, error) {
	var resp []view.ApplicationDevelopmentServiceInventoryView
	return resp, cli.List("v1/ai/model-services/app/", params, &resp)
}
