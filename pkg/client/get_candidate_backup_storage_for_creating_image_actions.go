// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetCandidateBackupStorageForCreatingImage gets CandidateBackupStorageForCreatingImage by uuid
func (cli *ZSClient) GetCandidateBackupStorageForCreatingImage(uuid string) (*view.GetCandidateBackupStorageForCreatingImageView, error) {
	var resp view.GetCandidateBackupStorageForCreatingImageView
	if err := cli.Get("v1/images/candidate-backup-storage", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
