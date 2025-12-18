// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateFlkSecSecretResourcePool 更新FlkSecSecretResourcePool
func (cli *ZSClient) UpdateFlkSecSecretResourcePool(uuid string, params param.UpdateFlkSecSecretResourcePoolParam) (*view.UpdateSecretResourcePoolEventView, error) {
	resp := view.UpdateSecretResourcePoolEventView{}
	if err := cli.Put("v1/secret-resource-pools/flkSec/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

