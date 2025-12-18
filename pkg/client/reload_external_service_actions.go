// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ReloadExternalService 操作ReloadExternalService
func (cli *ZSClient) ReloadExternalService(uuid string, params param.ReloadExternalServiceParam) (*view.ReloadExternalServiceEventView, error) {
	resp := view.ReloadExternalServiceEventView{}
	if err := cli.Put("v1/external/services", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

