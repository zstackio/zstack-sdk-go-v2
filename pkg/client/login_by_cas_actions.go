// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// LoginByCas 操作LoginByCas
func (cli *ZSClient) LoginByCas(uuid string, params param.LoginByCasParam) (*view.LoginByCasView, error) {
	resp := view.LoginByCasView{}
	if err := cli.Put("v1/cas/login/", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

