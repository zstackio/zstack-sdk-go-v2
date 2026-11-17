// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// CreateBuildApp creates BuildApp
func (cli *ZSClient) CreateBuildApp(params param.CreateBuildAppParam) (*view.BuildApplicationInventoryView, error) {
	resp := view.BuildApplicationInventoryView{}
	if err := cli.Post("v1/appcenter/buildapp", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryBuildApp queries BuildApp list
func (cli *ZSClient) QueryBuildApp(params *param.QueryParam) ([]view.BuildApplicationInventoryView, error) {
	var resp []view.BuildApplicationInventoryView
	return resp, cli.List("v1/appcenter/buildapp", params, &resp)
}

func (cli *ZSClient) GetBuildApp(uuid string) (*view.BuildApplicationInventoryView, error) {
	var resp view.BuildApplicationInventoryView
	if err := cli.Get("v1/appcenter/buildapp", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBuildApp Pagination
func (cli *ZSClient) PageBuildApp(params *param.QueryParam) ([]view.BuildApplicationInventoryView, int, error) {
	var buildApps []view.BuildApplicationInventoryView
	total, err := cli.Page("v1/appcenter/buildapp", params, &buildApps)
	return buildApps, total, err
}

// AddBuildApp adds BuildApp
func (cli *ZSClient) AddBuildApp(params param.AddBuildAppParam) (*view.BuildApplicationInventoryView, error) {
	resp := view.BuildApplicationInventoryView{}
	if err := cli.Post("v1/appcenter/buildapp/add", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteBuildApp deletes BuildApp
func (cli *ZSClient) DeleteBuildApp(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/appcenter/buildapp", uuid, string(deleteMode))
}

// CheckBuildAppParameters operates on BuildAppParameters
func (cli *ZSClient) CheckBuildAppParameters(params param.CheckBuildAppParametersParam) (*view.CheckBuildAppParametersView, error) {
	resp := view.CheckBuildAppParametersView{}
	if err := cli.Post("v1/appcenter/buildapp/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateBuildApp updates BuildApp
func (cli *ZSClient) UpdateBuildApp(uuid string, params param.UpdateBuildAppParam) (*view.BuildApplicationInventoryView, error) {
	resp := view.BuildApplicationInventoryView{}
	if err := cli.PutWithRespKey("v1/appcenter/buildapp", uuid, "", map[string]interface{}{
		"updateBuildApp": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
