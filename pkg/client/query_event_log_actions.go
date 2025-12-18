// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryEventLog queries EventLog list
func (cli *ZSClient) QueryEventLog(params param.QueryParam) ([]view.EventLogInventoryView, error) {
	var resp []view.EventLogInventoryView
	return resp, cli.List("v1/eventlogs", &params, &resp)
}
