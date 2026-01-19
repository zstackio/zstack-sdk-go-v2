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
	resp := view.SNSWeComEndpointInventoryView{}
	if err := cli.Post("v1/sns/application-endpoints/we-com", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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

// PageSNSWeComEndpoint Pagination
func (cli *ZSClient) PageSNSWeComEndpoint(params *param.QueryParam) ([]view.SNSWeComEndpointInventoryView, int, error) {
	var sNSWeComEndpoints []view.SNSWeComEndpointInventoryView
	total, err := cli.Page("v1/sns/application-endpoints/we-com", params, &sNSWeComEndpoints)
	return sNSWeComEndpoints, total, err
}
// UpdateSNSWeComEndpoint updates SNSWeComEndpoint
func (cli *ZSClient) UpdateSNSWeComEndpoint(uuid string, params param.UpdateSNSWeComEndpointParam) (*view.SNSApplicationEndpointInventoryView, error) {
	resp := view.SNSApplicationEndpointInventoryView{}
	if err := cli.Put("v1/sns/application-endpoints/we-com", uuid, map[string]interface{}{
		"updateSNSWeComEndpoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
