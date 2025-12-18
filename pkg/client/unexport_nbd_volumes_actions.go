// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UnexportNbdVolumes 操作UnexportNbdVolumes
func (cli *ZSClient) UnexportNbdVolumes(params param.UnexportNbdVolumesParam) (*view.UnexportNbdVolumesEventView, error) {
	resp := view.UnexportNbdVolumesEventView{}
	if err := cli.Post("v1/cbt-task/unexportvolume", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

