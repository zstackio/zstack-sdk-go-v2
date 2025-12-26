// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddSNSSmsReceiver adds SNSSmsReceiver
func (cli *ZSClient) AddSNSSmsReceiver(params param.AddSNSSmsReceiverParam) (*view.AddSNSSmsReceiverEventView, error) {
	resp := view.AddSNSSmsReceiverEventView{}
	if err := cli.Post("v1/sns/sms-endpoints/receivers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
