// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// LogIn operates on LogIn
func (cli *ZSClient) LogIn(uuid string, params param.LogInParam) (*view.LogInView, error) {
	resp := view.LogInView{}
	if err := cli.Put("v1/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
