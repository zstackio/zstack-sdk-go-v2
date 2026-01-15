// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryApplicationDevelopmentService queries ApplicationDevelopmentService list
func (cli *ZSClient) QueryApplicationDevelopmentService(params *param.QueryParam) ([]view.ApplicationDevelopmentServiceInventoryView, error) {
	var resp []view.ApplicationDevelopmentServiceInventoryView
	return resp, cli.List("v1/ai/model-services/app/", params, &resp)
}

// PageApplicationDevelopmentService Pagination
func (cli *ZSClient) PageApplicationDevelopmentService(params *param.QueryParam) ([]view.ApplicationDevelopmentServiceInventoryView, int, error) {
	var applicationDevelopmentServices []view.ApplicationDevelopmentServiceInventoryView
	total, err := cli.Page("v1/ai/model-services/app/", params, &applicationDevelopmentServices)
	return applicationDevelopmentServices, total, err
}
