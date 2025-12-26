// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// LoginByCas operates on LoginByCas
func (cli *ZSClient) LoginByCas(uuid string, params param.LoginByCasParam) (*view.LoginByCasView, error) {
	resp := view.LoginByCasView{}
	if err := cli.Put("v1/cas/login/", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
