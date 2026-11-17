// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// QueryIAM2LdapBinding queries IAM2LdapBinding list
func (cli *ZSClient) QueryIAM2LdapBinding(params *param.QueryParam) ([]view.LdapResourceRefInventoryView, error) {
	var resp []view.LdapResourceRefInventoryView
	return resp, cli.List("v1/iam2/ldap/bindings", params, &resp)
}

func (cli *ZSClient) GetIAM2LdapBinding(uuid string) (*view.LdapResourceRefInventoryView, error) {
	var resp view.LdapResourceRefInventoryView
	if err := cli.Get("v1/iam2/ldap/bindings", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIAM2LdapBinding Pagination
func (cli *ZSClient) PageIAM2LdapBinding(params *param.QueryParam) ([]view.LdapResourceRefInventoryView, int, error) {
	var iAM2LdapBindings []view.LdapResourceRefInventoryView
	total, err := cli.Page("v1/iam2/ldap/bindings", params, &iAM2LdapBindings)
	return iAM2LdapBindings, total, err
}

// GetCandidateLdapEntryForIAM2Binding gets CandidateLdapEntryForIAM2Binding by uuid
func (cli *ZSClient) GetCandidateLdapEntryForIAM2Binding() (*view.GetLdapEntryView, error) {
	var resp view.GetLdapEntryView
	if err := cli.GetWithRespKey("v1/iam2/ldap/entries/candidates", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateIAM2VirtualIDLdapBinding creates IAM2VirtualIDLdapBinding
func (cli *ZSClient) CreateIAM2VirtualIDLdapBinding(params param.CreateIAM2VirtualIDLdapBindingParam) (*view.LdapResourceRefInventoryView, error) {
	resp := view.LdapResourceRefInventoryView{}
	if err := cli.Post("v1/iam2/ldap/bindings", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteIAM2VirtualIDLdapBinding deletes IAM2VirtualIDLdapBinding
func (cli *ZSClient) DeleteIAM2VirtualIDLdapBinding(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/iam2/ldap/bindings", uuid, string(deleteMode))
}
