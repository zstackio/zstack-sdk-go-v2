// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// TokenIntrospection 操作TokenIntrospection
func (cli *ZSClient) TokenIntrospection(params param.TokenIntrospectionParam) (*view.TokenIntrospectionView, error) {
	resp := view.TokenIntrospectionView{}
	if err := cli.Post("v1/token/introspect", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

