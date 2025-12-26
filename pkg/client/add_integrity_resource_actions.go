// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddIntegrityResource adds IntegrityResource
func (cli *ZSClient) AddIntegrityResource(params param.AddIntegrityResourceParam) (*view.AddIntegrityResourceEventView, error) {
	resp := view.AddIntegrityResourceEventView{}
	if err := cli.Post("v1/integrity/resource/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
