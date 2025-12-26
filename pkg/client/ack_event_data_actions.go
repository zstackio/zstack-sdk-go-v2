// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AckEventData operates on AckEventData
func (cli *ZSClient) AckEventData(params param.AckEventDataParam) (*view.AckAlertDataEventView, error) {
	resp := view.AckAlertDataEventView{}
	if err := cli.Post("v1/zwatch/event-histories/acknowledgments", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
