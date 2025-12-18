// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetCandidateBackupStorageForCreatingImage 获取CandidateBackupStorageForCreatingImage详情
func (cli *ZSClient) GetCandidateBackupStorageForCreatingImage(uuid string) (*view.GetCandidateBackupStorageForCreatingImageView, error) {
	var resp view.GetCandidateBackupStorageForCreatingImageView
	if err := cli.Get("v1/images/candidate-backup-storage", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

