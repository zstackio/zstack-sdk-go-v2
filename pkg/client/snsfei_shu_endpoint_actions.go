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
	resp := view.SNSApplicationEndpointInventoryView{}
	if err := cli.PutWithRespKey("v1/sns/application-endpoints/feishu", uuid, "", map[string]interface{}{
		"updateSNSFeiShuEndpoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateSNSFeiShuEndpoint creates SNSFeiShuEndpoint
func (cli *ZSClient) CreateSNSFeiShuEndpoint(params param.CreateSNSFeiShuEndpointParam) (*view.SNSFeiShuEndpointInventoryView, error) {
	resp := view.SNSFeiShuEndpointInventoryView{}
	if err := cli.Post("v1/sns/application-endpoints/feishu", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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

// PageSNSFeiShuEndpoint Pagination
func (cli *ZSClient) PageSNSFeiShuEndpoint(params *param.QueryParam) ([]view.SNSFeiShuEndpointInventoryView, int, error) {
	var sNSFeiShuEndpoints []view.SNSFeiShuEndpointInventoryView
	total, err := cli.Page("v1/sns/application-endpoints/feishu", params, &sNSFeiShuEndpoints)
	return sNSFeiShuEndpoints, total, err
}
