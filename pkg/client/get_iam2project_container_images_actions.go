// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetIAM2ProjectContainerImages gets IAM2ProjectContainerImages by uuid
func (cli *ZSClient) GetIAM2ProjectContainerImages(uuid string) (*view.GetIAM2ProjectContainerImagesView, error) {
	var resp view.GetIAM2ProjectContainerImagesView
	if err := cli.Get("v1/iam2/project/{projectId}/repository/{repositoryId}/image", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
