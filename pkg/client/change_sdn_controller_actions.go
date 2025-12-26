// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeSdnController changes SdnController
func (cli *ZSClient) ChangeSdnController(uuid string, params param.ChangeSdnControllerParam) (*view.ChangeSdnControllerEventView, error) {
	resp := view.ChangeSdnControllerEventView{}
	if err := cli.Put("v1/sdn-controllers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
