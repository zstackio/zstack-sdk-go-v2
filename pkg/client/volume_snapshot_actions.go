// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateVolumeSnapshot updates VolumeSnapshot
func (cli *ZSClient) UpdateVolumeSnapshot(uuid string, params param.UpdateVolumeSnapshotParam) (*view.VolumeSnapshotInventoryView, error) {
	resp := view.VolumeSnapshotInventoryView{}
	if err := cli.Put("v1/volume-snapshots", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVolumeSnapshot queries VolumeSnapshot list
func (cli *ZSClient) QueryVolumeSnapshot(params *param.QueryParam) ([]view.VolumeSnapshotInventoryView, error) {
	var resp []view.VolumeSnapshotInventoryView
	return resp, cli.List("v1/volume-snapshots", params, &resp)
}

func (cli *ZSClient) GetVolumeSnapshot(uuid string) (*view.VolumeSnapshotInventoryView, error) {
	var resp view.VolumeSnapshotInventoryView
	if err := cli.Get("v1/volume-snapshots", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVolumeSnapshot Pagination
func (cli *ZSClient) PageVolumeSnapshot(params *param.QueryParam) ([]view.VolumeSnapshotInventoryView, int, error) {
	var volumeSnapshots []view.VolumeSnapshotInventoryView
	total, err := cli.Page("v1/volume-snapshots", params, &volumeSnapshots)
	return volumeSnapshots, total, err
}
// DeleteVolumeSnapshot deletes VolumeSnapshot
func (cli *ZSClient) DeleteVolumeSnapshot(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volume-snapshots", uuid, string(deleteMode))
}
// CreateVolumeSnapshot creates VolumeSnapshot
func (cli *ZSClient) CreateVolumeSnapshot(params param.CreateVolumeSnapshotParam) (*view.VolumeSnapshotInventoryView, error) {
	resp := view.VolumeSnapshotInventoryView{}
	if err := cli.Post("v1/volumes/{volumeUuid}/volume-snapshots", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
