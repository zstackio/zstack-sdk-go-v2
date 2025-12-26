// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateEip updates Eip
func (cli *ZSClient) UpdateEip(uuid string, params param.UpdateEipParam) (*view.UpdateEipEventView, error) {
	resp := view.UpdateEipEventView{}
	if err := cli.Put("v1/eips/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
