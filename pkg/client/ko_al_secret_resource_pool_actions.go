// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateKoAlSecretResourcePool 更新KoAlSecretResourcePool
func (cli *ZSClient) UpdateKoAlSecretResourcePool(uuid string, params param.UpdateKoAlSecretResourcePoolParam) (*view.UpdateSecretResourcePoolEventView, error) {
	resp := view.UpdateSecretResourcePoolEventView{}
	if err := cli.Put("v1/secret-resource-pools/koal/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

