// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetIAM2ProjectContainerImages 获取IAM2ProjectContainerImages详情
func (cli *ZSClient) GetIAM2ProjectContainerImages(uuid string) (*view.GetIAM2ProjectContainerImagesView, error) {
	var resp view.GetIAM2ProjectContainerImagesView
	if err := cli.Get("v1/iam2/project/{projectId}/repository/{repositoryId}/image", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

