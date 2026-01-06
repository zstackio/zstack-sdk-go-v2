// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySNSUniversalSmsEndpoint queries SNSUniversalSmsEndpoint list
func (cli *ZSClient) QuerySNSUniversalSmsEndpoint(params *param.QueryParam) ([]view.SNSUniversalSmsEndpointInventoryView, error) {
	var resp []view.SNSUniversalSmsEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints/universal-sms", params, &resp)
}
// UpdateSNSUniversalSmsEndpoint updates SNSUniversalSmsEndpoint
func (cli *ZSClient) UpdateSNSUniversalSmsEndpoint(uuid string, params param.UpdateSNSUniversalSmsEndpointParam) (*view.SNSApplicationEndpointInventoryView, error) {
	var resp view.UpdateSNSApplicationEndpointEventView
	if err := cli.Put("v1/sns/application-endpoints/universal-sms/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreateSNSUniversalSmsEndpoint creates SNSUniversalSmsEndpoint
func (cli *ZSClient) CreateSNSUniversalSmsEndpoint(params param.CreateSNSUniversalSmsEndpointParam) (*view.SNSUniversalSmsEndpointInventoryView, error) {
	var resp view.CreateSNSUniversalSmsEndpointEventView
	if err := cli.Post("v1/sns/application-endpoints/universal-sms", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// ValidateSNSUniversalSmsEndpoint operates on SNSUniversalSmsEndpoint
func (cli *ZSClient) ValidateSNSUniversalSmsEndpoint(uuid string, params param.ValidateSNSUniversalSmsEndpointParam) (*view.SNSUniversalSmsEndpointInventoryView, error) {
	resp := view.SNSUniversalSmsEndpointInventoryView{}
	if err := cli.Put("v1/sns/application-endpoints/universal-sms/{uuid}/validate", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
