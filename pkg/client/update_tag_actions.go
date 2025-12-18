// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateTag updates Tag
func (cli *ZSClient) UpdateTag(uuid string, params param.UpdateTagParam) (*view.UpdateTagEventView, error) {
	resp := view.UpdateTagEventView{}
	if err := cli.Put("v1/tags/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
