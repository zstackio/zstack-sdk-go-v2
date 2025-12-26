// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ReconnectSdnController operates on ReconnectSdnController
func (cli *ZSClient) ReconnectSdnController(uuid string, params param.ReconnectSdnControllerParam) (*view.ReconnectSdnControllerEventView, error) {
	resp := view.ReconnectSdnControllerEventView{}
	if err := cli.Put("v1/sdn-controllers/{sdnControllerUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
