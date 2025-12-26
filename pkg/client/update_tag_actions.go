// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateTag updates Tag
func (cli *ZSClient) UpdateTag(uuid string, params param.UpdateTagParam) (*view.UpdateTagEventView, error) {
	resp := view.UpdateTagEventView{}
	if err := cli.Put("v1/tags/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
