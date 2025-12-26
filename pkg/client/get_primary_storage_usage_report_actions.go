// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetPrimaryStorageUsageReport gets PrimaryStorageUsageReport by uuid
func (cli *ZSClient) GetPrimaryStorageUsageReport(uuid string) (*view.GetPrimaryStorageUsageReportView, error) {
	var resp view.GetPrimaryStorageUsageReportView
	if err := cli.Get("v1/primary-storage/{primaryStorageUuid}/usage/report", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
