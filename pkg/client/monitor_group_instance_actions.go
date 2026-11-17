// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryMonitorGroupInstance queries MonitorGroupInstance list
func (cli *ZSClient) QueryMonitorGroupInstance(params *param.QueryParam) ([]view.MonitorGroupInstanceInventoryView, error) {
	var resp []view.MonitorGroupInstanceInventoryView
	return resp, cli.List("v1/zwatch/monitorgroups/instances", params, &resp)
}

func (cli *ZSClient) GetMonitorGroupInstance(uuid string) (*view.MonitorGroupInstanceInventoryView, error) {
	var resp view.MonitorGroupInstanceInventoryView
	if err := cli.Get("v1/zwatch/monitorgroups/instances", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMonitorGroupInstance Pagination
func (cli *ZSClient) PageMonitorGroupInstance(params *param.QueryParam) ([]view.MonitorGroupInstanceInventoryView, int, error) {
	var monitorGroupInstances []view.MonitorGroupInstanceInventoryView
	total, err := cli.Page("v1/zwatch/monitorgroups/instances", params, &monitorGroupInstances)
	return monitorGroupInstances, total, err
}

// AddInstanceToMonitorGroup adds InstanceToMonitorGroup
func (cli *ZSClient) AddInstanceToMonitorGroup(groupUuid string, params param.AddInstanceToMonitorGroupParam) (*view.MonitorGroupInstanceInventoryView, error) {
	resp := view.MonitorGroupInstanceInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/zwatch/monitorgroups/%s/actions", groupUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveInstanceFromMonitorGroup removes InstanceFromMonitorGroup
func (cli *ZSClient) RemoveInstanceFromMonitorGroup(groupUuid string, instanceUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/zwatch/monitorgroups", groupUuid, fmt.Sprintf("actions/%s", instanceUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}
