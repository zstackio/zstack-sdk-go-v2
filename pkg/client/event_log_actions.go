// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryEventLog queries EventLog list
func (cli *ZSClient) QueryEventLog(params *param.QueryParam) ([]view.EventLogInventoryView, error) {
	var resp []view.EventLogInventoryView
	return resp, cli.List("v1/eventlogs", params, &resp)
}

func (cli *ZSClient) GetEventLog(uuid string) (*view.EventLogInventoryView, error) {
	var resp view.EventLogInventoryView
	if err := cli.Get("v1/eventlogs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
