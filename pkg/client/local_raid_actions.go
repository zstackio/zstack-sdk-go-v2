// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// QueryLocalRaidController queries LocalRaidController list
func (cli *ZSClient) QueryLocalRaidController(params *param.QueryParam) ([]view.RaidPhysicalDriveInventoryView, error) {
	var resp []view.RaidPhysicalDriveInventoryView
	return resp, cli.List("v1/storage-devices/local-raid/controllers", params, &resp)
}

func (cli *ZSClient) GetLocalRaidController(uuid string) (*view.RaidPhysicalDriveInventoryView, error) {
	var resp view.RaidPhysicalDriveInventoryView
	if err := cli.Get("v1/storage-devices/local-raid/controllers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageLocalRaidController Pagination
func (cli *ZSClient) PageLocalRaidController(params *param.QueryParam) ([]view.RaidPhysicalDriveInventoryView, int, error) {
	var localRaidControllers []view.RaidPhysicalDriveInventoryView
	total, err := cli.Page("v1/storage-devices/local-raid/controllers", params, &localRaidControllers)
	return localRaidControllers, total, err
}

// LocateLocalRaidPhysicalDrive operates on LocalRaidPhysicalDrive
func (cli *ZSClient) LocateLocalRaidPhysicalDrive(uuid string, params param.LocateLocalRaidPhysicalDriveParam) (*view.RaidPhysicalDriveInventoryView, error) {
	resp := view.RaidPhysicalDriveInventoryView{}
	if err := cli.PutWithRespKey("v1/storage-devices/local-raid/physical-drives", uuid, "", map[string]interface{}{
		"locateLocalRaidPhysicalDrive": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SelfTestLocalRaid operates on LocalRaid
func (cli *ZSClient) SelfTestLocalRaid(uuid string, params param.SelfTestLocalRaidParam) (*view.SelfTestLocalRaidEventView, error) {
	resp := view.SelfTestLocalRaidEventView{}
	if err := cli.PutWithRespKey("v1/storage-devices/local-raid/physical-drives", uuid, "", map[string]interface{}{
		"selfTestLocalRaid": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryLocalRaidPhysicalDrive queries LocalRaidPhysicalDrive list
func (cli *ZSClient) QueryLocalRaidPhysicalDrive(params *param.QueryParam) ([]view.RaidPhysicalDriveInventoryView, error) {
	var resp []view.RaidPhysicalDriveInventoryView
	return resp, cli.List("v1/storage-devices/local-raid/physical-drives", params, &resp)
}

func (cli *ZSClient) GetLocalRaidPhysicalDrive(uuid string) (*view.RaidPhysicalDriveInventoryView, error) {
	var resp view.RaidPhysicalDriveInventoryView
	if err := cli.Get("v1/storage-devices/local-raid/physical-drives", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageLocalRaidPhysicalDrive Pagination
func (cli *ZSClient) PageLocalRaidPhysicalDrive(params *param.QueryParam) ([]view.RaidPhysicalDriveInventoryView, int, error) {
	var localRaidPhysicalDrives []view.RaidPhysicalDriveInventoryView
	total, err := cli.Page("v1/storage-devices/local-raid/physical-drives", params, &localRaidPhysicalDrives)
	return localRaidPhysicalDrives, total, err
}

// QueryPhysicalDriveSelfTestHistory queries PhysicalDriveSelfTestHistory list
func (cli *ZSClient) QueryPhysicalDriveSelfTestHistory(params *param.QueryParam) ([]view.PhysicalDriveSmartSelfTestHistoryInventoryView, error) {
	var resp []view.PhysicalDriveSmartSelfTestHistoryInventoryView
	return resp, cli.List("v1/storage-devices/local-raid/physical-drives/self-test", params, &resp)
}

func (cli *ZSClient) GetPhysicalDriveSelfTestHistory(uuid string) (*view.PhysicalDriveSmartSelfTestHistoryInventoryView, error) {
	var resp view.PhysicalDriveSmartSelfTestHistoryInventoryView
	if err := cli.Get("v1/storage-devices/local-raid/physical-drives/self-test", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePhysicalDriveSelfTestHistory Pagination
func (cli *ZSClient) PagePhysicalDriveSelfTestHistory(params *param.QueryParam) ([]view.PhysicalDriveSmartSelfTestHistoryInventoryView, int, error) {
	var physicalDriveSelfTestHistories []view.PhysicalDriveSmartSelfTestHistoryInventoryView
	total, err := cli.Page("v1/storage-devices/local-raid/physical-drives/self-test", params, &physicalDriveSelfTestHistories)
	return physicalDriveSelfTestHistories, total, err
}

// GetLocalRaidPhysicalDriveSmart gets LocalRaidPhysicalDriveSmart by uuid
func (cli *ZSClient) GetLocalRaidPhysicalDriveSmart(uuid string) (*view.GetLocalRaidPhysicalDriveSmartView, error) {
	var resp view.GetLocalRaidPhysicalDriveSmartView
	if err := cli.GetWithRespKey("v1/storage-devices/local-raid/physical-drives", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RefreshLocalRaid operates on LocalRaid
func (cli *ZSClient) RefreshLocalRaid(params param.RefreshLocalRaidParam) (*view.RaidControllerInventoryView, error) {
	resp := view.RaidControllerInventoryView{}
	if err := cli.PutWithRespKey("v1/storage-devices/local-raid/actions", "", "", map[string]interface{}{
		"refreshLocalRaid": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
