// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryGuestVmScriptExecutedRecordDetail queries GuestVmScriptExecutedRecordDetail list
func (cli *ZSClient) QueryGuestVmScriptExecutedRecordDetail(params *param.QueryParam) ([]view.GuestVmScriptExecutedRecordDetailInventoryView, error) {
	var resp []view.GuestVmScriptExecutedRecordDetailInventoryView
	return resp, cli.List("v1/scripts/records/details", params, &resp)
}

func (cli *ZSClient) GetGuestVmScriptExecutedRecordDetail(uuid string) (*view.GuestVmScriptExecutedRecordDetailInventoryView, error) {
	var resp view.GuestVmScriptExecutedRecordDetailInventoryView
	if err := cli.Get("v1/scripts/records/details", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageGuestVmScriptExecutedRecordDetail Pagination
func (cli *ZSClient) PageGuestVmScriptExecutedRecordDetail(params *param.QueryParam) ([]view.GuestVmScriptExecutedRecordDetailInventoryView, int, error) {
	var guestVmScriptExecutedRecordDetails []view.GuestVmScriptExecutedRecordDetailInventoryView
	total, err := cli.Page("v1/scripts/records/details", params, &guestVmScriptExecutedRecordDetails)
	return guestVmScriptExecutedRecordDetails, total, err
}
