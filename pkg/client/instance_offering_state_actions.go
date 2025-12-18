// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeInstanceOfferingState 操作InstanceOfferingState
func (cli *ZSClient) ChangeInstanceOfferingState(uuid string, params param.ChangeInstanceOfferingStateParam) (*view.ChangeInstanceOfferingStateEventView, error) {
	resp := view.ChangeInstanceOfferingStateEventView{}
	if err := cli.Put("v1/instance-offerings/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

