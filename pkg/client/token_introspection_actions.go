// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// TokenIntrospection operates on TokenIntrospection
func (cli *ZSClient) TokenIntrospection(params param.TokenIntrospectionParam) (*view.TokenIntrospectionView, error) {
	resp := view.TokenIntrospectionView{}
	if err := cli.Post("v1/token/introspect", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
