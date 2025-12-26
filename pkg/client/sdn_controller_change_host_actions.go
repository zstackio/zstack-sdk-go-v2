// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SdnControllerChangeHost operates on SdnControllerChangeHost
func (cli *ZSClient) SdnControllerChangeHost(uuid string, params param.SdnControllerChangeHostParam) (*view.SdnControllerChangeHostEventView, error) {
	resp := view.SdnControllerChangeHostEventView{}
	if err := cli.Put("v1/sdn-controllers/{sdnControllerUuid}/hosts/{hostUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
