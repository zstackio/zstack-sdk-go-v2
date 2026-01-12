// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSNSWeComEndpoint creates SNSWeComEndpoint
func (cli *ZSClient) CreateSNSWeComEndpoint(params param.CreateSNSWeComEndpointParam) (*view.SNSWeComEndpointInventoryView, error) {
	var resp view.CreateSNSWeComEndpointEventView
	if err := cli.Post("v1/sns/application-endpoints/we-com", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QuerySNSWeComEndpoint queries SNSWeComEndpoint list
func (cli *ZSClient) QuerySNSWeComEndpoint(params *param.QueryParam) ([]view.SNSWeComEndpointInventoryView, error) {
	var resp []view.SNSWeComEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints/we-com", params, &resp)
}

func (cli *ZSClient) GetSNSWeComEndpoint(uuid string) (*view.SNSWeComEndpointInventoryView, error) {
	var resp view.SNSWeComEndpointInventoryView
	if err := cli.Get("v1/sns/application-endpoints/we-com", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateSNSWeComEndpoint updates SNSWeComEndpoint
func (cli *ZSClient) UpdateSNSWeComEndpoint(uuid string, params param.UpdateSNSWeComEndpointParam) (*view.SNSApplicationEndpointInventoryView, error) {
	var resp view.UpdateSNSApplicationEndpointEventView
	err := cli.PutWithSpec("v1/sns/application-endpoints/we-com", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
