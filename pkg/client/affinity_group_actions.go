// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateAffinityGroup updates AffinityGroup
func (cli *ZSClient) UpdateAffinityGroup(uuid string, params param.UpdateAffinityGroupParam) (*view.AffinityGroupInventoryView, error) {
	resp := view.AffinityGroupInventoryView{}
	if err := cli.PutWithRespKey("v1/affinity-groups", uuid, "", map[string]interface{}{
		"updateAffinityGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteAffinityGroup deletes AffinityGroup
func (cli *ZSClient) DeleteAffinityGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/affinity-groups", uuid, string(deleteMode))
}
// CreateAffinityGroup creates AffinityGroup
func (cli *ZSClient) CreateAffinityGroup(params param.CreateAffinityGroupParam) (*view.AffinityGroupInventoryView, error) {
	resp := view.AffinityGroupInventoryView{}
	if err := cli.Post("v1/affinity-groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryAffinityGroup queries AffinityGroup list
func (cli *ZSClient) QueryAffinityGroup(params *param.QueryParam) ([]view.AffinityGroupInventoryView, error) {
	var resp []view.AffinityGroupInventoryView
	return resp, cli.List("v1/affinity-groups", params, &resp)
}

func (cli *ZSClient) GetAffinityGroup(uuid string) (*view.AffinityGroupInventoryView, error) {
	var resp view.AffinityGroupInventoryView
	if err := cli.Get("v1/affinity-groups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAffinityGroup Pagination
func (cli *ZSClient) PageAffinityGroup(params *param.QueryParam) ([]view.AffinityGroupInventoryView, int, error) {
	var affinityGroups []view.AffinityGroupInventoryView
	total, err := cli.Page("v1/affinity-groups", params, &affinityGroups)
	return affinityGroups, total, err
}

// RemoveVmFromAffinityGroup removes VmFromAffinityGroup
func (cli *ZSClient) RemoveVmFromAffinityGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/affinity-groups", uuid, string(deleteMode))
}

// GetCandidateAffinityGroupForCreatingVm gets CandidateAffinityGroupForCreatingVm by uuid
func (cli *ZSClient) GetCandidateAffinityGroupForCreatingVm() (*view.AffinityGroupInventoryView, error) {
	var resp view.AffinityGroupInventoryView
	if err := cli.GetWithRespKey("v1/vm-instances/candidate-affinityGroup", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeAffinityGroupState changes AffinityGroupState
func (cli *ZSClient) ChangeAffinityGroupState(uuid string, params param.ChangeAffinityGroupStateParam) (*view.AffinityGroupInventoryView, error) {
	resp := view.AffinityGroupInventoryView{}
	if err := cli.PutWithRespKey("v1/affinity-groups", uuid, "", map[string]interface{}{
		"changeAffinityGroupState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateAffinityGroupForAttachingVm gets CandidateAffinityGroupForAttachingVm by uuid
func (cli *ZSClient) GetCandidateAffinityGroupForAttachingVm() (*view.AffinityGroupInventoryView, error) {
	var resp view.AffinityGroupInventoryView
	if err := cli.GetWithRespKey("v1/affinityGroup/attachingVm", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddVmToAffinityGroup adds VmToAffinityGroup
func (cli *ZSClient) AddVmToAffinityGroup(affinityGroupUuid string, uuid string, params param.AddVmToAffinityGroupParam) (*view.AffinityGroupInventoryView, error) {
	resp := view.AffinityGroupInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/affinity-groups/%s/vm-instances/%s", affinityGroupUuid, uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateVMForAttachingAffinityGroup gets CandidateVMForAttachingAffinityGroup by uuid
func (cli *ZSClient) GetCandidateVMForAttachingAffinityGroup() (*view.VmInstanceInventoryView, error) {
	var resp view.VmInstanceInventoryView
	if err := cli.GetWithRespKey("v1/VM/attachingGroup", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
