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
// UpdateSNSApplicationEndpoint updates SNSApplicationEndpoint
func (cli *ZSClient) UpdateSNSApplicationEndpoint(uuid string, params param.UpdateSNSApplicationEndpointParam) (*view.SNSApplicationEndpointInventoryView, error) {
	var resp view.UpdateSNSApplicationEndpointEventView
	err := cli.PutWithSpec("v1/sns/application-endpoints", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteSNSApplicationEndpoint deletes SNSApplicationEndpoint
func (cli *ZSClient) DeleteSNSApplicationEndpoint(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/sns/application-endpoints", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
