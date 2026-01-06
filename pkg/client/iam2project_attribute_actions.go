// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryIAM2ProjectAttribute queries IAM2ProjectAttribute list
func (cli *ZSClient) QueryIAM2ProjectAttribute(params *param.QueryParam) ([]view.IAM2ProjectAttributeInventoryView, error) {
	var resp []view.IAM2ProjectAttributeInventoryView
	return resp, cli.List("v1/iam2/projects/attributes", params, &resp)
}
// UpdateIAM2ProjectAttribute updates IAM2ProjectAttribute
func (cli *ZSClient) UpdateIAM2ProjectAttribute(uuid string, params param.UpdateIAM2ProjectAttributeParam) (*view.IAM2ProjectAttributeInventoryView, error) {
	var resp view.UpdateIAM2ProjectAttributeEventView
	if err := cli.Put("v1/iam2/projects/attributes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
