// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddSNSSmsReceiver 操作AddSNSSmsReceiver
func (cli *ZSClient) AddSNSSmsReceiver(params param.AddSNSSmsReceiverParam) (*view.AddSNSSmsReceiverEventView, error) {
	resp := view.AddSNSSmsReceiverEventView{}
	if err := cli.Post("v1/sns/sms-endpoints/receivers", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

