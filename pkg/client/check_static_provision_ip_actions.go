// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CheckStaticProvisionIp 操作CheckStaticProvisionIp
func (cli *ZSClient) CheckStaticProvisionIp(params param.CheckStaticProvisionIpParam) (*view.CheckStaticProvisionIpView, error) {
	resp := view.CheckStaticProvisionIpView{}
	if err := cli.Post("v1/baremetal2/bm-instances/static/provision/ip/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

