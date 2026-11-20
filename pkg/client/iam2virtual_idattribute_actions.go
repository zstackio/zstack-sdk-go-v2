// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"

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

// AddAttributesToIAM2VirtualID adds AttributesToIAM2VirtualID
func (cli *ZSClient) AddAttributesToIAM2VirtualID(uuid string, params param.AddAttributesToIAM2VirtualIDParam) (*view.AddAttributesToIAM2VirtualIDEventView, error) {
	resp := view.AddAttributesToIAM2VirtualIDEventView{}
	if err := cli.Post(fmt.Sprintf("v1/iam2/virtual-ids/%s/attributes", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveAttributesFromIAM2VirtualID requires additional arguments.
//
// Deprecated: use RemoveAttributesFromIAM2VirtualIDWithParams to supply the required fields.
func (cli *ZSClient) RemoveAttributesFromIAM2VirtualID(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("RemoveAttributesFromIAM2VirtualID requires additional arguments; use RemoveAttributesFromIAM2VirtualIDWithParams")
}

// RemoveAttributesFromIAM2VirtualIDWithParams sends the DELETE request with its operation parameters.
func (cli *ZSClient) RemoveAttributesFromIAM2VirtualIDWithParams(uuid string, deleteMode param.DeleteMode, params param.RemoveAttributesFromIAM2VirtualIDParamDetail) error {
	query := url.Values{
		"attributeUuids": params.AttributeUuids,
	}
	return cli.deleteActionWithQuery("v1/iam2/virtual-ids", uuid, "attributes", deleteMode, query, "attributeUuids")
}
