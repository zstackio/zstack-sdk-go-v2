// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySecurityGroup queries SecurityGroup list
func (cli *ZSClient) QuerySecurityGroup(params *param.QueryParam) ([]view.SecurityGroupInventoryView, error) {
	var resp []view.SecurityGroupInventoryView
	return resp, cli.List("v1/security-groups", params, &resp)
}

func (cli *ZSClient) GetSecurityGroup(uuid string) (*view.SecurityGroupInventoryView, error) {
	var resp view.SecurityGroupInventoryView
	if err := cli.Get("v1/security-groups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSecurityGroup Pagination
func (cli *ZSClient) PageSecurityGroup(params *param.QueryParam) ([]view.SecurityGroupInventoryView, int, error) {
	var securityGroups []view.SecurityGroupInventoryView
	total, err := cli.Page("v1/security-groups", params, &securityGroups)
	return securityGroups, total, err
}
// CreateSecurityGroup creates SecurityGroup
func (cli *ZSClient) CreateSecurityGroup(params param.CreateSecurityGroupParam) (*view.SecurityGroupInventoryView, error) {
	resp := view.SecurityGroupInventoryView{}
	if err := cli.Post("v1/security-groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteSecurityGroup deletes SecurityGroup
func (cli *ZSClient) DeleteSecurityGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/security-groups", uuid, string(deleteMode))
}
// UpdateSecurityGroup updates SecurityGroup
func (cli *ZSClient) UpdateSecurityGroup(uuid string, params param.UpdateSecurityGroupParam) (*view.SecurityGroupInventoryView, error) {
	resp := view.SecurityGroupInventoryView{}
	if err := cli.PutWithRespKey("v1/security-groups", uuid, "", map[string]interface{}{
		"updateSecurityGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeSecurityGroupState changes SecurityGroupState
func (cli *ZSClient) ChangeSecurityGroupState(uuid string, params param.ChangeSecurityGroupStateParam) (*view.SecurityGroupInventoryView, error) {
	resp := view.SecurityGroupInventoryView{}
	if err := cli.PutWithRespKey("v1/security-groups", uuid, "", map[string]interface{}{
		"changeSecurityGroupState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryVmNicInSecurityGroup queries VmNicInSecurityGroup list
func (cli *ZSClient) QueryVmNicInSecurityGroup(params *param.QueryParam) ([]view.VmNicSecurityGroupRefInventoryView, error) {
	var resp []view.VmNicSecurityGroupRefInventoryView
	return resp, cli.List("v1/security-groups/vm-instances/nics", params, &resp)
}

func (cli *ZSClient) GetVmNicInSecurityGroup(uuid string) (*view.VmNicSecurityGroupRefInventoryView, error) {
	var resp view.VmNicSecurityGroupRefInventoryView
	if err := cli.Get("v1/security-groups/vm-instances/nics", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVmNicInSecurityGroup Pagination
func (cli *ZSClient) PageVmNicInSecurityGroup(params *param.QueryParam) ([]view.VmNicSecurityGroupRefInventoryView, int, error) {
	var vmNicInSecurityGroups []view.VmNicSecurityGroupRefInventoryView
	total, err := cli.Page("v1/security-groups/vm-instances/nics", params, &vmNicInSecurityGroups)
	return vmNicInSecurityGroups, total, err
}

// GetCandidateVmNicForSecurityGroup gets CandidateVmNicForSecurityGroup by uuid
func (cli *ZSClient) GetCandidateVmNicForSecurityGroup(uuid string) (*view.VmNicInventoryView, error) {
	var resp view.VmNicInventoryView
	if err := cli.GetWithRespKey("v1/security-groups", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetVmNicSecurityGroup operates on VmNicSecurityGroup
func (cli *ZSClient) SetVmNicSecurityGroup(vmNicUuid string, params param.SetVmNicSecurityGroupParam) (*view.VmNicSecurityGroupRefInventoryView, error) {
	resp := view.VmNicSecurityGroupRefInventoryView{}
	if err := cli.PutWithRespKey("v1/security-groups/nics", vmNicUuid, "", map[string]interface{}{
		"setVmNicSecurityGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachSecurityGroupFromL3Network operates on SecurityGroupFromL3Network
func (cli *ZSClient) DetachSecurityGroupFromL3Network(securityGroupUuid string, l3NetworkUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/security-groups", securityGroupUuid, fmt.Sprintf("l3-networks/%s", l3NetworkUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// AttachSecurityGroupToL3Network operates on SecurityGroupToL3Network
func (cli *ZSClient) AttachSecurityGroupToL3Network(securityGroupUuid string, l3NetworkUuid string, params param.AttachSecurityGroupToL3NetworkParam) (*view.SecurityGroupInventoryView, error) {
	resp := view.SecurityGroupInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/security-groups/%s/l3-networks/%s", securityGroupUuid, l3NetworkUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddVmNicToSecurityGroup adds VmNicToSecurityGroup
func (cli *ZSClient) AddVmNicToSecurityGroup(securityGroupUuid string, params param.AddVmNicToSecurityGroupParam) (*view.AddVmNicToSecurityGroupEventView, error) {
	resp := view.AddVmNicToSecurityGroupEventView{}
	if _, err := cli.PostWithAsync(fmt.Sprintf("v1/security-groups/%s/vm-instances/nics", securityGroupUuid), "", params, &resp, false); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVmNicFromSecurityGroup deletes VmNicFromSecurityGroup
func (cli *ZSClient) DeleteVmNicFromSecurityGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/security-groups", uuid, string(deleteMode))
}
