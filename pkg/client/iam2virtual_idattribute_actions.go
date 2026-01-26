// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateIAM2VirtualIDAttribute updates IAM2VirtualIDAttribute
func (cli *ZSClient) UpdateIAM2VirtualIDAttribute(uuid string, params param.UpdateIAM2VirtualIDAttributeParam) (*view.IAM2VirtualIDAttributeInventoryView, error) {
	resp := view.IAM2VirtualIDAttributeInventoryView{}
	if err := cli.PutWithRespKey("v1/iam2/virtual-ids/attributes", uuid, "", map[string]interface{}{
		"updateIAM2VirtualIDAttribute": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryIAM2VirtualIDAttribute queries IAM2VirtualIDAttribute list
func (cli *ZSClient) QueryIAM2VirtualIDAttribute(params *param.QueryParam) ([]view.IAM2VirtualIDAttributeInventoryView, error) {
	var resp []view.IAM2VirtualIDAttributeInventoryView
	return resp, cli.List("v1/iam2/virtual-ids/attributes", params, &resp)
}

func (cli *ZSClient) GetIAM2VirtualIDAttribute(uuid string) (*view.IAM2VirtualIDAttributeInventoryView, error) {
	var resp view.IAM2VirtualIDAttributeInventoryView
	if err := cli.Get("v1/iam2/virtual-ids/attributes", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIAM2VirtualIDAttribute Pagination
func (cli *ZSClient) PageIAM2VirtualIDAttribute(params *param.QueryParam) ([]view.IAM2VirtualIDAttributeInventoryView, int, error) {
	var iAM2VirtualIDAttributes []view.IAM2VirtualIDAttributeInventoryView
	total, err := cli.Page("v1/iam2/virtual-ids/attributes", params, &iAM2VirtualIDAttributes)
	return iAM2VirtualIDAttributes, total, err
}
