// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSNSHttpEndpoint creates SNSHttpEndpoint
func (cli *ZSClient) CreateSNSHttpEndpoint(params param.CreateSNSHttpEndpointParam) (*view.SNSHttpEndpointInventoryView, error) {
	resp := view.SNSHttpEndpointInventoryView{}
	if err := cli.Post("v1/sns/application-endpoints/http", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateSNSHttpEndpoint updates SNSHttpEndpoint
func (cli *ZSClient) UpdateSNSHttpEndpoint(uuid string, params param.UpdateSNSHttpEndpointParam) (*view.SNSApplicationEndpointInventoryView, error) {
	resp := view.SNSApplicationEndpointInventoryView{}
	if err := cli.Put("v1/sns/application-endpoints/http", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySNSHttpEndpoint queries SNSHttpEndpoint list
func (cli *ZSClient) QuerySNSHttpEndpoint(params *param.QueryParam) ([]view.SNSHttpEndpointInventoryView, error) {
	var resp []view.SNSHttpEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints/http", params, &resp)
}

func (cli *ZSClient) GetSNSHttpEndpoint(uuid string) (*view.SNSHttpEndpointInventoryView, error) {
	var resp view.SNSHttpEndpointInventoryView
	if err := cli.Get("v1/sns/application-endpoints/http", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSHttpEndpoint Pagination
func (cli *ZSClient) PageSNSHttpEndpoint(params *param.QueryParam) ([]view.SNSHttpEndpointInventoryView, int, error) {
	var sNSHttpEndpoints []view.SNSHttpEndpointInventoryView
	total, err := cli.Page("v1/sns/application-endpoints/http", params, &sNSHttpEndpoints)
	return sNSHttpEndpoints, total, err
}
