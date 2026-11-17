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
	if err := cli.PutWithRespKey("v1/volume-snapshots/group", uuid, "", map[string]interface{}{
		"updateVolumeSnapshotGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteVolumeSnapshotGroup deletes VolumeSnapshotGroup
func (cli *ZSClient) DeleteVolumeSnapshotGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volume-snapshots/group", uuid, string(deleteMode))
}

// CheckVolumeSnapshotGroupAvailability operates on VolumeSnapshotGroupAvailability
func (cli *ZSClient) CheckVolumeSnapshotGroupAvailability() (*view.CheckVolumeSnapshotGroupAvailabilityView, error) {
	var resp view.CheckVolumeSnapshotGroupAvailabilityView
	if err := cli.GetWithRespKey("v1/volume-snapshots/groups/availabilities", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RevertVmFromSnapshotGroup operates on VmFromSnapshotGroup
func (cli *ZSClient) RevertVmFromSnapshotGroup(uuid string, params param.RevertVmFromSnapshotGroupParam) (*view.RevertVmFromSnapshotGroupEventView, error) {
	resp := view.RevertVmFromSnapshotGroupEventView{}
	if err := cli.PutWithRespKey("v1/volume-snapshots/group", uuid, "", map[string]interface{}{
		"revertVmFromSnapshotGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UngroupVolumeSnapshotGroup operates on UngroupVolumeSnapshotGroup
func (cli *ZSClient) UngroupVolumeSnapshotGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volume-snapshots/ungroup", uuid, string(deleteMode))
}

// GetMemorySnapshotGroupReference gets MemorySnapshotGroupReference by uuid
func (cli *ZSClient) GetMemorySnapshotGroupReference() (*view.VolumeSnapshotGroupInventoryView, error) {
	var resp view.VolumeSnapshotGroupInventoryView
	if err := cli.GetWithRespKey("v1/memory-snapshots/group/reference", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
