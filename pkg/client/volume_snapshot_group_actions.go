// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVolumeSnapshotGroup queries VolumeSnapshotGroup list
func (cli *ZSClient) QueryVolumeSnapshotGroup(params *param.QueryParam) ([]view.VolumeSnapshotGroupInventoryView, error) {
	var resp []view.VolumeSnapshotGroupInventoryView
	return resp, cli.List("v1/volume-snapshots/group", params, &resp)
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
	if err := cli.Put("v1/volume-snapshots/group/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteVolumeSnapshotGroup deletes VolumeSnapshotGroup
func (cli *ZSClient) DeleteVolumeSnapshotGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volume-snapshots/group/{uuid}", uuid, string(deleteMode))
}
