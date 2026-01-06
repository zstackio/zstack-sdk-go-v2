// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySNSApplicationEndpoint queries SNSApplicationEndpoint list
func (cli *ZSClient) QuerySNSApplicationEndpoint(params *param.QueryParam) ([]view.SNSApplicationEndpointInventoryView, error) {
	var resp []view.SNSApplicationEndpointInventoryView
	return resp, cli.List("v1/sns/application-endpoints", params, &resp)
}
// UpdateSNSApplicationEndpoint updates SNSApplicationEndpoint
func (cli *ZSClient) UpdateSNSApplicationEndpoint(uuid string, params param.UpdateSNSApplicationEndpointParam) (*view.SNSApplicationEndpointInventoryView, error) {
	var resp view.UpdateSNSApplicationEndpointEventView
	if err := cli.Put("v1/sns/application-endpoints/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteSNSApplicationEndpoint deletes SNSApplicationEndpoint
func (cli *ZSClient) DeleteSNSApplicationEndpoint(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/application-endpoints/{uuid}", uuid, string(deleteMode))
}
