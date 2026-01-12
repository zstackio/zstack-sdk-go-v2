// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVolumeSnapshotGroup queries VolumeSnapshotGroup list
func (cli *ZSClient) QueryVolumeSnapshotGroup(params *param.QueryParam) ([]view.VolumeSnapshotGroupInventoryView, error) {
	var resp []view.VolumeSnapshotGroupInventoryView
	return resp, cli.List("v1/volume-snapshots/group", params, &resp)
}

func (cli *ZSClient) GetVolumeSnapshotGroup(uuid string) (*view.VolumeSnapshotGroupInventoryView, error) {
	var resp view.VolumeSnapshotGroupInventoryView
	if err := cli.Get("v1/volume-snapshots/group", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateVolumeSnapshotGroup creates VolumeSnapshotGroup
func (cli *ZSClient) CreateVolumeSnapshotGroup(params param.CreateVolumeSnapshotGroupParam) (*view.VolumeSnapshotGroupInventoryView, error) {
	var resp view.CreateVolumeSnapshotGroupEventView
	if err := cli.Post("v1/volume-snapshots/group", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateVolumeSnapshotGroup updates VolumeSnapshotGroup
func (cli *ZSClient) UpdateVolumeSnapshotGroup(uuid string, params param.UpdateVolumeSnapshotGroupParam) (*view.VolumeSnapshotGroupInventoryView, error) {
	var resp view.UpdateVolumeSnapshotGroupEventView
	err := cli.PutWithSpec("v1/volume-snapshots/group", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteVolumeSnapshotGroup deletes VolumeSnapshotGroup
func (cli *ZSClient) DeleteVolumeSnapshotGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/volume-snapshots/group", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
