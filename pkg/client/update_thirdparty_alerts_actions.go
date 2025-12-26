// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateThirdpartyAlerts updates ThirdpartyAlerts
func (cli *ZSClient) UpdateThirdpartyAlerts(uuid string, params param.UpdateThirdpartyAlertsParam) (*view.UpdateThirdpartyAlertsEventView, error) {
	resp := view.UpdateThirdpartyAlertsEventView{}
	if err := cli.Put("v1/zwatch/third-party/alerts/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
