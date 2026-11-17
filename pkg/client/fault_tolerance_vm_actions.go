// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// CreateFaultToleranceVmInstance creates FaultToleranceVmInstance
func (cli *ZSClient) CreateFaultToleranceVmInstance(params param.CreateFaultToleranceVmInstanceParam) (*view.CreateFaultToleranceVmInstanceEventView, error) {
	resp := view.CreateFaultToleranceVmInstanceEventView{}
	if err := cli.Post("v1/vm-instances/fault-tolerance", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFaultToleranceVms gets FaultToleranceVms by uuid
func (cli *ZSClient) GetFaultToleranceVms() (*view.GetFaultToleranceVmsView, error) {
	var resp view.GetFaultToleranceVmsView
	if err := cli.GetWithRespKey("v1/vm-instances/fault-tolerance/sub-vms", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// FailoverFaultToleranceVm operates on FailoverFaultToleranceVm
func (cli *ZSClient) FailoverFaultToleranceVm(params param.FailoverFaultToleranceVmParam) (*view.FailoverFaultToleranceVmEventView, error) {
	resp := view.FailoverFaultToleranceVmEventView{}
	if err := cli.PutWithRespKey("v1/vm-instances/fault-tolerance", "", "", map[string]interface{}{
		"failoverFaultToleranceVm": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryFaultToleranceVm queries FaultToleranceVm list
func (cli *ZSClient) QueryFaultToleranceVm(params *param.QueryParam) ([]view.FaultToleranceVmGroupInventoryView, error) {
	var resp []view.FaultToleranceVmGroupInventoryView
	return resp, cli.List("v1/vm-instances/fault-tolerance", params, &resp)
}

func (cli *ZSClient) GetFaultToleranceVm(uuid string) (*view.FaultToleranceVmGroupInventoryView, error) {
	var resp view.FaultToleranceVmGroupInventoryView
	if err := cli.Get("v1/vm-instances/fault-tolerance", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageFaultToleranceVm Pagination
func (cli *ZSClient) PageFaultToleranceVm(params *param.QueryParam) ([]view.FaultToleranceVmGroupInventoryView, int, error) {
	var faultToleranceVms []view.FaultToleranceVmGroupInventoryView
	total, err := cli.Page("v1/vm-instances/fault-tolerance", params, &faultToleranceVms)
	return faultToleranceVms, total, err
}
