// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetLdapEntry 获取LdapEntry详情
func (cli *ZSClient) GetLdapEntry(uuid string) (*view.GetLdapEntryView, error) {
	var resp view.GetLdapEntryView
	if err := cli.Get("v1/ldap/entry", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

