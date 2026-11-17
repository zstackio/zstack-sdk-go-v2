// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CloneVmInstance operates on VmInstance
func (cli *ZSClient) CloneVmInstance(vmInstanceUuid string, params param.CloneVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"cloneVmInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ResumeVmInstance operates on VmInstance
func (cli *ZSClient) ResumeVmInstance(uuid string, params param.ResumeVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"resumeVmInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// StartVmInstance starts VmInstance
func (cli *ZSClient) StartVmInstance(uuid string, params param.StartVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"startVmInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// StopVmInstance stops VmInstance
func (cli *ZSClient) StopVmInstance(uuid string, params param.StopVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"stopVmInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryVmInstance queries VmInstance list
func (cli *ZSClient) QueryVmInstance(params *param.QueryParam) ([]view.VmInstanceInventoryView, error) {
	var resp []view.VmInstanceInventoryView
	return resp, cli.List("v1/vm-instances", params, &resp)
}

func (cli *ZSClient) GetVmInstance(uuid string) (*view.VmInstanceInventoryView, error) {
	var resp view.VmInstanceInventoryView
	if err := cli.Get("v1/vm-instances", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmInstance Pagination
func (cli *ZSClient) PageVmInstance(params *param.QueryParam) ([]view.VmInstanceInventoryView, int, error) {
	var vmInstances []view.VmInstanceInventoryView
	total, err := cli.Page("v1/vm-instances", params, &vmInstances)
	return vmInstances, total, err
}
// ExpungeVmInstance operates on VmInstance
func (cli *ZSClient) ExpungeVmInstance(uuid string) error {
	params := map[string]interface{}{
		"expungeVmInstance": map[string]interface{}{},
	}
	return cli.Put("v1/vm-instances", uuid, params, nil)
}
// RebootVmInstance operates on VmInstance
func (cli *ZSClient) RebootVmInstance(uuid string, params param.RebootVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"rebootVmInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateVmInstance updates VmInstance
func (cli *ZSClient) UpdateVmInstance(uuid string, params param.UpdateVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"updateVmInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DestroyVmInstance destroys VmInstance
func (cli *ZSClient) DestroyVmInstance(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances", uuid, string(deleteMode))
}
// CreateVmInstance creates VmInstance
func (cli *ZSClient) CreateVmInstance(params param.CreateVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post("v1/vm-instances", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// RecoverVmInstance operates on VmInstance
func (cli *ZSClient) RecoverVmInstance(uuid string, params param.RecoverVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"recoverVmInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmClockTrack operates on VmClockTrack
func (cli *ZSClient) SetVmClockTrack(uuid string, params param.SetVmClockTrackParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmClockTrack": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmInstanceHygonMdev operates on VmInstanceHygonMdev
func (cli *ZSClient) SetVmInstanceHygonMdev(uuid string, params param.SetVmInstanceHygonMdevParam) (*view.SetVmInstanceHygonMdevEventView, error) {
	resp := view.SetVmInstanceHygonMdevEventView{}
	if err := cli.Post(fmt.Sprintf("v1/vm-instances/%s/hygon-mdev", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachL3NetworkToVm operates on L3NetworkToVm
func (cli *ZSClient) AttachL3NetworkToVm(vmInstanceUuid string, l3NetworkUuid string, params param.AttachL3NetworkToVmParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/vm-instances/%s/l3-networks/%s", vmInstanceUuid, l3NetworkUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmBootOrder gets VmBootOrder by uuid
func (cli *ZSClient) GetVmBootOrder(uuid string) (*view.GetVmBootOrderView, error) {
	var resp view.GetVmBootOrderView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmBootOrder operates on VmBootOrder
func (cli *ZSClient) SetVmBootOrder(uuid string, params param.SetVmBootOrderParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmBootOrder": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmInstanceHaLevel deletes VmInstanceHaLevel
func (cli *ZSClient) DeleteVmInstanceHaLevel(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances", uuid, string(deleteMode))
}

// ReimageVmInstance operates on ReimageVmInstance
func (cli *ZSClient) ReimageVmInstance(vmInstanceUuid string, params param.ReimageVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"reimageVmInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmHostname gets VmHostname by uuid
func (cli *ZSClient) GetVmHostname(uuid string) (*view.GetVmHostnameView, error) {
	var resp view.GetVmHostnameView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVmNetworkConfig updates VmNetworkConfig
func (cli *ZSClient) UpdateVmNetworkConfig(vmInstanceUuid string, params param.UpdateVmNetworkConfigParam) (*view.UpdateVmNetworkConfigEventView, error) {
	resp := view.UpdateVmNetworkConfigEventView{}
	if err := cli.PutWithRespKey("v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"updateVmNetworkConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmStaticIp operates on VmStaticIp
func (cli *ZSClient) SetVmStaticIp(vmInstanceUuid string, params param.SetVmStaticIpParam) (*view.SetVmStaticIpEventView, error) {
	resp := view.SetVmStaticIpEventView{}
	if err := cli.PutWithRespKey("v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"setVmStaticIp": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmSshKey gets VmSshKey by uuid
func (cli *ZSClient) GetVmSshKey(uuid string) (*view.GetVmSshKeyView, error) {
	var resp view.GetVmSshKeyView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmRDP operates on VmRDP
func (cli *ZSClient) SetVmRDP(uuid string, params param.SetVmRDPParam) (*view.SetVmRDPEventView, error) {
	resp := view.SetVmRDPEventView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmRDP": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmSshKey deletes VmSshKey
func (cli *ZSClient) DeleteVmSshKey(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances", uuid, string(deleteMode))
}

// DeleteVmConsolePassword deletes VmConsolePassword
func (cli *ZSClient) DeleteVmConsolePassword(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances", uuid, string(deleteMode))
}

// GetVmStartingCandidateClustersHosts gets VmStartingCandidateClustersHosts by uuid
func (cli *ZSClient) GetVmStartingCandidateClustersHosts(uuid string) (*view.GetVmStartingCandidateClustersHostsView, error) {
	var resp view.GetVmStartingCandidateClustersHostsView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UploadFileToVm operates on UploadFileToVm
func (cli *ZSClient) UploadFileToVm(params param.UploadFileToVmParam) (*view.UploadFileToVmEventView, error) {
	resp := view.UploadFileToVmEventView{}
	if err := cli.Post("v1/upload-file", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmUserDefinedXml operates on VmUserDefinedXml
func (cli *ZSClient) SetVmUserDefinedXml(vmInstanceUuid string, params param.SetVmUserDefinedXmlParam) (*view.SetVmUserDefinedXmlEventView, error) {
	resp := view.SetVmUserDefinedXmlEventView{}
	if err := cli.PutWithRespKey("v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"setVmUserDefinedXml": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListVMsFromKVMHost operates on ListVMsFromKVMHost
func (cli *ZSClient) ListVMsFromKVMHost(params param.ListVMsFromKVMHostParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post("v1/v2v", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// TakeVmConsoleScreenshot operates on TakeVmConsoleScreenshot
func (cli *ZSClient) TakeVmConsoleScreenshot(uuid string, params param.TakeVmConsoleScreenshotParam) (*view.TakeVmConsoleScreenshotEventView, error) {
	resp := view.TakeVmConsoleScreenshotEventView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"takeVmConsoleScreenshot": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmSecurityLevel operates on VmSecurityLevel
func (cli *ZSClient) SetVmSecurityLevel(uuid string, params param.SetVmSecurityLevelParam) (*view.SetVmSecurityLevelEventView, error) {
	resp := view.SetVmSecurityLevelEventView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmSecurityLevel": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmStaticIp deletes VmStaticIp
func (cli *ZSClient) DeleteVmStaticIp(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances", uuid, string(deleteMode))
}

// GetVmCapabilities gets VmCapabilities by uuid
func (cli *ZSClient) GetVmCapabilities(uuid string) (*view.GetVmCapabilitiesView, error) {
	var resp view.GetVmCapabilitiesView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmFromVolumeBackup creates VmFromVolumeBackup
func (cli *ZSClient) CreateVmFromVolumeBackup(backupUuid string, params param.CreateVmFromVolumeBackupParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/vm-instances/from/vm-backup/%s", backupUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmHostname operates on VmHostname
func (cli *ZSClient) SetVmHostname(uuid string, params param.SetVmHostnameParam) (*view.SetVmHostnameEventView, error) {
	resp := view.SetVmHostnameEventView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmHostname": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmConsolePassword operates on VmConsolePassword
func (cli *ZSClient) SetVmConsolePassword(uuid string, params param.SetVmConsolePasswordParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmConsolePassword": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmRDP gets VmRDP by uuid
func (cli *ZSClient) GetVmRDP(uuid string) (*view.GetVmRDPView, error) {
	var resp view.GetVmRDPView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateConsolePassword updates ConsolePassword
func (cli *ZSClient) UpdateConsolePassword(uuid string, params param.UpdateConsolePasswordParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"updateConsolePassword": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmInstanceFromVolumeSnapshotGroup creates VmInstanceFromVolumeSnapshotGroup
func (cli *ZSClient) CreateVmInstanceFromVolumeSnapshotGroup(volumeSnapshotGroupUuid string, params param.CreateVmInstanceFromVolumeSnapshotGroupParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/vm-instances/from/volume-snapshots/group/%s", volumeSnapshotGroupUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmEmulatorPinning gets VmEmulatorPinning by uuid
func (cli *ZSClient) GetVmEmulatorPinning(uuid string) (*view.GetVmEmulatorPinningView, error) {
	var resp view.GetVmEmulatorPinningView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmInstanceFromOvf creates VmInstanceFromOvf
func (cli *ZSClient) CreateVmInstanceFromOvf(params param.CreateVmInstanceFromOvfParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post("v1/ovf/create-vm-instance", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmInstanceFromOvfAsync Async
func (cli *ZSClient) CreateVmInstanceFromOvfAsync(params param.CreateVmInstanceFromOvfParam) (string, error) {

	resource := "v1/ovf/create-vm-instance"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// ChangeVmImage changes VmImage
func (cli *ZSClient) ChangeVmImage(vmInstanceUuid string, params param.ChangeVmImageParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"changeVmImage": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmUsbRedirect operates on VmUsbRedirect
func (cli *ZSClient) SetVmUsbRedirect(uuid string, params param.SetVmUsbRedirectParam) (*view.SetVmUsbRedirectEventView, error) {
	resp := view.SetVmUsbRedirectEventView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmUsbRedirect": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmsCapabilities gets VmsCapabilities by uuid
func (cli *ZSClient) GetVmsCapabilities() (*view.GetVmsCapabilitiesView, error) {
	var resp view.GetVmsCapabilitiesView
	if err := cli.GetWithRespKey("v1/vm-instances/capabilities", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmXml gets VmXml by uuid
func (cli *ZSClient) GetVmXml(uuid string) (*view.GetVmXmlView, error) {
	var resp view.GetVmXmlView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmInstanceFirstBootDevice gets VmInstanceFirstBootDevice by uuid
func (cli *ZSClient) GetVmInstanceFirstBootDevice(uuid string) (*view.GetVmInstanceFirstBootDeviceView, error) {
	var resp view.GetVmInstanceFirstBootDeviceView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnmountVmInstanceRecoveryPoint operates on UnmountVmInstanceRecoveryPoint
func (cli *ZSClient) UnmountVmInstanceRecoveryPoint(params param.UnmountVmInstanceRecoveryPointParam) (*view.UnmountVmInstanceRecoveryPointEventView, error) {
	resp := view.UnmountVmInstanceRecoveryPointEventView{}
	if err := cli.Post("v1/cdp-backup-storage/unmount-recovery-point", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DescribeVmInstanceRecoveryPoint operates on DescribeVmInstanceRecoveryPoint
func (cli *ZSClient) DescribeVmInstanceRecoveryPoint(uuid string) (*view.DescribeVmInstanceRecoveryPointView, error) {
	var resp view.DescribeVmInstanceRecoveryPointView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmCleanTraffic operates on VmCleanTraffic
func (cli *ZSClient) SetVmCleanTraffic(uuid string, params param.SetVmCleanTrafficParam) (*view.SetVmCleanTrafficEventView, error) {
	resp := view.SetVmCleanTrafficEventView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmCleanTraffic": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmBootMode operates on VmBootMode
func (cli *ZSClient) SetVmBootMode(uuid string, params param.SetVmBootModeParam) (*view.SetVmBootModeEventView, error) {
	resp := view.SetVmBootModeEventView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmBootMode": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ProtectVmInstanceRecoveryPoint operates on ProtectVmInstanceRecoveryPoint
func (cli *ZSClient) ProtectVmInstanceRecoveryPoint(vmInstanceUuid string, params param.ProtectVmInstanceRecoveryPointParam) (*view.ProtectVmInstanceRecoveryPointEventView, error) {
	resp := view.ProtectVmInstanceRecoveryPointEventView{}
	if err := cli.PutWithRespKey("v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"protectVmInstanceRecoveryPoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmInstanceHaLevel operates on VmInstanceHaLevel
func (cli *ZSClient) SetVmInstanceHaLevel(uuid string, params param.SetVmInstanceHaLevelParam) (*view.SetVmInstanceHaLevelEventView, error) {
	resp := view.SetVmInstanceHaLevelEventView{}
	if err := cli.Post(fmt.Sprintf("v1/vm-instances/%s/ha-levels", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmConsolePassword gets VmConsolePassword by uuid
func (cli *ZSClient) GetVmConsolePassword(uuid string) (*view.GetVmConsolePasswordView, error) {
	var resp view.GetVmConsolePasswordView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmInstanceHaLevel gets VmInstanceHaLevel by uuid
func (cli *ZSClient) GetVmInstanceHaLevel(uuid string) (*view.GetVmInstanceHaLevelView, error) {
	var resp view.GetVmInstanceHaLevelView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmFromVmBackup creates VmFromVmBackup
func (cli *ZSClient) CreateVmFromVmBackup(groupUuid string, params param.CreateVmFromVmBackupParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/vm-instances/from/vm-backups/%s", groupUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmMonitorNumber gets VmMonitorNumber by uuid
func (cli *ZSClient) GetVmMonitorNumber(uuid string) (*view.GetVmMonitorNumberView, error) {
	var resp view.GetVmMonitorNumberView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVirtualizerInfo gets VirtualizerInfo by uuid
func (cli *ZSClient) GetVirtualizerInfo() (*view.VirtualizerInfoInventoryView, error) {
	var resp view.VirtualizerInfoInventoryView
	if err := cli.GetWithRespKey("v1/vm-instances/virtualizer-info", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmUsbRedirect gets VmUsbRedirect by uuid
func (cli *ZSClient) GetVmUsbRedirect(uuid string) (*view.GetVmUsbRedirectView, error) {
	var resp view.GetVmUsbRedirectView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PauseVmInstance operates on PauseVmInstance
func (cli *ZSClient) PauseVmInstance(uuid string, params param.PauseVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"pauseVmInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmQga operates on VmQga
func (cli *ZSClient) SetVmQga(uuid string, params param.SetVmQgaParam) (*view.SetVmQgaEventView, error) {
	resp := view.SetVmQgaEventView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmQga": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmQxlMemory operates on VmQxlMemory
func (cli *ZSClient) SetVmQxlMemory(uuid string, params param.SetVmQxlMemoryParam) (*view.SetVmQxlMemoryEventView, error) {
	resp := view.SetVmQxlMemoryEventView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmQxlMemory": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// MigrateVm operates on Vm
func (cli *ZSClient) MigrateVm(vmInstanceUuid string, params param.MigrateVmParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"migrateVm": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// MigrateVmAsync Async
func (cli *ZSClient) MigrateVmAsync(vmInstanceUuid string, params param.MigrateVmParam) (string, error) {

	resource := fmt.Sprintf("v1/vm-instances/%s/actions", vmInstanceUuid)
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// ChangeVmPassword changes VmPassword
func (cli *ZSClient) ChangeVmPassword(uuid string, params param.ChangeVmPasswordParam) (*view.ChangeVmPasswordEventView, error) {
	resp := view.ChangeVmPasswordEventView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"changeVmPassword": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FlattenVmInstance operates on FlattenVmInstance
func (cli *ZSClient) FlattenVmInstance(uuid string, params param.FlattenVmInstanceParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"flattenVmInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FlattenVmInstanceAsync Async
func (cli *ZSClient) FlattenVmInstanceAsync(uuid string, params param.FlattenVmInstanceParam) (string, error) {

	resource := fmt.Sprintf("v1/vm-instances/%s/actions", uuid)
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// GetVmNuma gets VmNuma by uuid
func (cli *ZSClient) GetVmNuma(uuid string) (*view.GetVmNumaView, error) {
	var resp view.GetVmNumaView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVmPriority updates VmPriority
func (cli *ZSClient) UpdateVmPriority(uuid string, params param.UpdateVmPriorityParam) (*view.UpdateVmPriorityEventView, error) {
	resp := view.UpdateVmPriorityEventView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"updateVmPriority": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmHostname deletes VmHostname
func (cli *ZSClient) DeleteVmHostname(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances", uuid, string(deleteMode))
}

// SetVmSoundType operates on VmSoundType
func (cli *ZSClient) SetVmSoundType(uuid string, params param.SetVmSoundTypeParam) (*view.SetVmSoundTypeEventView, error) {
	resp := view.SetVmSoundTypeEventView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmSoundType": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateZonesClustersHostsForCreatingVm gets CandidateZonesClustersHostsForCreatingVm by uuid
func (cli *ZSClient) GetCandidateZonesClustersHostsForCreatingVm() (*view.GetCandidateZonesClustersHostsForCreatingVmView, error) {
	var resp view.GetCandidateZonesClustersHostsForCreatingVmView
	if err := cli.GetWithRespKey("v1/vm-instances/candidate-destinations", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmInstanceFromVolumeSnapshot creates VmInstanceFromVolumeSnapshot
func (cli *ZSClient) CreateVmInstanceFromVolumeSnapshot(volumeSnapshotUuid string, params param.CreateVmInstanceFromVolumeSnapshotParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/vm-instances/from/volume-snapshots/%s", volumeSnapshotUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmSshKey operates on VmSshKey
func (cli *ZSClient) SetVmSshKey(uuid string, params param.SetVmSshKeyParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmSshKey": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExecuteGuestVmCommand operates on ExecuteGuestVmCommand
func (cli *ZSClient) ExecuteGuestVmCommand(params param.ExecuteGuestVmCommandParam) (*view.ExecuteGuestVmCommandEventView, error) {
	resp := view.ExecuteGuestVmCommandEventView{}
	if err := cli.Post("v1/vm-instances/commands/exec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmDeviceAddress gets VmDeviceAddress by uuid
func (cli *ZSClient) GetVmDeviceAddress() (*view.GetVmDeviceAddressView, error) {
	var resp view.GetVmDeviceAddressView
	if err := cli.GetWithRespKey("v1/vm-instances/devices", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmQga gets VmQga by uuid
func (cli *ZSClient) GetVmQga(uuid string) (*view.GetVmQgaView, error) {
	var resp view.GetVmQgaView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmvNUMATopology gets VmvNUMATopology by uuid
func (cli *ZSClient) GetVmvNUMATopology(uuid string) (*view.GetVmvNUMATopologyView, error) {
	var resp view.GetVmvNUMATopologyView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FstrimVm operates on FstrimVm
func (cli *ZSClient) FstrimVm(uuid string, params param.FstrimVmParam) (*view.FstrimVmEventView, error) {
	resp := view.FstrimVmEventView{}
	if err := cli.Post(fmt.Sprintf("v1/vm-instances/%s/actions", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnprotectVmInstanceRecoveryPoint operates on UnprotectVmInstanceRecoveryPoint
func (cli *ZSClient) UnprotectVmInstanceRecoveryPoint(vmInstanceUuid string, params param.UnprotectVmInstanceRecoveryPointParam) (*view.UnprotectVmInstanceRecoveryPointEventView, error) {
	resp := view.UnprotectVmInstanceRecoveryPointEventView{}
	if err := cli.PutWithRespKey("v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"unprotectVmInstanceRecoveryPoint": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmMonitorNumber operates on VmMonitorNumber
func (cli *ZSClient) SetVmMonitorNumber(uuid string, params param.SetVmMonitorNumberParam) (*view.SetVmMonitorNumberEventView, error) {
	resp := view.SetVmMonitorNumberEventView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmMonitorNumber": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmUserDefinedXml deletes VmUserDefinedXml
func (cli *ZSClient) DeleteVmUserDefinedXml(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances", uuid, string(deleteMode))
}

// GetVmInstanceProtectedRecoveryPoints gets VmInstanceProtectedRecoveryPoints by uuid
func (cli *ZSClient) GetVmInstanceProtectedRecoveryPoints(uuid string) (*view.GetVmInstanceProtectedRecoveryPointsView, error) {
	var resp view.GetVmInstanceProtectedRecoveryPointsView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncVmClock operates on VmClock
func (cli *ZSClient) SyncVmClock(uuid string, params param.SyncVmClockParam) (*view.SyncVmClockEventView, error) {
	resp := view.SyncVmClockEventView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"syncVmClock": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmBootVolume operates on VmBootVolume
func (cli *ZSClient) SetVmBootVolume(vmInstanceUuid string, params param.SetVmBootVolumeParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances", vmInstanceUuid, "", map[string]interface{}{
		"setVmBootVolume": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmConsoleMode operates on VmConsoleMode
func (cli *ZSClient) SetVmConsoleMode(uuid string, params param.SetVmConsoleModeParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmConsoleMode": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmEmulatorPinning operates on VmEmulatorPinning
func (cli *ZSClient) SetVmEmulatorPinning(uuid string, params param.SetVmEmulatorPinningParam) (*view.SetVmEmulatorPinningEventView, error) {
	resp := view.SetVmEmulatorPinningEventView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmEmulatorPinning": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmNuma operates on VmNuma
func (cli *ZSClient) SetVmNuma(uuid string, params param.SetVmNumaParam) (*view.SetVmNumaEventView, error) {
	resp := view.SetVmNumaEventView{}
	if err := cli.PutWithRespKey("v1/vm-instances", uuid, "", map[string]interface{}{
		"setVmNuma": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmTask gets VmTask by uuid
func (cli *ZSClient) GetVmTask() (*view.GetChainTaskView, error) {
	var resp view.GetChainTaskView
	if err := cli.GetWithRespKey("v1/vm-instances/task-details", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVmInstanceFromVolume creates VmInstanceFromVolume
func (cli *ZSClient) CreateVmInstanceFromVolume(params param.CreateVmInstanceFromVolumeParam) (*view.VmInstanceInventoryView, error) {
	resp := view.VmInstanceInventoryView{}
	if err := cli.Post("v1/vm-instances/from/volume", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmInstanceRecoveryPoints gets VmInstanceRecoveryPoints by uuid
func (cli *ZSClient) GetVmInstanceRecoveryPoints(uuid string) (*view.GetVmInstanceRecoveryPointsView, error) {
	var resp view.GetVmInstanceRecoveryPointsView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmBootMode deletes VmBootMode
func (cli *ZSClient) DeleteVmBootMode(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/vm-instances", uuid, string(deleteMode))
}

// MountVmInstanceRecoveryPoint operates on MountVmInstanceRecoveryPoint
func (cli *ZSClient) MountVmInstanceRecoveryPoint(params param.MountVmInstanceRecoveryPointParam) (*view.MountVmInstanceRecoveryPointEventView, error) {
	resp := view.MountVmInstanceRecoveryPointEventView{}
	if err := cli.Post("v1/cdp-backup-storage/mount-recovery-point", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
