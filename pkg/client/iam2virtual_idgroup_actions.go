// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryIAM2VirtualIDGroup queries IAM2VirtualIDGroup list
func (cli *ZSClient) QueryIAM2VirtualIDGroup(params *param.QueryParam) ([]view.IAM2VirtualIDGroupInventoryView, error) {
	var resp []view.IAM2VirtualIDGroupInventoryView
	return resp, cli.List("v1/iam2/projects/groups", params, &resp)
}

func (cli *ZSClient) GetIAM2VirtualIDGroup(uuid string) (*view.IAM2VirtualIDGroupInventoryView, error) {
	var resp view.IAM2VirtualIDGroupInventoryView
	if err := cli.Get("v1/iam2/projects/groups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIAM2VirtualIDGroup Pagination
func (cli *ZSClient) PageIAM2VirtualIDGroup(params *param.QueryParam) ([]view.IAM2VirtualIDGroupInventoryView, int, error) {
	var iAM2VirtualIDGroups []view.IAM2VirtualIDGroupInventoryView
	total, err := cli.Page("v1/iam2/projects/groups", params, &iAM2VirtualIDGroups)
	return iAM2VirtualIDGroups, total, err
}
// CreateIAM2VirtualIDGroup creates IAM2VirtualIDGroup
func (cli *ZSClient) CreateIAM2VirtualIDGroup(params param.CreateIAM2VirtualIDGroupParam) (*view.IAM2VirtualIDGroupInventoryView, error) {
	resp := view.IAM2VirtualIDGroupInventoryView{}
	if err := cli.Post("v1/iam2/groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteIAM2VirtualIDGroup deletes IAM2VirtualIDGroup
func (cli *ZSClient) DeleteIAM2VirtualIDGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/groups", uuid, string(deleteMode))
}
// UpdateIAM2VirtualIDGroup updates IAM2VirtualIDGroup
func (cli *ZSClient) UpdateIAM2VirtualIDGroup(uuid string, params param.UpdateIAM2VirtualIDGroupParam) (*view.IAM2VirtualIDGroupInventoryView, error) {
	resp := view.IAM2VirtualIDGroupInventoryView{}
	if err := cli.PutWithRespKey("v1/iam2/projects/groups", uuid, "", map[string]interface{}{
		"updateIAM2VirtualIDGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
