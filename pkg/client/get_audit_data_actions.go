// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetAuditData gets AuditData by uuid
func (cli *ZSClient) GetAuditData(uuid string) (*view.GetAuditDataView, error) {
	var resp view.GetAuditDataView
	if err := cli.Get("v1/zwatch/audits", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
