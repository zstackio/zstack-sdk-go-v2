// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCandidateLdapEntryForIAM2Binding 获取CandidateLdapEntryForIAM2Binding详情
func (cli *ZSClient) GetCandidateLdapEntryForIAM2Binding(uuid string) (*view.GetLdapEntryView, error) {
	var resp view.GetLdapEntryView
	if err := cli.Get("v1/iam2/ldap/entries/candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

