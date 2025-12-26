// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UnexportNbdVolumes operates on UnexportNbdVolumes
func (cli *ZSClient) UnexportNbdVolumes(params param.UnexportNbdVolumesParam) (*view.UnexportNbdVolumesEventView, error) {
	resp := view.UnexportNbdVolumesEventView{}
	if err := cli.Post("v1/cbt-task/unexportvolume", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
