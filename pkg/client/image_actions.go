// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddImage adds Image
func (cli *ZSClient) AddImage(params param.AddImageParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.Post("v1/images", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddImageAsync Async
func (cli *ZSClient) AddImageAsync(params param.AddImageParam) (string, error) {

	resource := "v1/images"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}
// QueryImage queries Image list
func (cli *ZSClient) QueryImage(params *param.QueryParam) ([]view.ImageInventoryView, error) {
	var resp []view.ImageInventoryView
	return resp, cli.List("v1/images", params, &resp)
}

func (cli *ZSClient) GetImage(uuid string) (*view.ImageInventoryView, error) {
	var resp view.ImageInventoryView
	if err := cli.Get("v1/images", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageImage Pagination
func (cli *ZSClient) PageImage(params *param.QueryParam) ([]view.ImageInventoryView, int, error) {
	var images []view.ImageInventoryView
	total, err := cli.Page("v1/images", params, &images)
	return images, total, err
}
// SyncImage operates on Image
func (cli *ZSClient) SyncImage(imageStoreUuid string, params param.SyncImageParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithRespKey("v1/backup-storage/image-store", imageStoreUuid, "", map[string]interface{}{
		"syncImage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// RecoverImage operates on Image
func (cli *ZSClient) RecoverImage(imageUuid string, params param.RecoverImageParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithRespKey("v1/images", imageUuid, "", map[string]interface{}{
		"recoverImage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CloneImage operates on Image
func (cli *ZSClient) CloneImage(imageUuid string, params param.CloneImageParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/image/clone/%s", imageUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteImage deletes Image
func (cli *ZSClient) DeleteImage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/images", uuid, string(deleteMode))
}
// UpdateImage updates Image
func (cli *ZSClient) UpdateImage(uuid string, params param.UpdateImageParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.PutWithRespKey("v1/images", uuid, "", map[string]interface{}{
		"updateImage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ExpungeImage operates on Image
func (cli *ZSClient) ExpungeImage(uuid string) error {
	params := map[string]interface{}{
		"expungeImage": map[string]interface{}{},
	}
	return cli.Put("v1/images", uuid, params, nil)
}
