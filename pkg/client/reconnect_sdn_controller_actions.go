// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ReconnectSdnController 操作ReconnectSdnController
func (cli *ZSClient) ReconnectSdnController(uuid string, params param.ReconnectSdnControllerParam) (*view.ReconnectSdnControllerEventView, error) {
	resp := view.ReconnectSdnControllerEventView{}
	if err := cli.Put("v1/sdn-controllers/{sdnControllerUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

