// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateCasClient updates CasClient
func (cli *ZSClient) UpdateCasClient(uuid string, params param.UpdateCasClientParam) (*view.UpdateCasClientEventView, error) {
	resp := view.UpdateCasClientEventView{}
	if err := cli.Put("v1/update/cas/client", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
