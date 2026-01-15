// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySNSUniversalSmsEndpoint queries SNSUniversalSmsEndpoint list
func (cli *ZSClient) QuerySNSUniversalSmsEndpoint(params *param.QueryParam) ([]view.SNSUniversalSmsEndpointInventoryView, error) {
	var resp []view.SNSUniversalSmsEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints/universal-sms", params, &resp)
}

// PageSNSUniversalSmsEndpoint Pagination
func (cli *ZSClient) PageSNSUniversalSmsEndpoint(params *param.QueryParam) ([]view.SNSUniversalSmsEndpointInventoryView, int, error) {
	var sNSUniversalSmsEndpoints []view.SNSUniversalSmsEndpointInventoryView
	total, err := cli.Page("v1/sns/application-endpoints/universal-sms", params, &sNSUniversalSmsEndpoints)
	return sNSUniversalSmsEndpoints, total, err
}
// UpdateSNSUniversalSmsEndpoint updates SNSUniversalSmsEndpoint
func (cli *ZSClient) UpdateSNSUniversalSmsEndpoint(uuid string, params param.UpdateSNSUniversalSmsEndpointParam) (*view.SNSApplicationEndpointInventoryView, error) {
	resp := view.SNSApplicationEndpointInventoryView{}
	if err := cli.Put("v1/sns/application-endpoints/universal-sms", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateSNSUniversalSmsEndpoint creates SNSUniversalSmsEndpoint
func (cli *ZSClient) CreateSNSUniversalSmsEndpoint(params param.CreateSNSUniversalSmsEndpointParam) (*view.SNSUniversalSmsEndpointInventoryView, error) {
	resp := view.SNSUniversalSmsEndpointInventoryView{}
	if err := cli.Post("v1/sns/application-endpoints/universal-sms", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ValidateSNSUniversalSmsEndpoint operates on SNSUniversalSmsEndpoint
func (cli *ZSClient) ValidateSNSUniversalSmsEndpoint(uuid string, params param.ValidateSNSUniversalSmsEndpointParam) (*view.SNSUniversalSmsEndpointInventoryView, error) {
	resp := view.SNSUniversalSmsEndpointInventoryView{}
	if err := cli.Put("v1/sns/application-endpoints/universal-sms", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
