// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateSdnController 更新SdnController
func (cli *ZSClient) UpdateSdnController(uuid string, params param.UpdateSdnControllerParam) (*view.UpdateSdnControllerEventView, error) {
	resp := view.UpdateSdnControllerEventView{}
	if err := cli.Put("v1/sdn-controllers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

