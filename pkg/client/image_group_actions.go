// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ExpungeImageGroup operates on ImageGroup
func (cli *ZSClient) ExpungeImageGroup(uuid string) error {
	params := map[string]interface{}{
		"expungeImageGroup": map[string]interface{}{},
	}
	return cli.Put("v1/imagegroups", uuid, params, nil)
}
// QueryImageGroup queries ImageGroup list
func (cli *ZSClient) QueryImageGroup(params *param.QueryParam) ([]view.ImageGroupInventoryView, error) {
	var resp []view.ImageGroupInventoryView
	return resp, cli.List("v1/imagegroups", params, &resp)
}

func (cli *ZSClient) GetImageGroup(uuid string) (*view.ImageGroupInventoryView, error) {
	var resp view.ImageGroupInventoryView
	if err := cli.Get("v1/imagegroups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageImageGroup Pagination
func (cli *ZSClient) PageImageGroup(params *param.QueryParam) ([]view.ImageGroupInventoryView, int, error) {
	var imageGroups []view.ImageGroupInventoryView
	total, err := cli.Page("v1/imagegroups", params, &imageGroups)
	return imageGroups, total, err
}

// CreateImageGroupFromVmInstance creates ImageGroupFromVmInstance
func (cli *ZSClient) CreateImageGroupFromVmInstance(params param.CreateImageGroupFromVmInstanceParam) (*view.ImageGroupInventoryView, error) {
	resp := view.ImageGroupInventoryView{}
	if err := cli.Post("v1/images/groups/from/vm-instance", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateImageGroupFromImage creates ImageGroupFromImage
func (cli *ZSClient) CreateImageGroupFromImage(rootVolumeTemplateUuid string, params param.CreateImageGroupFromImageParam) (*view.ImageGroupInventoryView, error) {
	resp := view.ImageGroupInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/imagegroup/from/image/%s", rootVolumeTemplateUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateImageGroupFromSnapshot creates ImageGroupFromSnapshot
func (cli *ZSClient) CreateImageGroupFromSnapshot(rootVolumeSnapshotUuid string, params param.CreateImageGroupFromSnapshotParam) (*view.ImageGroupInventoryView, error) {
	resp := view.ImageGroupInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/imagegroup/from/snapshot/%s", rootVolumeSnapshotUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
