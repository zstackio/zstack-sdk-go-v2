// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CalculateImageHash operates on CalculateImageHash
func (cli *ZSClient) CalculateImageHash(uuid string, params param.CalculateImageHashParam) (*view.CalculateImageHashEventView, error) {
	resp := view.CalculateImageHashEventView{}
	if err := cli.Put("v1/images/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
