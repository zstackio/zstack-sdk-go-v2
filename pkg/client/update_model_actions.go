// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateModel updates Model
func (cli *ZSClient) UpdateModel(uuid string, params param.UpdateModelParam) (*view.UpdateModelEventView, error) {
	resp := view.UpdateModelEventView{}
	if err := cli.Put("v1/ai/models/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
