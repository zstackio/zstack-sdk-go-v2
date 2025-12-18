// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RefreshSSOServerToken 操作RefreshSSOServerToken
func (cli *ZSClient) RefreshSSOServerToken(uuid string, params param.RefreshSSOServerTokenParam) (*view.RefreshSSOServerTokenEventView, error) {
	resp := view.RefreshSSOServerTokenEventView{}
	if err := cli.Put("v1/sso/server/token/refresh", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

