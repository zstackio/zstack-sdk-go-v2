// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// LogInByUser operates on LogInByUser
func (cli *ZSClient) LogInByUser(uuid string, params param.LogInByUserParam) (*view.LogInView, error) {
	resp := view.LogInView{}
	if err := cli.Put("v1/accounts/users/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
