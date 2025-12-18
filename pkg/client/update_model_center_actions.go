// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateModelCenter updates ModelCenter
func (cli *ZSClient) UpdateModelCenter(uuid string, params param.UpdateModelCenterParam) (*view.UpdateModelCenterEventView, error) {
	resp := view.UpdateModelCenterEventView{}
	if err := cli.Put("v1/ai/model-centers/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
