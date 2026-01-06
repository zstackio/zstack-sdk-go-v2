// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
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
// SyncLdapServer operates on LdapServer
func (cli *ZSClient) SyncLdapServer(uuid string, params param.SyncLdapServerParam) (*view.LongJobInventoryView, error) {
	var resp view.SyncLdapServerEventView
	if err := cli.Put("v1/ldap/servers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteLdapServer deletes LdapServer
func (cli *ZSClient) DeleteLdapServer(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ldap/servers/{uuid}", uuid, string(deleteMode))
}
// UpdateLdapServer updates LdapServer
func (cli *ZSClient) UpdateLdapServer(uuid string, params param.UpdateLdapServerParam) (*view.LdapServerInventoryView, error) {
	var resp view.UpdateLdapServerEventView
	if err := cli.Put("v1/ldap/servers/{ldapServerUuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
