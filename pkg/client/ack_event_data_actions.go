// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AckEventData operates on AckEventData
func (cli *ZSClient) AckEventData(params param.AckEventDataParam) (*view.AckAlertDataEventView, error) {
	resp := view.AckAlertDataEventView{}
	if err := cli.Post("v1/zwatch/event-histories/acknowledgments", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
