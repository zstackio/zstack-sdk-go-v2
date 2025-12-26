// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateSdnController updates SdnController
func (cli *ZSClient) UpdateSdnController(uuid string, params param.UpdateSdnControllerParam) (*view.UpdateSdnControllerEventView, error) {
	resp := view.UpdateSdnControllerEventView{}
	if err := cli.Put("v1/sdn-controllers/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
