// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryGuestVmScriptExecutedRecord queries GuestVmScriptExecutedRecord list
func (cli *ZSClient) QueryGuestVmScriptExecutedRecord(ctx context.Context, params *param.QueryParam) ([]view.GuestVmScriptExecutedRecordInventoryView, error) {
	var resp []view.GuestVmScriptExecutedRecordInventoryView
	return resp, cli.List(ctx, "v1/scripts/records", params, &resp)
}

func (cli *ZSClient) GetGuestVmScriptExecutedRecord(ctx context.Context, uuid string) (*view.GuestVmScriptExecutedRecordInventoryView, error) {
	var resp view.GuestVmScriptExecutedRecordInventoryView
	if err := cli.Get(ctx, "v1/scripts/records", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageGuestVmScriptExecutedRecord Pagination
func (cli *ZSClient) PageGuestVmScriptExecutedRecord(ctx context.Context, params *param.QueryParam) ([]view.GuestVmScriptExecutedRecordInventoryView, int, error) {
	var guestVmScriptExecutedRecords []view.GuestVmScriptExecutedRecordInventoryView
	total, err := cli.Page(ctx, "v1/scripts/records", params, &guestVmScriptExecutedRecords)
	return guestVmScriptExecutedRecords, total, err
}
