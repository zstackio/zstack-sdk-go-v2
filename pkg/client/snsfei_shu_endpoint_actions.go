// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateSNSFeiShuEndpoint updates SNSFeiShuEndpoint
func (cli *ZSClient) UpdateSNSFeiShuEndpoint(uuid string, params param.UpdateSNSFeiShuEndpointParam) (*view.SNSApplicationEndpointInventoryView, error) {
	var resp view.UpdateSNSApplicationEndpointEventView
	err := cli.PutWithSpec("v1/sns/application-endpoints/feishu", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
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

func (cli *ZSClient) GetSNSFeiShuEndpoint(uuid string) (*view.SNSFeiShuEndpointInventoryView, error) {
	var resp view.SNSFeiShuEndpointInventoryView
	if err := cli.Get("v1/sns/application-endpoints/feishu", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
