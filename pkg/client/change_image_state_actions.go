// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeImageState changes ImageState
func (cli *ZSClient) ChangeImageState(uuid string, params param.ChangeImageStateParam) (*view.ChangeImageStateEventView, error) {
	resp := view.ChangeImageStateEventView{}
	if err := cli.Put("v1/images/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
