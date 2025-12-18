// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateSlbInstance 创建SlbInstance
func (cli *ZSClient) CreateSlbInstance(params param.CreateSlbInstanceParam) (*view.CreateSlbInstanceEventView, error) {
	resp := view.CreateSlbInstanceEventView{}
	if err := cli.Post("v1/load-balancers/slb/instances", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

