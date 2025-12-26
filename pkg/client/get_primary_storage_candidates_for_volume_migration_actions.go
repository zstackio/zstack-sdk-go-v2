// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetPrimaryStorageCandidatesForVolumeMigration gets PrimaryStorageCandidatesForVolumeMigration by uuid
func (cli *ZSClient) GetPrimaryStorageCandidatesForVolumeMigration(uuid string) (*view.GetPrimaryStorageCandidatesForVolumeMigrationView, error) {
	var resp view.GetPrimaryStorageCandidatesForVolumeMigrationView
	if err := cli.Get("v1/primary-storage/volumes/{volumeUuid}/migration-candidates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
