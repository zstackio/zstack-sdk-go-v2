// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateVolume updates Volume
func (cli *ZSClient) UpdateVolume(uuid string, params param.UpdateVolumeParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithRespKey("v1/volumes", uuid, "", map[string]interface{}{
		"updateVolume": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVolume queries Volume list
func (cli *ZSClient) QueryVolume(params *param.QueryParam) ([]view.VolumeInventoryView, error) {
	var resp []view.VolumeInventoryView
	return resp, cli.List("v1/volumes", params, &resp)
}

func (cli *ZSClient) GetVolume(uuid string) (*view.VolumeInventoryView, error) {
	var resp view.VolumeInventoryView
	if err := cli.Get("v1/volumes", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVolume Pagination
func (cli *ZSClient) PageVolume(params *param.QueryParam) ([]view.VolumeInventoryView, int, error) {
	var volumes []view.VolumeInventoryView
	total, err := cli.Page("v1/volumes", params, &volumes)
	return volumes, total, err
}

// ExportNbdVolumes operates on NbdVolumes
func (cli *ZSClient) ExportNbdVolumes(params param.ExportNbdVolumesParam) (*view.ExportNbdVolumesEventView, error) {
	resp := view.ExportNbdVolumesEventView{}
	if err := cli.Post("v1/cbt-task/exportvolume", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVolumeIoThreadPin operates on VolumeIoThreadPin
func (cli *ZSClient) SetVolumeIoThreadPin(uuid string, params param.SetVolumeIoThreadPinParam) (*view.SetVolumeIoThreadPinEventView, error) {
	resp := view.SetVolumeIoThreadPinEventView{}
	if err := cli.PutWithRespKey("v1/volumes", uuid, "", map[string]interface{}{
		"setVolumeIoThreadPin": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RecoverDataVolume operates on DataVolume
func (cli *ZSClient) RecoverDataVolume(uuid string, params param.RecoverDataVolumeParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithRespKey("v1/volumes", uuid, "", map[string]interface{}{
		"recoverDataVolume": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVolumeIoThreadPin gets VolumeIoThreadPin by uuid
func (cli *ZSClient) GetVolumeIoThreadPin(uuid string) (*view.GetVolumeIoThreadPinView, error) {
	var resp view.GetVolumeIoThreadPinView
	if err := cli.GetWithRespKey("v1/volumes", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncVolumeSize operates on VolumeSize
func (cli *ZSClient) SyncVolumeSize(uuid string, params param.SyncVolumeSizeParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithRespKey("v1/volumes", uuid, "", map[string]interface{}{
		"syncVolumeSize": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ResizeRootVolume operates on RootVolume
func (cli *ZSClient) ResizeRootVolume(uuid string, params param.ResizeRootVolumeParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithRespKey("v1/volumes/resize", uuid, "", map[string]interface{}{
		"resizeRootVolume": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDataVolumeAttachableVm gets DataVolumeAttachableVm by uuid
func (cli *ZSClient) GetDataVolumeAttachableVm(uuid string) (*view.VmInstanceInventoryView, error) {
	var resp view.VmInstanceInventoryView
	if err := cli.GetWithRespKey("v1/volumes", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachDataVolumeToHost operates on DataVolumeToHost
func (cli *ZSClient) AttachDataVolumeToHost(volumeUuid string, hostUuid string, params param.AttachDataVolumeToHostParam) (*view.AttachDataVolumeToHostEventView, error) {
	resp := view.AttachDataVolumeToHostEventView{}
	if err := cli.Post(fmt.Sprintf("v1/volumes/%s/hosts/%s", volumeUuid, hostUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDataVolumeFromVolumeBackup creates DataVolumeFromVolumeBackup
func (cli *ZSClient) CreateDataVolumeFromVolumeBackup(backupUuid string, params param.CreateDataVolumeFromVolumeBackupParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/volumes/data-volume/from/volume-template/%s", backupUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachDataVolumeToVm operates on DataVolumeToVm
func (cli *ZSClient) AttachDataVolumeToVm(volumeUuid string, vmInstanceUuid string, params param.AttachDataVolumeToVmParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/volumes/%s/vm-instances/%s", volumeUuid, vmInstanceUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteDataVolume deletes DataVolume
func (cli *ZSClient) DeleteDataVolume(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volumes", uuid, string(deleteMode))
}

// SetVolumeQos operates on VolumeQos
func (cli *ZSClient) SetVolumeQos(uuid string, params param.SetVolumeQosParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithRespKey("v1/volumes", uuid, "", map[string]interface{}{
		"setVolumeQos": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVolumeCapabilities gets VolumeCapabilities by uuid
func (cli *ZSClient) GetVolumeCapabilities(uuid string) (*view.GetVolumeCapabilitiesView, error) {
	var resp view.GetVolumeCapabilitiesView
	if err := cli.GetWithRespKey("v1/volumes", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnexportNbdVolumes operates on UnexportNbdVolumes
func (cli *ZSClient) UnexportNbdVolumes(params param.UnexportNbdVolumesParam) (*view.UnexportNbdVolumesEventView, error) {
	resp := view.UnexportNbdVolumesEventView{}
	if err := cli.Post("v1/cbt-task/unexportvolume", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ResizeDataVolume operates on DataVolume
func (cli *ZSClient) ResizeDataVolume(uuid string, params param.ResizeDataVolumeParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithRespKey("v1/volumes/data/resize", uuid, "", map[string]interface{}{
		"resizeDataVolume": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// BatchSyncVolumeSize operates on SyncVolumeSize
func (cli *ZSClient) BatchSyncVolumeSize(params param.BatchSyncVolumeSizeParam) (*view.BatchSyncVolumeSizeView, error) {
	resp := view.BatchSyncVolumeSizeView{}
	if err := cli.Post("v1/volumes/batch-sync-volumes", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmAttachableDataVolume gets VmAttachableDataVolume by uuid
func (cli *ZSClient) GetVmAttachableDataVolume(uuid string) (*view.VolumeInventoryView, error) {
	var resp view.VolumeInventoryView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVolumeQos gets VolumeQos by uuid
func (cli *ZSClient) GetVolumeQos(uuid string) (*view.GetVolumeQosView, error) {
	var resp view.GetVolumeQosView
	if err := cli.GetWithRespKey("v1/volumes", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDataVolumeFromVolumeSnapshot creates DataVolumeFromVolumeSnapshot
func (cli *ZSClient) CreateDataVolumeFromVolumeSnapshot(volumeSnapshotUuid string, params param.CreateDataVolumeFromVolumeSnapshotParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/volumes/data/from/volume-snapshots/%s", volumeSnapshotUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExpungeDataVolume operates on DataVolume
func (cli *ZSClient) ExpungeDataVolume(uuid string) error {
	params := map[string]interface{}{
		"expungeDataVolume": map[string]interface{}{},
	}
	return cli.Put("v1/volumes", uuid, params, nil)
}

// GetVolumeFormat gets VolumeFormat by uuid
func (cli *ZSClient) GetVolumeFormat() (*view.GetVolumeFormatView, error) {
	var resp view.GetVolumeFormatView
	if err := cli.GetWithRespKey("v1/volumes/formats", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDataVolume creates DataVolume
func (cli *ZSClient) CreateDataVolume(params param.CreateDataVolumeParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.Post("v1/volumes/data", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVolumeQos deletes VolumeQos
func (cli *ZSClient) DeleteVolumeQos(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/volumes", uuid, string(deleteMode))
}

// FlattenVolume operates on FlattenVolume
func (cli *ZSClient) FlattenVolume(uuid string, params param.FlattenVolumeParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithRespKey("v1/volumes", uuid, "", map[string]interface{}{
		"flattenVolume": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FlattenVolumeAsync Async
func (cli *ZSClient) FlattenVolumeAsync(uuid string, params param.FlattenVolumeParam) (string, error) {

	resource := fmt.Sprintf("v1/volumes/%s/actions", uuid)
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// DetachDataVolumeFromVm detaches a data volume without deleting it.
func (cli *ZSClient) DetachDataVolumeFromVm(uuid string, params param.DetachDataVolumeFromVmParam) (*view.VolumeInventoryView, error) {
	body := struct {
		VmUuid *string `json:"vmUuid,omitempty"`
		param.BaseParam
	}{
		VmUuid:    params.Params.VmUuid,
		BaseParam: params.BaseParam,
	}
	path := fmt.Sprintf("v1/volumes/%s/vm-instances", uuid)
	query := url.Values{}
	if params.Params.VmUuid != nil && *params.Params.VmUuid != "" {
		query.Set("vmUuid", *params.Params.VmUuid)
		path = fmt.Sprintf("%s?%s", path, query.Encode())
	}

	resp := view.VolumeInventoryView{}
	if err := cli.DeleteWithBody(path, body, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachDataVolumeFromHost detaches a data volume without deleting it.
func (cli *ZSClient) DetachDataVolumeFromHost(volumeUuid string, params param.DetachDataVolumeFromHostParam) error {
	body := struct {
		HostUuid *string `json:"hostUuid,omitempty"`
		param.BaseParam
	}{
		HostUuid:  params.Params.HostUuid,
		BaseParam: params.BaseParam,
	}
	return cli.DeleteWithBody(fmt.Sprintf("v1/volumes/%s/hosts", volumeUuid), body, nil)
}

// CreateDataVolumeFromVolumeTemplate creates DataVolumeFromVolumeTemplate
func (cli *ZSClient) CreateDataVolumeFromVolumeTemplate(imageUuid string, params param.CreateDataVolumeFromVolumeTemplateParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/volumes/data/from/data-volume-templates/%s", imageUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeVolumeState changes VolumeState
func (cli *ZSClient) ChangeVolumeState(uuid string, params param.ChangeVolumeStateParam) (*view.VolumeInventoryView, error) {
	resp := view.VolumeInventoryView{}
	if err := cli.PutWithRespKey("v1/volumes", uuid, "", map[string]interface{}{
		"changeVolumeState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
