// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateIAM2VirtualIDFromLdapUid creates IAM2VirtualIDFromLdapUid
func (cli *ZSClient) CreateIAM2VirtualIDFromLdapUid(params param.CreateIAM2VirtualIDFromLdapUidParam) (*view.CreateIAM2VirtualIDFromLdapUidEventView, error) {
	resp := view.CreateIAM2VirtualIDFromLdapUidEventView{}
	if err := cli.Post("v1/iam2/virtual-id/ldap/uid", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
