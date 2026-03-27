// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateIAM2Project creates IAM2Project
func (cli *ZSClient) CreateIAM2Project(ctx context.Context, params param.CreateIAM2ProjectParam) (*view.IAM2ProjectInventoryView, error) {
	resp := view.IAM2ProjectInventoryView{}
	if err := cli.Post(ctx, "v1/iam2/projects", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteIAM2Project deletes IAM2Project
func (cli *ZSClient) DeleteIAM2Project(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/iam2/projects", uuid, string(deleteMode))
}
// RecoverIAM2Project operates on IAM2Project
func (cli *ZSClient) RecoverIAM2Project(ctx context.Context, uuid string, params param.RecoverIAM2ProjectParam) (*view.IAM2ProjectInventoryView, error) {
	resp := view.IAM2ProjectInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/iam2/projects", uuid, "", map[string]interface{}{
		"recoverIAM2Project": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryIAM2Project queries IAM2Project list
func (cli *ZSClient) QueryIAM2Project(ctx context.Context, params *param.QueryParam) ([]view.IAM2ProjectInventoryView, error) {
	var resp []view.IAM2ProjectInventoryView
	return resp, cli.List(ctx, "v1/iam2/projects", params, &resp)
}

func (cli *ZSClient) GetIAM2Project(ctx context.Context, uuid string) (*view.IAM2ProjectInventoryView, error) {
	var resp view.IAM2ProjectInventoryView
	if err := cli.Get(ctx, "v1/iam2/projects", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIAM2Project Pagination
func (cli *ZSClient) PageIAM2Project(ctx context.Context, params *param.QueryParam) ([]view.IAM2ProjectInventoryView, int, error) {
	var iAM2Projects []view.IAM2ProjectInventoryView
	total, err := cli.Page(ctx, "v1/iam2/projects", params, &iAM2Projects)
	return iAM2Projects, total, err
}
// ExpungeIAM2Project operates on IAM2Project
func (cli *ZSClient) ExpungeIAM2Project(ctx context.Context, uuid string) error {
	params := map[string]interface{}{
		"expungeIAM2Project": map[string]interface{}{},
	}
	return cli.Put(ctx, "v1/iam2/projects", uuid, params, nil)
}
// LoginIAM2Project operates on IAM2Project
func (cli *ZSClient) LoginIAM2Project(ctx context.Context, params param.LoginIAM2ProjectParam) (*view.SessionInventoryView, error) {
	resp := view.SessionInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/iam2/projects/login", "", "", map[string]interface{}{
		"loginIAM2Project": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateIAM2Project updates IAM2Project
func (cli *ZSClient) UpdateIAM2Project(ctx context.Context, uuid string, params param.UpdateIAM2ProjectParam) (*view.IAM2ProjectInventoryView, error) {
	resp := view.IAM2ProjectInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/iam2/projects", uuid, "", map[string]interface{}{
		"updateIAM2Project": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
