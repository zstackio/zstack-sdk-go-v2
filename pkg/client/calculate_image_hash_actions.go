// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CalculateImageHash operates on CalculateImageHash
func (cli *ZSClient) CalculateImageHash(uuid string, params param.CalculateImageHashParam) (*view.CalculateImageHashEventView, error) {
	resp := view.CalculateImageHashEventView{}
	if err := cli.Put("v1/images/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
