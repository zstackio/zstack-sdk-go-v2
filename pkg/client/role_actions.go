// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateRole creates Role
func (cli *ZSClient) CreateRole(params param.CreateRoleParam) (*view.RoleInventoryView, error) {
	resp := view.RoleInventoryView{}
	if err := cli.Post("v1/identities/roles", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryRole queries Role list
func (cli *ZSClient) QueryRole(params *param.QueryParam) ([]view.RoleInventoryView, error) {
	var resp []view.RoleInventoryView
	return resp, cli.List("v1/identities/roles", params, &resp)
}

func (cli *ZSClient) GetRole(uuid string) (*view.RoleInventoryView, error) {
	var resp view.RoleInventoryView
	if err := cli.Get("v1/identities/roles", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageRole Pagination
func (cli *ZSClient) PageRole(params *param.QueryParam) ([]view.RoleInventoryView, int, error) {
	var roles []view.RoleInventoryView
	total, err := cli.Page("v1/identities/roles", params, &roles)
	return roles, total, err
}
// DeleteRole deletes Role
func (cli *ZSClient) DeleteRole(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/identities/roles", uuid, string(deleteMode))
}
// UpdateRole updates Role
func (cli *ZSClient) UpdateRole(uuid string, params param.UpdateRoleParam) (*view.RoleInventoryView, error) {
	resp := view.RoleInventoryView{}
	if err := cli.PutWithRespKey("v1/identities/roles", uuid, "", map[string]interface{}{
		"updateRole": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeRoleState changes RoleState
func (cli *ZSClient) ChangeRoleState(uuid string, params param.ChangeRoleStateParam) (*view.RoleInventoryView, error) {
	resp := view.RoleInventoryView{}
	if err := cli.PutWithRespKey("v1/identities/roles", uuid, "", map[string]interface{}{
		"changeRoleState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachRoleToAccount operates on RoleToAccount
func (cli *ZSClient) AttachRoleToAccount(accountUuid string, roleUuid string, params param.AttachRoleToAccountParam) (*view.AttachRoleToAccountEventView, error) {
	resp := view.AttachRoleToAccountEventView{}
	if err := cli.Post(fmt.Sprintf("v1/identities/accounts/%s/roles/%s", accountUuid, roleUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemovePolicyStatementsFromRole removes PolicyStatementsFromRole
func (cli *ZSClient) RemovePolicyStatementsFromRole(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/identities/roles", uuid, string(deleteMode))
}

// DetachPolicyFromRole operates on PolicyFromRole
func (cli *ZSClient) DetachPolicyFromRole(policyUuid string, roleUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/identities/policies", policyUuid, fmt.Sprintf("roles/%s", roleUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// AddPolicyStatementsToRole adds PolicyStatementsToRole
func (cli *ZSClient) AddPolicyStatementsToRole(uuid string, params param.AddPolicyStatementsToRoleParam) (*view.AddPolicyStatementsToRoleEventView, error) {
	resp := view.AddPolicyStatementsToRoleEventView{}
	if err := cli.Post(fmt.Sprintf("v1/identities/roles/%s/policy-statements", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachRoleFromAccount operates on RoleFromAccount
func (cli *ZSClient) DetachRoleFromAccount(accountUuid string, roleUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/identities/accounts", accountUuid, fmt.Sprintf("roles/%s", roleUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// AttachPolicyToRole operates on PolicyToRole
func (cli *ZSClient) AttachPolicyToRole(policyUuid string, roleUuid string, params param.AttachPolicyToRoleParam) (*view.AttachPolicyToRoleEventView, error) {
	resp := view.AttachPolicyToRoleEventView{}
	if err := cli.Post(fmt.Sprintf("v1/identities/policies/%s/roles/%s", policyUuid, roleUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
