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
	resp := view.SNSApplicationEndpointInventoryView{}
	if err := cli.Post("v1/sns/application-endpoints/emails", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySNSEmailEndpoint queries SNSEmailEndpoint list
func (cli *ZSClient) QuerySNSEmailEndpoint(params *param.QueryParam) ([]view.SNSEmailEndpointInventoryView, error) {
	var resp []view.SNSEmailEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints/emails", params, &resp)
}

// PageSNSEmailEndpoint Pagination
func (cli *ZSClient) PageSNSEmailEndpoint(params *param.QueryParam) ([]view.SNSEmailEndpointInventoryView, int, error) {
	var sNSEmailEndpoints []view.SNSEmailEndpointInventoryView
	total, err := cli.Page("v1/sns/application-endpoints/emails", params, &sNSEmailEndpoints)
	return sNSEmailEndpoints, total, err
}
