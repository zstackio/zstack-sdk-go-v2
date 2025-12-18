// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddProxyToResource 操作AddProxyToResource
func (cli *ZSClient) AddProxyToResource(params param.AddProxyToResourceParam) (*view.AddProxyToResourceEventView, error) {
	resp := view.AddProxyToResourceEventView{}
	if err := cli.Post("v1/proxy/{proxyUuid}/resource/{resourceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

