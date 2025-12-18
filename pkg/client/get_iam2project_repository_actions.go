// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetIAM2ProjectRepository gets IAM2ProjectRepository by uuid
func (cli *ZSClient) GetIAM2ProjectRepository(uuid string) (*view.GetIAM2ProjectRepositoryView, error) {
	var resp view.GetIAM2ProjectRepositoryView
	if err := cli.Get("v1/iam2/projects/repositories", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
