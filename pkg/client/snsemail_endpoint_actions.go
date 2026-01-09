// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
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

func (cli *ZSClient) GetSNSEmailEndpoint(uuid string) (*view.SNSEmailEndpointInventoryView, error) {
	var resp view.SNSEmailEndpointInventoryView
	if err := cli.Get("v1/sns/application-endpoints/emails", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
