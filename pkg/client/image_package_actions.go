// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateImagePackage updates ImagePackage
func (cli *ZSClient) UpdateImagePackage(uuid string, params param.UpdateImagePackageParam) (*view.ImagePackageInventoryView, error) {
	resp := view.ImagePackageInventoryView{}
	if err := cli.Put("v1/image-packages", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryImagePackage queries ImagePackage list
func (cli *ZSClient) QueryImagePackage(params *param.QueryParam) ([]view.ImagePackageInventoryView, error) {
	var resp []view.ImagePackageInventoryView
	return resp, cli.List("v1/image-packages", params, &resp)
}

func (cli *ZSClient) GetImagePackage(uuid string) (*view.ImagePackageInventoryView, error) {
	var resp view.ImagePackageInventoryView
	if err := cli.Get("v1/image-packages", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageImagePackage Pagination
func (cli *ZSClient) PageImagePackage(params *param.QueryParam) ([]view.ImagePackageInventoryView, int, error) {
	var imagePackages []view.ImagePackageInventoryView
	total, err := cli.Page("v1/image-packages", params, &imagePackages)
	return imagePackages, total, err
}
// DeleteImagePackage deletes ImagePackage
func (cli *ZSClient) DeleteImagePackage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/image-packages", uuid, string(deleteMode))
}
