// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// QueryEventRecord queries EventRecord list
func (cli *ZSClient) QueryEventRecord(params *param.QueryParam) ([]view.EventRecordsInventoryView, error) {
	var resp []view.EventRecordsInventoryView
	return resp, cli.List("v1/zwatch/event-records", params, &resp)
}

func (cli *ZSClient) GetEventRecord(uuid string) (*view.EventRecordsInventoryView, error) {
	var resp view.EventRecordsInventoryView
	if err := cli.Get("v1/zwatch/event-records", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageEventRecord Pagination
func (cli *ZSClient) PageEventRecord(params *param.QueryParam) ([]view.EventRecordsInventoryView, int, error) {
	var eventRecords []view.EventRecordsInventoryView
	total, err := cli.Page("v1/zwatch/event-records", params, &eventRecords)
	return eventRecords, total, err
}
