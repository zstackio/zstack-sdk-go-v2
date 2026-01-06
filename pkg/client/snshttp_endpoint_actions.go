// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSNSHttpEndpoint creates SNSHttpEndpoint
func (cli *ZSClient) CreateSNSHttpEndpoint(params param.CreateSNSHttpEndpointParam) (*view.SNSHttpEndpointInventoryView, error) {
	var resp view.CreateSNSHttpEndpointEventView
	if err := cli.Post("v1/sns/application-endpoints/http", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateSNSHttpEndpoint updates SNSHttpEndpoint
func (cli *ZSClient) UpdateSNSHttpEndpoint(uuid string, params param.UpdateSNSHttpEndpointParam) (*view.SNSApplicationEndpointInventoryView, error) {
	var resp view.UpdateSNSApplicationEndpointEventView
	if err := cli.Put("v1/sns/application-endpoints/http/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QuerySNSHttpEndpoint queries SNSHttpEndpoint list
func (cli *ZSClient) QuerySNSHttpEndpoint(params *param.QueryParam) ([]view.SNSHttpEndpointInventoryView, error) {
	var resp []view.SNSHttpEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints/http", params, &resp)
}
