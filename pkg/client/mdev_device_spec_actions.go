// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateMdevDeviceSpec updates MdevDeviceSpec
func (cli *ZSClient) UpdateMdevDeviceSpec(uuid string, params param.UpdateMdevDeviceSpecParam) (*view.MdevDeviceSpecInventoryView, error) {
	resp := view.MdevDeviceSpecInventoryView{}
	if err := cli.PutWithRespKey("v1/mdev-device-specs", uuid, "", map[string]interface{}{
		"updateMdevDeviceSpec": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryMdevDeviceSpec queries MdevDeviceSpec list
func (cli *ZSClient) QueryMdevDeviceSpec(params *param.QueryParam) ([]view.MdevDeviceSpecInventoryView, error) {
	var resp []view.MdevDeviceSpecInventoryView
	return resp, cli.List("v1/mdev-device-specs", params, &resp)
}

func (cli *ZSClient) GetMdevDeviceSpec(uuid string) (*view.MdevDeviceSpecInventoryView, error) {
	var resp view.MdevDeviceSpecInventoryView
	if err := cli.Get("v1/mdev-device-specs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMdevDeviceSpec Pagination
func (cli *ZSClient) PageMdevDeviceSpec(params *param.QueryParam) ([]view.MdevDeviceSpecInventoryView, int, error) {
	var mdevDeviceSpecs []view.MdevDeviceSpecInventoryView
	total, err := cli.Page("v1/mdev-device-specs", params, &mdevDeviceSpecs)
	return mdevDeviceSpecs, total, err
}

// RemoveMdevDeviceSpecFromVmInstance removes MdevDeviceSpecFromVmInstance
func (cli *ZSClient) RemoveMdevDeviceSpecFromVmInstance(mdevSpecUuid string, vmInstanceUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/mdev-device-specs", mdevSpecUuid, fmt.Sprintf("vm-instances/%s", vmInstanceUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// GetMdevDeviceSpecCandidates gets MdevDeviceSpecCandidates by uuid
func (cli *ZSClient) GetMdevDeviceSpecCandidates() (*view.MdevDeviceSpecInventoryView, error) {
	var resp view.MdevDeviceSpecInventoryView
	if err := cli.GetWithRespKey("v1/mdev-device-specs/candidates", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddMdevDeviceSpecToVmInstance adds MdevDeviceSpecToVmInstance
func (cli *ZSClient) AddMdevDeviceSpecToVmInstance(mdevSpecUuid string, vmInstanceUuid string, params param.AddMdevDeviceSpecToVmInstanceParam) (*view.VmInstanceMdevDeviceSpecRefInventoryView, error) {
	resp := view.VmInstanceMdevDeviceSpecRefInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/mdev-device-specs/%s/vm-instances/%s", mdevSpecUuid, vmInstanceUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
