// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryUser queries User list
func (cli *ZSClient) QueryUser(params *param.QueryParam) ([]view.UserInventoryView, error) {
	var resp []view.UserInventoryView
	return resp, cli.List("v1/accounts/users", params, &resp)
}

func (cli *ZSClient) GetUser(uuid string) (*view.UserInventoryView, error) {
	var resp view.UserInventoryView
	if err := cli.Get("v1/accounts/users", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageUser Pagination
func (cli *ZSClient) PageUser(params *param.QueryParam) ([]view.UserInventoryView, int, error) {
	var users []view.UserInventoryView
	total, err := cli.Page("v1/accounts/users", params, &users)
	return users, total, err
}
// DeleteUser deletes User
func (cli *ZSClient) DeleteUser(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/users", uuid, string(deleteMode))
}
// UpdateUser updates User
func (cli *ZSClient) UpdateUser(params param.UpdateUserParam) (*view.UserInventoryView, error) {
	resp := view.UserInventoryView{}
	if err := cli.PutWithRespKey("v1/accounts/users/actions", "", "", map[string]interface{}{
		"updateUser": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateUser creates User
func (cli *ZSClient) CreateUser(params param.CreateUserParam) (*view.UserInventoryView, error) {
	resp := view.UserInventoryView{}
	if err := cli.Post("v1/accounts/users", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachPoliciesToUser operates on PoliciesToUser
func (cli *ZSClient) AttachPoliciesToUser(userUuid string, params param.AttachPoliciesToUserParam) (*view.AttachPoliciesToUserEventView, error) {
	resp := view.AttachPoliciesToUserEventView{}
	if err := cli.Post(fmt.Sprintf("v1/accounts/users/%s/policy-collection", userUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachPolicyFromUser operates on PolicyFromUser
func (cli *ZSClient) DetachPolicyFromUser(userUuid string, policyUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/accounts/users", userUuid, fmt.Sprintf("policies/%s", policyUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}

// AttachPolicyToUser operates on PolicyToUser
func (cli *ZSClient) AttachPolicyToUser(userUuid string, params param.AttachPolicyToUserParam) (*view.AttachPolicyToUserEventView, error) {
	resp := view.AttachPolicyToUserEventView{}
	if err := cli.Post(fmt.Sprintf("v1/accounts/users/%s/policies", userUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachPoliciesFromUser operates on PoliciesFromUser
func (cli *ZSClient) DetachPoliciesFromUser(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/accounts/users", uuid, string(deleteMode))
}
