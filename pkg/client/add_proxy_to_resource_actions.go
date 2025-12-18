// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddProxyToResource adds ProxyToResource
func (cli *ZSClient) AddProxyToResource(params param.AddProxyToResourceParam) (*view.AddProxyToResourceEventView, error) {
	resp := view.AddProxyToResourceEventView{}
	if err := cli.Post("v1/proxy/{proxyUuid}/resource/{resourceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
