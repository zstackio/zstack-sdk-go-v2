// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryGuestVmScriptExecutedRecord queries GuestVmScriptExecutedRecord list
func (cli *ZSClient) QueryGuestVmScriptExecutedRecord(params *param.QueryParam) ([]view.GuestVmScriptExecutedRecordInventoryView, error) {
	var resp []view.GuestVmScriptExecutedRecordInventoryView
	return resp, cli.List("v1/scripts/records", params, &resp)
}

func (cli *ZSClient) GetGuestVmScriptExecutedRecord(uuid string) (*view.GuestVmScriptExecutedRecordInventoryView, error) {
	var resp view.GuestVmScriptExecutedRecordInventoryView
	if err := cli.Get("v1/scripts/records", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
