// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateIAM2Project creates IAM2Project
func (cli *ZSClient) CreateIAM2Project(params param.CreateIAM2ProjectParam) (*view.IAM2ProjectInventoryView, error) {
	var resp view.CreateIAM2ProjectEventView
	if err := cli.Post("v1/iam2/projects", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteIAM2Project deletes IAM2Project
func (cli *ZSClient) DeleteIAM2Project(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/{uuid}", uuid, string(deleteMode))
}
// RecoverIAM2Project operates on IAM2Project
func (cli *ZSClient) RecoverIAM2Project(uuid string, params param.RecoverIAM2ProjectParam) (*view.IAM2ProjectInventoryView, error) {
	var resp view.RecoverIAM2ProjectEventView
	if err := cli.Put("v1/iam2/projects/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryIAM2Project queries IAM2Project list
func (cli *ZSClient) QueryIAM2Project(params *param.QueryParam) ([]view.IAM2ProjectInventoryView, error) {
	var resp []view.IAM2ProjectInventoryView
	return resp, cli.List("v1/iam2/projects", params, &resp)
}
// ExpungeIAM2Project operates on IAM2Project
func (cli *ZSClient) ExpungeIAM2Project(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects/{uuid}/actions", uuid, string(deleteMode))
}
// LoginIAM2Project operates on IAM2Project
func (cli *ZSClient) LoginIAM2Project(uuid string, params param.LoginIAM2ProjectParam) (*view.SessionInventoryView, error) {
	var resp view.LoginIAM2ProjectView
	if err := cli.Put("v1/iam2/projects/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateIAM2Project updates IAM2Project
func (cli *ZSClient) UpdateIAM2Project(uuid string, params param.UpdateIAM2ProjectParam) (*view.IAM2ProjectInventoryView, error) {
	var resp view.UpdateIAM2ProjectEventView
	if err := cli.Put("v1/iam2/projects/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
