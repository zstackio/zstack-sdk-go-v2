// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetLicenseRecords gets LicenseRecords by uuid
func (cli *ZSClient) GetLicenseRecords(uuid string) (*view.GetLicenseRecordsView, error) {
	var resp view.GetLicenseRecordsView
	if err := cli.Get("v1/licenses/records", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
