// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetAuditData gets AuditData by uuid
func (cli *ZSClient) GetAuditData(uuid string) (*view.GetAuditDataView, error) {
	var resp view.GetAuditDataView
	if err := cli.Get("v1/zwatch/audits", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
