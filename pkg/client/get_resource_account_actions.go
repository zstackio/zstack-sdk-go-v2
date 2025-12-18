// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetResourceAccount gets ResourceAccount by uuid
func (cli *ZSClient) GetResourceAccount(uuid string) (*view.GetResourceAccountView, error) {
	var resp view.GetResourceAccountView
	if err := cli.Get("v1/resources/accounts", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
