// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateVolumeSnapshot updates VolumeSnapshot
func (cli *ZSClient) UpdateVolumeSnapshot(uuid string, params param.UpdateVolumeSnapshotParam) (*view.VolumeSnapshotInventoryView, error) {
	resp := view.VolumeSnapshotInventoryView{}
	if err := cli.PutWithRespKey("v1/volume-snapshots", uuid, "", map[string]interface{}{
		"updateVolumeSnapshot": params.Params,
	}, &resp); err != nil {
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
func (cli *ZSClient) CreateVolumeSnapshot(volumeUuid string, params param.CreateVolumeSnapshotParam) (*view.VolumeSnapshotInventoryView, error) {
	resp := view.VolumeSnapshotInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/volumes/%s/volume-snapshots", volumeUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BatchDeleteVolumeSnapshot operates on DeleteVolumeSnapshot
func (cli *ZSClient) BatchDeleteVolumeSnapshot(params param.BatchDeleteVolumeSnapshotParam) (*view.BatchDeleteVolumeSnapshotEventView, error) {
	resp := view.BatchDeleteVolumeSnapshotEventView{}
	if err := cli.PutWithRespKey("v1/volume-snapshots/batch-delete", "", "", map[string]interface{}{
		"batchDeleteVolumeSnapshot": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ShrinkVolumeSnapshot operates on ShrinkVolumeSnapshot
func (cli *ZSClient) ShrinkVolumeSnapshot(uuid string, params param.ShrinkVolumeSnapshotParam) (*view.ShrinkVolumeSnapshotEventView, error) {
	resp := view.ShrinkVolumeSnapshotEventView{}
	if err := cli.PutWithRespKey("v1/volume-snapshots/shrink", uuid, "", map[string]interface{}{
		"shrinkVolumeSnapshot": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateVolumeSnapshotChain operates on VolumeSnapshotChain
func (cli *ZSClient) ValidateVolumeSnapshotChain(uuid string, params param.ValidateVolumeSnapshotChainParam) (*view.ValidateVolumeSnapshotChainEventView, error) {
	resp := view.ValidateVolumeSnapshotChainEventView{}
	if err := cli.PutWithRespKey("v1/volumes", uuid, "", map[string]interface{}{
		"validateVolumeSnapshotChain": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UndoSnapshotCreation operates on UndoSnapshotCreation
func (cli *ZSClient) UndoSnapshotCreation(uuid string, params param.UndoSnapshotCreationParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithRespKey("v1/volumes", uuid, "", map[string]interface{}{
		"undoSnapshotCreation": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVolumeSnapshotSize gets VolumeSnapshotSize by uuid
func (cli *ZSClient) GetVolumeSnapshotSize(uuid string, params param.GetVolumeSnapshotSizeParam) (*view.GetVolumeSnapshotSizeEventView, error) {
	resp := view.GetVolumeSnapshotSizeEventView{}
	if err := cli.PutWithRespKey("v1/volume-snapshots", uuid, "", map[string]interface{}{
		"getVolumeSnapshotSize": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVolumesSnapshot creates VolumesSnapshot
func (cli *ZSClient) CreateVolumesSnapshot(params param.CreateVolumesSnapshotParam) (*view.VolumeSnapshotInventoryView, error) {
	resp := view.VolumeSnapshotInventoryView{}
	if err := cli.Post("v1/volumes/volume-snapshots", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RevertVolumeFromSnapshot operates on VolumeFromSnapshot
func (cli *ZSClient) RevertVolumeFromSnapshot(uuid string, params param.RevertVolumeFromSnapshotParam) (*view.RevertVolumeFromSnapshotEventView, error) {
	resp := view.RevertVolumeFromSnapshotEventView{}
	if err := cli.PutWithRespKey("v1/volume-snapshots", uuid, "", map[string]interface{}{
		"revertVolumeFromSnapshot": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RevertVolumeFromSnapshotAsync Async
func (cli *ZSClient) RevertVolumeFromSnapshotAsync(uuid string, params param.RevertVolumeFromSnapshotParam) (string, error) {

	resource := fmt.Sprintf("v1/volume-snapshots/%s/actions", uuid)
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}
