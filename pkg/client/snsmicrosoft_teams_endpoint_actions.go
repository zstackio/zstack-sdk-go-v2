// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateSNSMicrosoftTeamsEndpoint updates SNSMicrosoftTeamsEndpoint
func (cli *ZSClient) UpdateSNSMicrosoftTeamsEndpoint(uuid string, params param.UpdateSNSMicrosoftTeamsEndpointParam) (*view.SNSApplicationEndpointInventoryView, error) {
	var resp view.UpdateSNSApplicationEndpointEventView
	if err := cli.Put("v1/sns/application-endpoints/microsoft-teams", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreateSNSMicrosoftTeamsEndpoint creates SNSMicrosoftTeamsEndpoint
func (cli *ZSClient) CreateSNSMicrosoftTeamsEndpoint(params param.CreateSNSMicrosoftTeamsEndpointParam) (*view.SNSMicrosoftTeamsEndpointInventoryView, error) {
	var resp view.CreateSNSMicrosoftTeamsEndpointEventView
	if err := cli.Post("v1/sns/application-endpoints/microsoft-teams", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QuerySNSMicrosoftTeamsEndpoint queries SNSMicrosoftTeamsEndpoint list
func (cli *ZSClient) QuerySNSMicrosoftTeamsEndpoint(params *param.QueryParam) ([]view.SNSMicrosoftTeamsEndpointInventoryView, error) {
	var resp []view.SNSMicrosoftTeamsEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints/microsoft-teams", params, &resp)
}

func (cli *ZSClient) GetSNSMicrosoftTeamsEndpoint(uuid string) (*view.SNSMicrosoftTeamsEndpointInventoryView, error) {
	var resp view.SNSMicrosoftTeamsEndpointInventoryView
	if err := cli.Get("v1/sns/application-endpoints/microsoft-teams", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
