// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSNSEmailEndpoint creates SNSEmailEndpoint
func (cli *ZSClient) CreateSNSEmailEndpoint(params param.CreateSNSEmailEndpointParam) (*view.SNSApplicationEndpointInventoryView, error) {
	var resp view.CreateSNSApplicationEndpointEventView
	if err := cli.Post("v1/sns/application-endpoints/emails", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QuerySNSEmailEndpoint queries SNSEmailEndpoint list
func (cli *ZSClient) QuerySNSEmailEndpoint(params *param.QueryParam) ([]view.SNSEmailEndpointInventoryView, error) {
	var resp []view.SNSEmailEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints/emails", params, &resp)
}
