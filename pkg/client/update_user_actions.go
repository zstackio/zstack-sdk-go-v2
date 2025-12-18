// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateUser updates User
func (cli *ZSClient) UpdateUser(uuid string, params param.UpdateUserParam) (*view.UpdateUserEventView, error) {
	resp := view.UpdateUserEventView{}
	if err := cli.Put("v1/accounts/users/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
