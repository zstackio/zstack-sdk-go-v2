// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateThirdpartyAlerts 更新ThirdpartyAlerts
func (cli *ZSClient) UpdateThirdpartyAlerts(uuid string, params param.UpdateThirdpartyAlertsParam) (*view.UpdateThirdpartyAlertsEventView, error) {
	resp := view.UpdateThirdpartyAlertsEventView{}
	if err := cli.Put("v1/zwatch/third-party/alerts/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

