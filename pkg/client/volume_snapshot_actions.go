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
	var resp view.UpdateVolumeSnapshotEventView
	err := cli.PutWithSpec("v1/volume-snapshots", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
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
// DeleteVolumeSnapshot deletes VolumeSnapshot
func (cli *ZSClient) DeleteVolumeSnapshot(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/volume-snapshots", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// CreateVolumeSnapshot creates VolumeSnapshot
func (cli *ZSClient) CreateVolumeSnapshot(params param.CreateVolumeSnapshotParam) (*view.VolumeSnapshotInventoryView, error) {
	var resp view.CreateVolumeSnapshotEventView
	if err := cli.Post("v1/volumes/{volumeUuid}/volume-snapshots", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
