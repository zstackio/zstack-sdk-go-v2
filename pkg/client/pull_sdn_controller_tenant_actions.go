// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// PullSdnControllerTenant operates on PullSdnControllerTenant
func (cli *ZSClient) PullSdnControllerTenant(uuid string, params param.PullSdnControllerTenantParam) (*view.PullSdnControllerTenantEventView, error) {
	resp := view.PullSdnControllerTenantEventView{}
	if err := cli.Put("v1/sdn-controllers/{uuid}/tenant/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
