// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ExportNbdVolumes operates on ExportNbdVolumes
func (cli *ZSClient) ExportNbdVolumes(params param.ExportNbdVolumesParam) (*view.ExportNbdVolumesEventView, error) {
	resp := view.ExportNbdVolumesEventView{}
	if err := cli.Post("v1/cbt-task/exportvolume", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
