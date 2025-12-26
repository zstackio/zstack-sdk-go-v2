// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeVolumeState changes VolumeState
func (cli *ZSClient) ChangeVolumeState(uuid string, params param.ChangeVolumeStateParam) (*view.ChangeVolumeStateEventView, error) {
	resp := view.ChangeVolumeStateEventView{}
	if err := cli.Put("v1/volumes/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
