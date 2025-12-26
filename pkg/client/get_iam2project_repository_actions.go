// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetIAM2ProjectRepository gets IAM2ProjectRepository by uuid
func (cli *ZSClient) GetIAM2ProjectRepository(uuid string) (*view.GetIAM2ProjectRepositoryView, error) {
	var resp view.GetIAM2ProjectRepositoryView
	if err := cli.Get("v1/iam2/projects/repositories", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
