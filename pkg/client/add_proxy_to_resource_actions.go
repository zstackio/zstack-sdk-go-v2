// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddProxyToResource adds ProxyToResource
func (cli *ZSClient) AddProxyToResource(params param.AddProxyToResourceParam) (*view.AddProxyToResourceEventView, error) {
	resp := view.AddProxyToResourceEventView{}
	if err := cli.Post("v1/proxy/{proxyUuid}/resource/{resourceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
