// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ReconnectHost operates on Host
func (cli *ZSClient) ReconnectHost(uuid string, params param.ReconnectHostParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.PutWithRespKey("v1/hosts", uuid, "", map[string]interface{}{
		"reconnectHost": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateHost updates Host
func (cli *ZSClient) UpdateHost(uuid string, params param.UpdateHostParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.PutWithRespKey("v1/hosts", uuid, "", map[string]interface{}{
		"updateHost": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteHost deletes Host
func (cli *ZSClient) DeleteHost(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hosts", uuid, string(deleteMode))
}
// QueryHost queries Host list
func (cli *ZSClient) QueryHost(params *param.QueryParam) ([]view.HostInventoryView, error) {
	var resp []view.HostInventoryView
	return resp, cli.List("v1/hosts", params, &resp)
}

func (cli *ZSClient) GetHost(uuid string) (*view.HostInventoryView, error) {
	var resp view.HostInventoryView
	if err := cli.Get("v1/hosts", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageHost Pagination
func (cli *ZSClient) PageHost(params *param.QueryParam) ([]view.HostInventoryView, int, error) {
	var hosts []view.HostInventoryView
	total, err := cli.Page("v1/hosts", params, &hosts)
	return hosts, total, err
}

// ChangeHostState changes HostState
func (cli *ZSClient) ChangeHostState(uuid string, params param.ChangeHostStateParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.PutWithRespKey("v1/hosts", uuid, "", map[string]interface{}{
		"changeHostState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostPowerStatus gets HostPowerStatus by uuid
func (cli *ZSClient) GetHostPowerStatus(uuid string, params param.GetHostPowerStatusParam) (*view.HostIpmiInventoryView, error) {
	resp := view.HostIpmiInventoryView{}
	if err := cli.PutWithRespKey("v1/hosts/power", uuid, "", map[string]interface{}{
		"getHostPowerStatus": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PowerOnHost operates on PowerOnHost
func (cli *ZSClient) PowerOnHost(uuid string, params param.PowerOnHostParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.PutWithRespKey("v1/hosts/power", uuid, "", map[string]interface{}{
		"powerOnHost": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateHostIommuState updates HostIommuState
func (cli *ZSClient) UpdateHostIommuState(uuid string, params param.UpdateHostIommuStateParam) (*view.UpdateHostIommuStateEventView, error) {
	resp := view.UpdateHostIommuStateEventView{}
	if err := cli.PutWithRespKey("v1/pci-device/hosts", uuid, "", map[string]interface{}{
		"updateHostIommuState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostAllocatorStrategies gets HostAllocatorStrategies by uuid
func (cli *ZSClient) GetHostAllocatorStrategies() (*view.GetHostAllocatorStrategiesView, error) {
	var resp view.GetHostAllocatorStrategiesView
	if err := cli.GetWithRespKey("v1/hosts/allocators/strategies", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInterfaceServiceTypeStatistic gets InterfaceServiceTypeStatistic by uuid
func (cli *ZSClient) GetInterfaceServiceTypeStatistic() (*view.GetInterfaceServiceTypeStatisticView, error) {
	var resp view.GetInterfaceServiceTypeStatisticView
	if err := cli.GetWithRespKey("v1/hosts/hosts-network-interfaces/service-type-statistic", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PowerOffHost operates on PowerOffHost
func (cli *ZSClient) PowerOffHost(params param.PowerOffHostParam) (*view.PowerOffHostEventView, error) {
	resp := view.PowerOffHostEventView{}
	if err := cli.PutWithRespKey("v1/hosts/power-off/actions", "", "", map[string]interface{}{
		"powerOffHost": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCpuMemoryCapacity gets CpuMemoryCapacity by uuid
func (cli *ZSClient) GetCpuMemoryCapacity() (*view.GetCpuMemoryCapacityView, error) {
	var resp view.GetCpuMemoryCapacityView
	if err := cli.GetWithRespKey("v1/hosts/capacities/cpu-memory", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCpuMemoryCapacityWithParams passes the scope required by Cloud. When no
// zone, cluster, or host is selected, callers must set All to true.
func (cli *ZSClient) GetCpuMemoryCapacityWithParams(params param.GetCpuMemoryCapacityParamDetail) (*view.GetCpuMemoryCapacityView, error) {
	var resp view.GetCpuMemoryCapacityView
	if err := cli.GetWithRespKey("v1/hosts/capacities/cpu-memory", "", "", &params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostMultipathTopology gets HostMultipathTopology by uuid
func (cli *ZSClient) GetHostMultipathTopology() (*view.GetHostMultipathTopologyView, error) {
	var resp view.GetHostMultipathTopologyView
	if err := cli.GetWithRespKey("v1/storage-devices/multipath/topology", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostNetworkFacts gets HostNetworkFacts by uuid
func (cli *ZSClient) GetHostNetworkFacts(uuid string) (*view.GetHostNetworkFactsView, error) {
	var resp view.GetHostNetworkFactsView
	if err := cli.GetWithRespKey("v1/hosts/network-facts", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostIommuState gets HostIommuState by uuid
func (cli *ZSClient) GetHostIommuState(uuid string) (*view.GetHostIommuStateView, error) {
	var resp view.GetHostIommuStateView
	if err := cli.GetWithRespKey("v1/pci-device/hosts", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostCandidatesForVmMigration gets HostCandidatesForVmMigration by uuid
func (cli *ZSClient) GetHostCandidatesForVmMigration(uuid string) (*view.HostInventoryView, error) {
	var resp view.HostInventoryView
	if err := cli.GetWithRespKey("v1/primary-storage/hosts", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostIommuStatus gets HostIommuStatus by uuid
func (cli *ZSClient) GetHostIommuStatus(uuid string) (*view.GetHostIommuStatusView, error) {
	var resp view.GetHostIommuStatusView
	if err := cli.GetWithRespKey("v1/pci-device/hosts", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHypervisorTypes gets HypervisorTypes by uuid
func (cli *ZSClient) GetHypervisorTypes() (*view.GetHypervisorTypesView, error) {
	var resp view.GetHypervisorTypesView
	if err := cli.GetWithRespKey("v1/hosts/hypervisor-types", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostTask gets HostTask by uuid
func (cli *ZSClient) GetHostTask() (*view.GetChainTaskView, error) {
	var resp view.GetChainTaskView
	if err := cli.GetWithRespKey("v1/hosts/task-details", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVmMigrationCandidateHosts gets VmMigrationCandidateHosts by uuid
func (cli *ZSClient) GetVmMigrationCandidateHosts(uuid string) (*view.HostInventoryView, error) {
	var resp view.HostInventoryView
	if err := cli.GetWithRespKey("v1/vm-instances", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostWebSshUrl gets HostWebSshUrl by uuid
func (cli *ZSClient) GetHostWebSshUrl(params param.GetHostWebSshUrlParam) (*view.GetHostWebSshUrlEventView, error) {
	resp := view.GetHostWebSshUrlEventView{}
	if err := cli.Post("v1/hosts/webssh", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PowerResetHost operates on PowerResetHost
func (cli *ZSClient) PowerResetHost(uuid string, params param.PowerResetHostParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.PutWithRespKey("v1/hosts/power", uuid, "", map[string]interface{}{
		"powerResetHost": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ShutdownHost operates on ShutdownHost
func (cli *ZSClient) ShutdownHost(uuid string, params param.ShutdownHostParam) (*view.HostInventoryView, error) {
	resp := view.HostInventoryView{}
	if err := cli.PutWithRespKey("v1/hosts/power", uuid, "", map[string]interface{}{
		"shutdownHost": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AllocateHostResource operates on HostResource
func (cli *ZSClient) AllocateHostResource(uuid string, params param.AllocateHostResourceParam) (*view.AllocateHostResourceEventView, error) {
	resp := view.AllocateHostResourceEventView{}
	if err := cli.Post(fmt.Sprintf("v1/hosts/%s/allocate-resource", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostResourceAllocation gets HostResourceAllocation by uuid
func (cli *ZSClient) GetHostResourceAllocation(uuid string, params param.GetHostResourceAllocationParam) (*view.GetHostResourceAllocationEventView, error) {
	resp := view.GetHostResourceAllocationEventView{}
	if err := cli.Post(fmt.Sprintf("v1/hosts/%s/resource-allocation", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHostNUMATopology gets HostNUMATopology by uuid
func (cli *ZSClient) GetHostNUMATopology(uuid string, params param.GetHostNUMATopologyParam) (*view.GetHostNUMATopologyEventView, error) {
	resp := view.GetHostNUMATopologyEventView{}
	if err := cli.Post(fmt.Sprintf("v1/hosts/%s/numa", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
