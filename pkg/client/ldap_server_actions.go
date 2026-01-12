// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddLdapServer adds LdapServer
func (cli *ZSClient) AddLdapServer(params param.AddLdapServerParam) (*view.LdapServerInventoryView, error) {
	var resp view.AddLdapServerEventView
	if err := cli.Post("v1/ldap/servers", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryLdapServer queries LdapServer list
func (cli *ZSClient) QueryLdapServer(params *param.QueryParam) ([]view.LdapServerInventoryView, error) {
	var resp []view.LdapServerInventoryView
	return resp, cli.List("v1/ldap/servers", params, &resp)
}

func (cli *ZSClient) GetLdapServer(uuid string) (*view.LdapServerInventoryView, error) {
	var resp view.LdapServerInventoryView
	if err := cli.Get("v1/ldap/servers", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// SyncLdapServer operates on LdapServer
func (cli *ZSClient) SyncLdapServer(uuid string, params param.SyncLdapServerParam) (*view.LongJobInventoryView, error) {
	var resp view.SyncLdapServerEventView
	err := cli.PutWithSpec("v1/ldap/servers", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}

// SyncLdapServerAsync Async
func (cli *ZSClient) SyncLdapServerAsync(params param.SyncLdapServerParam) (string, error) {

	resource := "v1/ldap/servers/{uuid}/actions"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}
// DeleteLdapServer deletes LdapServer
func (cli *ZSClient) DeleteLdapServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/ldap/servers", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// UpdateLdapServer updates LdapServer
func (cli *ZSClient) UpdateLdapServer(ldapServerUuid string, params param.UpdateLdapServerParam) (*view.LdapServerInventoryView, error) {
	var resp view.UpdateLdapServerEventView
	err := cli.PutWithSpec("v1/ldap/servers", fmt.Sprintf(\"%s\", ldapServerUuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
