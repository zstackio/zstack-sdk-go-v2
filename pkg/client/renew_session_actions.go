// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RenewSession 操作RenewSession
func (cli *ZSClient) RenewSession(uuid string, params param.RenewSessionParam) (*view.RenewSessionEventView, error) {
	resp := view.RenewSessionEventView{}
	if err := cli.Put("v1/accounts/sessions/{sessionUuid}/renew", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

