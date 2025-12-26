// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SdnControllerAddHost operates on SdnControllerAddHost
func (cli *ZSClient) SdnControllerAddHost(params param.SdnControllerAddHostParam) (*view.SdnControllerAddHostEventView, error) {
	resp := view.SdnControllerAddHostEventView{}
	if err := cli.Post("v1/sdn-controllers/{sdnControllerUuid}/hosts/{hostUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
