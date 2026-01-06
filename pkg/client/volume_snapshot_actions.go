// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateVolumeSnapshot updates VolumeSnapshot
func (cli *ZSClient) UpdateVolumeSnapshot(uuid string, params param.UpdateVolumeSnapshotParam) (*view.VolumeSnapshotInventoryView, error) {
	var resp view.UpdateVolumeSnapshotEventView
	if err := cli.Put("v1/volume-snapshots/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryVolumeSnapshot queries VolumeSnapshot list
func (cli *ZSClient) QueryVolumeSnapshot(params *param.QueryParam) ([]view.VolumeSnapshotInventoryView, error) {
	var resp []view.VolumeSnapshotInventoryView
	return resp, cli.List("v1/volume-snapshots", params, &resp)
}
// DeleteVolumeSnapshot deletes VolumeSnapshot
func (cli *ZSClient) DeleteVolumeSnapshot(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volume-snapshots/{uuid}", uuid, string(deleteMode))
}
// CreateVolumeSnapshot creates VolumeSnapshot
func (cli *ZSClient) CreateVolumeSnapshot(params param.CreateVolumeSnapshotParam) (*view.VolumeSnapshotInventoryView, error) {
	var resp view.CreateVolumeSnapshotEventView
	if err := cli.Post("v1/volumes/{volumeUuid}/volume-snapshots", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
