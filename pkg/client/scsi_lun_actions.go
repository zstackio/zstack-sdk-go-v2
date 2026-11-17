// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateScsiLun updates ScsiLun
func (cli *ZSClient) UpdateScsiLun(uuid string, params param.UpdateScsiLunParam) (*view.ScsiLunInventoryView, error) {
	resp := view.ScsiLunInventoryView{}
	if err := cli.PutWithRespKey("v1/storage-devices/scsi-lun", uuid, "", map[string]interface{}{
		"updateScsiLun": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryScsiLun queries ScsiLun list
func (cli *ZSClient) QueryScsiLun(params *param.QueryParam) ([]view.ScsiLunInventoryView, error) {
	var resp []view.ScsiLunInventoryView
	return resp, cli.List("v1/storage-devices/scsi-lun/luns", params, &resp)
}

func (cli *ZSClient) GetScsiLun(uuid string) (*view.ScsiLunInventoryView, error) {
	var resp view.ScsiLunInventoryView
	if err := cli.Get("v1/storage-devices/scsi-lun/luns", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageScsiLun Pagination
func (cli *ZSClient) PageScsiLun(params *param.QueryParam) ([]view.ScsiLunInventoryView, int, error) {
	var scsiLuns []view.ScsiLunInventoryView
	total, err := cli.Page("v1/storage-devices/scsi-lun/luns", params, &scsiLuns)
	return scsiLuns, total, err
}

// CheckScsiLunClusterStatus operates on ScsiLunClusterStatus
func (cli *ZSClient) CheckScsiLunClusterStatus(uuid string, clusterUuid string, params param.CheckScsiLunClusterStatusParam) (*view.ScsiLunClusterStatusInventoryView, error) {
	resp := view.ScsiLunClusterStatusInventoryView{}
	err := cli.PutWithSpec("v1/storage-devices/scsi-lun", uuid, fmt.Sprintf("cluster/%s", clusterUuid), "", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetScsiLunCandidatesForAttachingVm gets ScsiLunCandidatesForAttachingVm by uuid
func (cli *ZSClient) GetScsiLunCandidatesForAttachingVm(uuid string) (*view.ScsiLunInventoryView, error) {
	var resp view.ScsiLunInventoryView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachScsiLunFromHost operates on ScsiLunFromHost
func (cli *ZSClient) DetachScsiLunFromHost(uuid string, params param.DetachScsiLunFromHostParam) (*view.ScsiLunInventoryView, error) {
	resp := view.ScsiLunInventoryView{}
	if err := cli.PutWithRespKey("v1/storage-devices/scsi-lun", uuid, "", map[string]interface{}{
		"detachScsiLunFromHost": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachScsiLunToVmInstance operates on ScsiLunToVmInstance
func (cli *ZSClient) AttachScsiLunToVmInstance(vmInstanceUuid string, uuid string, params param.AttachScsiLunToVmInstanceParam) (*view.ScsiLunInventoryView, error) {
	resp := view.ScsiLunInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/vm-instances/%s/scsi-lun/%s", vmInstanceUuid, uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachScsiLunFromVmInstance operates on ScsiLunFromVmInstance
func (cli *ZSClient) DetachScsiLunFromVmInstance(vmInstanceUuid string, uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/vm-instances", vmInstanceUuid, fmt.Sprintf("scsi-lun/%s", uuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}
