// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateSystemTag updates SystemTag
func (cli *ZSClient) UpdateSystemTag(uuid string, params param.UpdateSystemTagParam) (*view.UpdateSystemTagEventView, error) {
	resp := view.UpdateSystemTagEventView{}
	if err := cli.Put("v1/system-tags/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
