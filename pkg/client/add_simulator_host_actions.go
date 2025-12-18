// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddSimulatorHost 操作AddSimulatorHost
func (cli *ZSClient) AddSimulatorHost(params param.AddSimulatorHostParam) (*view.AddHostEventView, error) {
	resp := view.AddHostEventView{}
	if err := cli.Post("v1/hosts/simulators", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

