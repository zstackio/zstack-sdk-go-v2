// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ValidatePassword 操作ValidatePassword
func (cli *ZSClient) ValidatePassword(uuid string, params param.ValidatePasswordParam) (*view.ValidatePasswordView, error) {
	resp := view.ValidatePasswordView{}
	if err := cli.Put("v1/password/verify", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

