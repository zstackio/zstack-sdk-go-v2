// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySharedResource queries SharedResource list
func (cli *ZSClient) QuerySharedResource(params *param.QueryParam) ([]view.SharedResourceInventoryView, error) {
	var resp []view.SharedResourceInventoryView
	return resp, cli.List("v1/accounts/resources", params, &resp)
}

// PageSharedResource Pagination
func (cli *ZSClient) PageSharedResource(params *param.QueryParam) ([]view.SharedResourceInventoryView, int, error) {
	var sharedResources []view.SharedResourceInventoryView
	total, err := cli.Page("v1/accounts/resources", params, &sharedResources)
	return sharedResources, total, err
}
