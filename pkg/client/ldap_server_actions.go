// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddLdapServer adds LdapServer
func (cli *ZSClient) AddLdapServer(params param.AddLdapServerParam) (*view.LdapServerInventoryView, error) {
	resp := view.LdapServerInventoryView{}
	if err := cli.Post("v1/ldap/servers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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

// PageLdapServer Pagination
func (cli *ZSClient) PageLdapServer(params *param.QueryParam) ([]view.LdapServerInventoryView, int, error) {
	var ldapServers []view.LdapServerInventoryView
	total, err := cli.Page("v1/ldap/servers", params, &ldapServers)
	return ldapServers, total, err
}
// SyncLdapServer operates on LdapServer
func (cli *ZSClient) SyncLdapServer(uuid string, params param.SyncLdapServerParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.PutWithRespKey("v1/ldap/servers", uuid, "", map[string]interface{}{
		"syncLdapServer": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncLdapServerAsync Async
func (cli *ZSClient) SyncLdapServerAsync(uuid string, params param.SyncLdapServerParam) (string, error) {

	resource := fmt.Sprintf("v1/ldap/servers/%s/actions", uuid)
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
	return cli.Delete("v1/ldap/servers", uuid, string(deleteMode))
}
// UpdateLdapServer updates LdapServer
func (cli *ZSClient) UpdateLdapServer(ldapServerUuid string, params param.UpdateLdapServerParam) (*view.LdapServerInventoryView, error) {
	resp := view.LdapServerInventoryView{}
	if err := cli.PutWithRespKey("v1/ldap/servers", ldapServerUuid, "", map[string]interface{}{
		"updateLdapServer": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
