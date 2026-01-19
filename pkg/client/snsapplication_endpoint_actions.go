// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySNSApplicationEndpoint queries SNSApplicationEndpoint list
func (cli *ZSClient) QuerySNSApplicationEndpoint(params *param.QueryParam) ([]view.SNSApplicationEndpointInventoryView, error) {
	var resp []view.SNSApplicationEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints", params, &resp)
}

func (cli *ZSClient) GetSNSApplicationEndpoint(uuid string) (*view.SNSApplicationEndpointInventoryView, error) {
	var resp view.SNSApplicationEndpointInventoryView
	if err := cli.Get("v1/sns/application-endpoints", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSNSApplicationEndpoint Pagination
func (cli *ZSClient) PageSNSApplicationEndpoint(params *param.QueryParam) ([]view.SNSApplicationEndpointInventoryView, int, error) {
	var sNSApplicationEndpoints []view.SNSApplicationEndpointInventoryView
	total, err := cli.Page("v1/sns/application-endpoints", params, &sNSApplicationEndpoints)
	return sNSApplicationEndpoints, total, err
}
// UpdateSNSApplicationEndpoint updates SNSApplicationEndpoint
func (cli *ZSClient) UpdateSNSApplicationEndpoint(uuid string, params param.UpdateSNSApplicationEndpointParam) (*view.SNSApplicationEndpointInventoryView, error) {
	resp := view.SNSApplicationEndpointInventoryView{}
	if err := cli.Put("v1/sns/application-endpoints", uuid, map[string]interface{}{
		"updateSNSApplicationEndpoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteSNSApplicationEndpoint deletes SNSApplicationEndpoint
func (cli *ZSClient) DeleteSNSApplicationEndpoint(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/application-endpoints", uuid, string(deleteMode))
}
