// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryThirdpartyPlatform queries ThirdpartyPlatform list
func (cli *ZSClient) QueryThirdpartyPlatform(params *param.QueryParam) ([]view.ThirdpartyPlatformInventoryView, error) {
	var resp []view.ThirdpartyPlatformInventoryView
	return resp, cli.List("v1/zwatch/third-party/platforms", params, &resp)
}

func (cli *ZSClient) GetThirdpartyPlatform(uuid string) (*view.ThirdpartyPlatformInventoryView, error) {
	var resp view.ThirdpartyPlatformInventoryView
	if err := cli.Get("v1/zwatch/third-party/platforms", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageThirdpartyPlatform Pagination
func (cli *ZSClient) PageThirdpartyPlatform(params *param.QueryParam) ([]view.ThirdpartyPlatformInventoryView, int, error) {
	var thirdpartyPlatforms []view.ThirdpartyPlatformInventoryView
	total, err := cli.Page("v1/zwatch/third-party/platforms", params, &thirdpartyPlatforms)
	return thirdpartyPlatforms, total, err
}
// UpdateThirdpartyPlatform updates ThirdpartyPlatform
func (cli *ZSClient) UpdateThirdpartyPlatform(uuid string, params param.UpdateThirdpartyPlatformParam) (*view.ThirdpartyPlatformInventoryView, error) {
	resp := view.ThirdpartyPlatformInventoryView{}
	if err := cli.PutWithRespKey("v1/zwatch/third-party/platforms", uuid, "", map[string]interface{}{
		"updateThirdpartyPlatform": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddThirdpartyPlatform adds ThirdpartyPlatform
func (cli *ZSClient) AddThirdpartyPlatform(params param.AddThirdpartyPlatformParam) (*view.ThirdpartyPlatformInventoryView, error) {
	resp := view.ThirdpartyPlatformInventoryView{}
	if err := cli.Post("v1/zwatch/third-party/platforms", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteThirdpartyPlatform deletes ThirdpartyPlatform
func (cli *ZSClient) DeleteThirdpartyPlatform(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/third-party/platforms", uuid, string(deleteMode))
}
