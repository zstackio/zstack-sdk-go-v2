// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetIAM2ProjectContainerImageTags 获取IAM2ProjectContainerImageTags详情
func (cli *ZSClient) GetIAM2ProjectContainerImageTags(uuid string) (*view.GetIAM2ProjectContainerImageTagsView, error) {
	var resp view.GetIAM2ProjectContainerImageTagsView
	if err := cli.Get("v1/iam2/project/{projectId}/repository/{repositoryId}/image/{imageName}/tag", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

