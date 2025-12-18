// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddSdnController 操作AddSdnController
func (cli *ZSClient) AddSdnController(params param.AddSdnControllerParam) (*view.AddSdnControllerEventView, error) {
	resp := view.AddSdnControllerEventView{}
	if err := cli.Post("v1/sdn-controllers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

