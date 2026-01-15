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

// PageVolumeSnapshotGroup Pagination
func (cli *ZSClient) PageVolumeSnapshotGroup(params *param.QueryParam) ([]view.VolumeSnapshotGroupInventoryView, int, error) {
	var volumeSnapshotGroups []view.VolumeSnapshotGroupInventoryView
	total, err := cli.Page("v1/volume-snapshots/group", params, &volumeSnapshotGroups)
	return volumeSnapshotGroups, total, err
}
// CreateVolumeSnapshotGroup creates VolumeSnapshotGroup
func (cli *ZSClient) CreateVolumeSnapshotGroup(params param.CreateVolumeSnapshotGroupParam) (*view.VolumeSnapshotGroupInventoryView, error) {
	resp := view.VolumeSnapshotGroupInventoryView{}
	if err := cli.Post("v1/volume-snapshots/group", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateVolumeSnapshotGroup updates VolumeSnapshotGroup
func (cli *ZSClient) UpdateVolumeSnapshotGroup(uuid string, params param.UpdateVolumeSnapshotGroupParam) (*view.VolumeSnapshotGroupInventoryView, error) {
	resp := view.VolumeSnapshotGroupInventoryView{}
	if err := cli.Put("v1/volume-snapshots/group", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteVolumeSnapshotGroup deletes VolumeSnapshotGroup
func (cli *ZSClient) DeleteVolumeSnapshotGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volume-snapshots/group", uuid, string(deleteMode))
}
