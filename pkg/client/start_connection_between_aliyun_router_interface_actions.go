// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// StartConnectionBetweenAliyunRouterInterface starts ConnectionBetweenAliyunRouterInterface
func (cli *ZSClient) StartConnectionBetweenAliyunRouterInterface(uuid string, params param.StartConnectionBetweenAliyunRouterInterfaceParam) (*view.StartConnectionBetweenAliyunRouterInterfaceEventView, error) {
	resp := view.StartConnectionBetweenAliyunRouterInterfaceEventView{}
	if err := cli.Put("v1/hybrid/aliyun/router-interface/{vbrInterfaceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
