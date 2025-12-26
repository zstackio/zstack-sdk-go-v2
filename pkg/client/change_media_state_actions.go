// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeMediaState changes MediaState
func (cli *ZSClient) ChangeMediaState(uuid string, params param.ChangeMediaStateParam) (*view.ChangeMediaStateEventView, error) {
	resp := view.ChangeMediaStateEventView{}
	if err := cli.Put("v1/media/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
