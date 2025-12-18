// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// StartConnectionBetweenAliyunRouterInterface starts ConnectionBetweenAliyunRouterInterface
func (cli *ZSClient) StartConnectionBetweenAliyunRouterInterface(uuid string, params param.StartConnectionBetweenAliyunRouterInterfaceParam) (*view.StartConnectionBetweenAliyunRouterInterfaceEventView, error) {
	resp := view.StartConnectionBetweenAliyunRouterInterfaceEventView{}
	if err := cli.Put("v1/hybrid/aliyun/router-interface/{vbrInterfaceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
