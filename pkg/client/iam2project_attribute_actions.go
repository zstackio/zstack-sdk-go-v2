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

// QueryIAM2ProjectAttribute queries IAM2ProjectAttribute list
func (cli *ZSClient) QueryIAM2ProjectAttribute(params *param.QueryParam) ([]view.IAM2ProjectAttributeInventoryView, error) {
	var resp []view.IAM2ProjectAttributeInventoryView
	return resp, cli.List("v1/iam2/projects/attributes", params, &resp)
}

func (cli *ZSClient) GetIAM2ProjectAttribute(uuid string) (*view.IAM2ProjectAttributeInventoryView, error) {
	var resp view.IAM2ProjectAttributeInventoryView
	if err := cli.Get("v1/iam2/projects/attributes", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIAM2ProjectAttribute Pagination
func (cli *ZSClient) PageIAM2ProjectAttribute(params *param.QueryParam) ([]view.IAM2ProjectAttributeInventoryView, int, error) {
	var iAM2ProjectAttributes []view.IAM2ProjectAttributeInventoryView
	total, err := cli.Page("v1/iam2/projects/attributes", params, &iAM2ProjectAttributes)
	return iAM2ProjectAttributes, total, err
}
// UpdateIAM2ProjectAttribute updates IAM2ProjectAttribute
func (cli *ZSClient) UpdateIAM2ProjectAttribute(uuid string, params param.UpdateIAM2ProjectAttributeParam) (*view.IAM2ProjectAttributeInventoryView, error) {
	resp := view.IAM2ProjectAttributeInventoryView{}
	if err := cli.PutWithRespKey("v1/iam2/projects/attributes", uuid, "", map[string]interface{}{
		"updateIAM2ProjectAttribute": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddAttributesToIAM2Project adds AttributesToIAM2Project
func (cli *ZSClient) AddAttributesToIAM2Project(uuid string, params param.AddAttributesToIAM2ProjectParam) (*view.AddAttributesToIAM2ProjectEventView, error) {
	resp := view.AddAttributesToIAM2ProjectEventView{}
	if err := cli.Post(fmt.Sprintf("v1/iam2/projects/%s/attributes", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveAttributesFromIAM2Project requires additional arguments.
//
// Deprecated: use RemoveAttributesFromIAM2ProjectWithParams to supply the required fields.
func (cli *ZSClient) RemoveAttributesFromIAM2Project(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("RemoveAttributesFromIAM2Project requires additional arguments; use RemoveAttributesFromIAM2ProjectWithParams")
}

// RemoveAttributesFromIAM2ProjectWithParams sends the DELETE request with its operation parameters.
func (cli *ZSClient) RemoveAttributesFromIAM2ProjectWithParams(uuid string, deleteMode param.DeleteMode, params param.RemoveAttributesFromIAM2ProjectParamDetail) error {
	query := url.Values{
		"attributeUuids": params.AttributeUuids,
	}
	return cli.deleteActionWithQuery("v1/iam2/projects", uuid, "attributes", deleteMode, query, "attributeUuids")
}
