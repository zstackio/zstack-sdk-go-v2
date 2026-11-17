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

// GetCandidateLdapEntryForBinding gets CandidateLdapEntryForBinding by uuid
func (cli *ZSClient) GetCandidateLdapEntryForBinding() (*view.GetLdapEntryView, error) {
	var resp view.GetLdapEntryView
	if err := cli.GetWithRespKey("v1/ldap/entries/candidates", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLdapEntry gets LdapEntry by uuid
func (cli *ZSClient) GetLdapEntry() (*view.GetLdapEntryView, error) {
	var resp view.GetLdapEntryView
	if err := cli.GetWithRespKey("v1/ldap/entry", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryLdapBinding queries LdapBinding list
func (cli *ZSClient) QueryLdapBinding(params *param.QueryParam) ([]view.LdapAccountRefInventoryView, error) {
	var resp []view.LdapAccountRefInventoryView
	return resp, cli.List("v1/ldap/bindings", params, &resp)
}

func (cli *ZSClient) GetLdapBinding(uuid string) (*view.LdapAccountRefInventoryView, error) {
	var resp view.LdapAccountRefInventoryView
	if err := cli.Get("v1/ldap/bindings", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageLdapBinding Pagination
func (cli *ZSClient) PageLdapBinding(params *param.QueryParam) ([]view.LdapAccountRefInventoryView, int, error) {
	var ldapBindings []view.LdapAccountRefInventoryView
	total, err := cli.Page("v1/ldap/bindings", params, &ldapBindings)
	return ldapBindings, total, err
}

// GetLdapServerAvailableAttributes gets LdapServerAvailableAttributes by uuid
func (cli *ZSClient) GetLdapServerAvailableAttributes(uuid string) (*view.GetLdapServerAvailableAttributesView, error) {
	var resp view.GetLdapServerAvailableAttributesView
	if err := cli.GetWithRespKey("v1/ldap/server/attributes", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateLdapBinding creates LdapBinding
func (cli *ZSClient) CreateLdapBinding(params param.CreateLdapBindingParam) (*view.LdapAccountRefInventoryView, error) {
	resp := view.LdapAccountRefInventoryView{}
	if err := cli.Post("v1/ldap/bindings", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteLdapBinding deletes LdapBinding
func (cli *ZSClient) DeleteLdapBinding(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ldap/bindings", uuid, string(deleteMode))
}
