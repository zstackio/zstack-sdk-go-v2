// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateAlertDataAck updates AlertDataAck
func (cli *ZSClient) UpdateAlertDataAck(uuid string, params param.UpdateAlertDataAckParam) (*view.UpdateAlertDataAckEventView, error) {
	resp := view.UpdateAlertDataAckEventView{}
	if err := cli.Put("v1/zwatch/alert-histories/acknowledgments/{alertDataUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
