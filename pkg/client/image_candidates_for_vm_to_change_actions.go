// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetImageCandidatesForVmToChange 获取ImageCandidatesForVmToChange详情
func (cli *ZSClient) GetImageCandidatesForVmToChange(uuid string) (*view.GetImageCandidatesForVmToChangeView, error) {
	var resp view.GetImageCandidatesForVmToChangeView
	if err := cli.Get("v1/vm-instances/{vmInstanceUuid}/image-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

