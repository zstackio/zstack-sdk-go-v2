// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CheckStaticProvisionIp operates on CheckStaticProvisionIp
func (cli *ZSClient) CheckStaticProvisionIp(params param.CheckStaticProvisionIpParam) (*view.CheckStaticProvisionIpView, error) {
	resp := view.CheckStaticProvisionIpView{}
	if err := cli.Post("v1/baremetal2/bm-instances/static/provision/ip/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
