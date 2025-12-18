// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateCSPSecretResourcePool 更新CSPSecretResourcePool
func (cli *ZSClient) UpdateCSPSecretResourcePool(uuid string, params param.UpdateCSPSecretResourcePoolParam) (*view.UpdateSecretResourcePoolEventView, error) {
	resp := view.UpdateSecretResourcePoolEventView{}
	if err := cli.Put("v1/secret-resource-pools/csp/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

