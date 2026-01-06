// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateSNSFeiShuEndpoint updates SNSFeiShuEndpoint
func (cli *ZSClient) UpdateSNSFeiShuEndpoint(uuid string, params param.UpdateSNSFeiShuEndpointParam) (*view.SNSApplicationEndpointInventoryView, error) {
	var resp view.UpdateSNSApplicationEndpointEventView
	if err := cli.Put("v1/sns/application-endpoints/feishu/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreateSNSFeiShuEndpoint creates SNSFeiShuEndpoint
func (cli *ZSClient) CreateSNSFeiShuEndpoint(params param.CreateSNSFeiShuEndpointParam) (*view.SNSFeiShuEndpointInventoryView, error) {
	var resp view.CreateSNSFeiShuEndpointEventView
	if err := cli.Post("v1/sns/application-endpoints/feishu", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QuerySNSFeiShuEndpoint queries SNSFeiShuEndpoint list
func (cli *ZSClient) QuerySNSFeiShuEndpoint(params *param.QueryParam) ([]view.SNSFeiShuEndpointInventoryView, error) {
	var resp []view.SNSFeiShuEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints/feishu", params, &resp)
}
