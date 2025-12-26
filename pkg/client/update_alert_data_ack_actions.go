// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateAlertDataAck updates AlertDataAck
func (cli *ZSClient) UpdateAlertDataAck(uuid string, params param.UpdateAlertDataAckParam) (*view.UpdateAlertDataAckEventView, error) {
	resp := view.UpdateAlertDataAckEventView{}
	if err := cli.Put("v1/zwatch/alert-histories/acknowledgments/{alertDataUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
