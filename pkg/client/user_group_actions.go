// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteUserGroup deletes UserGroup
func (cli *ZSClient) DeleteUserGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/groups", uuid, string(deleteMode))
}
// CreateUserGroup creates UserGroup
func (cli *ZSClient) CreateUserGroup(params param.CreateUserGroupParam) (*view.UserGroupInventoryView, error) {
	resp := view.UserGroupInventoryView{}
	if err := cli.Post("v1/accounts/groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryUserGroup queries UserGroup list
func (cli *ZSClient) QueryUserGroup(params *param.QueryParam) ([]view.UserGroupInventoryView, error) {
	var resp []view.UserGroupInventoryView
	return resp, cli.List("v1/accounts/groups", params, &resp)
}

func (cli *ZSClient) GetUserGroup(uuid string) (*view.UserGroupInventoryView, error) {
	var resp view.UserGroupInventoryView
	if err := cli.Get("v1/accounts/groups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageUserGroup Pagination
func (cli *ZSClient) PageUserGroup(params *param.QueryParam) ([]view.UserGroupInventoryView, int, error) {
	var userGroups []view.UserGroupInventoryView
	total, err := cli.Page("v1/accounts/groups", params, &userGroups)
	return userGroups, total, err
}
// UpdateUserGroup updates UserGroup
func (cli *ZSClient) UpdateUserGroup(params param.UpdateUserGroupParam) (*view.UserGroupInventoryView, error) {
	resp := view.UserGroupInventoryView{}
	if err := cli.PutWithRespKey("v1/accounts/groups/actions", "", "", map[string]interface{}{
		"updateUserGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddUserToGroup adds UserToGroup
func (cli *ZSClient) AddUserToGroup(groupUuid string, params param.AddUserToGroupParam) (*view.AddUserToGroupEventView, error) {
	resp := view.AddUserToGroupEventView{}
	if err := cli.Post(fmt.Sprintf("v1/accounts/groups/%s/users", groupUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveUserFromGroup removes UserFromGroup
func (cli *ZSClient) RemoveUserFromGroup(groupUuid string, userUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/accounts/groups", groupUuid, fmt.Sprintf("users/%s", userUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// AttachPolicyToUserGroup operates on PolicyToUserGroup
func (cli *ZSClient) AttachPolicyToUserGroup(groupUuid string, params param.AttachPolicyToUserGroupParam) (*view.AttachPolicyToUserGroupEventView, error) {
	resp := view.AttachPolicyToUserGroupEventView{}
	if err := cli.Post(fmt.Sprintf("v1/accounts/groups/%s/policies", groupUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachPolicyFromUserGroup operates on PolicyFromUserGroup
func (cli *ZSClient) DetachPolicyFromUserGroup(groupUuid string, policyUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/accounts/groups", groupUuid, fmt.Sprintf("policies/%s", policyUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}
