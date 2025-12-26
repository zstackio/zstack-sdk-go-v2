// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RefreshSSOServerToken operates on RefreshSSOServerToken
func (cli *ZSClient) RefreshSSOServerToken(uuid string, params param.RefreshSSOServerTokenParam) (*view.RefreshSSOServerTokenEventView, error) {
	resp := view.RefreshSSOServerTokenEventView{}
	if err := cli.Put("v1/sso/server/token/refresh", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
