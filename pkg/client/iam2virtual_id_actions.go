// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryIAM2VirtualID queries IAM2VirtualID list
func (cli *ZSClient) QueryIAM2VirtualID(params *param.QueryParam) ([]view.IAM2VirtualIDInventoryView, error) {
	var resp []view.IAM2VirtualIDInventoryView
	return resp, cli.List("v1/iam2/virtual-ids", params, &resp)
}

// PageIAM2VirtualID Pagination
func (cli *ZSClient) PageIAM2VirtualID(params *param.QueryParam) ([]view.IAM2VirtualIDInventoryView, int, error) {
	var iAM2VirtualIDs []view.IAM2VirtualIDInventoryView
	total, err := cli.Page("v1/iam2/virtual-ids", params, &iAM2VirtualIDs)
	return iAM2VirtualIDs, total, err
}
// CreateIAM2VirtualID creates IAM2VirtualID
func (cli *ZSClient) CreateIAM2VirtualID(params param.CreateIAM2VirtualIDParam) (*view.IAM2VirtualIDInventoryView, error) {
	resp := view.IAM2VirtualIDInventoryView{}
	if err := cli.Post("v1/iam2/virtual-ids", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteIAM2VirtualID deletes IAM2VirtualID
func (cli *ZSClient) DeleteIAM2VirtualID(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/virtual-ids", uuid, string(deleteMode))
}
// LoginIAM2VirtualID operates on IAM2VirtualID
func (cli *ZSClient) LoginIAM2VirtualID(uuid string, params param.LoginIAM2VirtualIDParam) (*view.SessionInventoryView, error) {
	resp := view.SessionInventoryView{}
	if err := cli.Put("v1/iam2/virtual-ids/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateIAM2VirtualID updates IAM2VirtualID
func (cli *ZSClient) UpdateIAM2VirtualID(uuid string, params param.UpdateIAM2VirtualIDParam) (*view.IAM2VirtualIDInventoryView, error) {
	resp := view.IAM2VirtualIDInventoryView{}
	if err := cli.Put("v1/iam2/virtual-ids", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
