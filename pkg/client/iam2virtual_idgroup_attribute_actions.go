// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateIAM2VirtualIDGroupAttribute updates IAM2VirtualIDGroupAttribute
func (cli *ZSClient) UpdateIAM2VirtualIDGroupAttribute(uuid string, params param.UpdateIAM2VirtualIDGroupAttributeParam) (*view.IAM2VirtualIDGroupAttributeInventoryView, error) {
	var resp view.UpdateIAM2VirtualIDGroupAttributeEventView
	if err := cli.Put("v1/iam2/projects/groups/attributes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryIAM2VirtualIDGroupAttribute queries IAM2VirtualIDGroupAttribute list
func (cli *ZSClient) QueryIAM2VirtualIDGroupAttribute(params *param.QueryParam) ([]view.IAM2VirtualIDGroupAttributeInventoryView, error) {
	var resp []view.IAM2VirtualIDGroupAttributeInventoryView
	return resp, cli.List("v1/iam2/projects/groups/attributes/", params, &resp)
}

func (cli *ZSClient) GetIAM2VirtualIDGroupAttribute(uuid string) (*view.IAM2VirtualIDGroupAttributeInventoryView, error) {
	var resp view.IAM2VirtualIDGroupAttributeInventoryView
	if err := cli.Get("v1/iam2/projects/groups/attributes/", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
