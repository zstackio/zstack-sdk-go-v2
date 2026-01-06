// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSNSDingTalkEndpoint creates SNSDingTalkEndpoint
func (cli *ZSClient) CreateSNSDingTalkEndpoint(params param.CreateSNSDingTalkEndpointParam) (*view.SNSDingTalkEndpointInventoryView, error) {
	var resp view.CreateSNSDingTalkEndpointEventView
	if err := cli.Post("v1/sns/application-endpoints/ding-talk", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QuerySNSDingTalkEndpoint queries SNSDingTalkEndpoint list
func (cli *ZSClient) QuerySNSDingTalkEndpoint(params *param.QueryParam) ([]view.SNSDingTalkEndpointInventoryView, error) {
	var resp []view.SNSDingTalkEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints/ding-talk", params, &resp)
}
// UpdateSNSDingTalkEndpoint updates SNSDingTalkEndpoint
func (cli *ZSClient) UpdateSNSDingTalkEndpoint(uuid string, params param.UpdateSNSDingTalkEndpointParam) (*view.SNSApplicationEndpointInventoryView, error) {
	var resp view.UpdateSNSApplicationEndpointEventView
	if err := cli.Put("v1/sns/application-endpoints/ding-talk/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
