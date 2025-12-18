// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// PullSdnControllerTenant operates on PullSdnControllerTenant
func (cli *ZSClient) PullSdnControllerTenant(uuid string, params param.PullSdnControllerTenantParam) (*view.PullSdnControllerTenantEventView, error) {
	resp := view.PullSdnControllerTenantEventView{}
	if err := cli.Put("v1/sdn-controllers/{uuid}/tenant/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
