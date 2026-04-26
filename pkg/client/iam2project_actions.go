// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{}    // avoid unused import

// CreateIAM2Project creates IAM2Project
func (cli *ZSClient) CreateIAM2Project(params param.CreateIAM2ProjectParam) (*view.IAM2ProjectInventoryView, error) {
	resp := view.IAM2ProjectInventoryView{}
	if err := cli.Post("v1/iam2/projects", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteIAM2Project deletes IAM2Project
func (cli *ZSClient) DeleteIAM2Project(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/projects", uuid, string(deleteMode))
}

// DeleteAndExpungeIAM2Project deletes IAM2Project and then expunges it from recycle bin.
func (cli *ZSClient) DeleteAndExpungeIAM2Project(uuid string, deleteMode param.DeleteMode) error {
	if err := cli.DeleteIAM2Project(uuid, deleteMode); err != nil {
		return err
	}
	return cli.ExpungeIAM2Project(uuid)
}

// RecoverIAM2Project operates on IAM2Project
func (cli *ZSClient) RecoverIAM2Project(uuid string, params param.RecoverIAM2ProjectParam) (*view.IAM2ProjectInventoryView, error) {
	resp := view.IAM2ProjectInventoryView{}
	if err := cli.PutWithRespKey("v1/iam2/projects", uuid, "", map[string]interface{}{
		"recoverIAM2Project": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryIAM2Project queries IAM2Project list
func (cli *ZSClient) QueryIAM2Project(params *param.QueryParam) ([]view.IAM2ProjectInventoryView, error) {
	var resp []view.IAM2ProjectInventoryView
	return resp, cli.List("v1/iam2/projects", params, &resp)
}

func (cli *ZSClient) GetIAM2Project(uuid string) (*view.IAM2ProjectInventoryView, error) {
	var resp view.IAM2ProjectInventoryView
	if err := cli.Get("v1/iam2/projects", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIAM2Project Pagination
func (cli *ZSClient) PageIAM2Project(params *param.QueryParam) ([]view.IAM2ProjectInventoryView, int, error) {
	var iAM2Projects []view.IAM2ProjectInventoryView
	total, err := cli.Page("v1/iam2/projects", params, &iAM2Projects)
	return iAM2Projects, total, err
}

// ExpungeIAM2Project operates on IAM2Project
func (cli *ZSClient) ExpungeIAM2Project(uuid string) error {
	params := map[string]interface{}{
		"expungeIAM2Project": map[string]interface{}{},
	}
	return cli.Put("v1/iam2/projects", uuid, params, nil)
}

// LoginIAM2Project operates on IAM2Project
func (cli *ZSClient) LoginIAM2Project(params param.LoginIAM2ProjectParam) (*view.SessionInventoryView, error) {
	resp := view.SessionInventoryView{}
	if err := cli.PutWithRespKey("v1/iam2/projects/login", "", "", map[string]interface{}{
		"loginIAM2Project": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateIAM2Project updates IAM2Project
func (cli *ZSClient) UpdateIAM2Project(uuid string, params param.UpdateIAM2ProjectParam) (*view.IAM2ProjectInventoryView, error) {
	resp := view.IAM2ProjectInventoryView{}
	if err := cli.PutWithRespKey("v1/iam2/projects", uuid, "", map[string]interface{}{
		"updateIAM2Project": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
