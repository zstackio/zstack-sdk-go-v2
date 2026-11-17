// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// GetAuditData gets AuditData by uuid
func (cli *ZSClient) GetAuditData() (*view.GetAuditDataView, error) {
	var resp view.GetAuditDataView
	if err := cli.GetWithRespKey("v1/zwatch/audits", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryAudit queries Audit list
func (cli *ZSClient) QueryAudit(params *param.QueryParam) ([]view.AuditsInventoryView, error) {
	var resp []view.AuditsInventoryView
	return resp, cli.List("v1/zwatch/audit-records", params, &resp)
}

func (cli *ZSClient) GetAudit(uuid string) (*view.AuditsInventoryView, error) {
	var resp view.AuditsInventoryView
	if err := cli.Get("v1/zwatch/audit-records", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAudit Pagination
func (cli *ZSClient) PageAudit(params *param.QueryParam) ([]view.AuditsInventoryView, int, error) {
	var audits []view.AuditsInventoryView
	total, err := cli.Page("v1/zwatch/audit-records", params, &audits)
	return audits, total, err
}
