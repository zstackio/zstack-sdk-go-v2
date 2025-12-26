// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetVmSoundType operates on SetVmSoundType
func (cli *ZSClient) SetVmSoundType(uuid string, params param.SetVmSoundTypeParam) (*view.SetVmSoundTypeEventView, error) {
	resp := view.SetVmSoundTypeEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
