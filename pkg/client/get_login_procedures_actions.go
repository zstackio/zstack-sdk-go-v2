// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetLoginProcedures gets LoginProcedures by uuid
func (cli *ZSClient) GetLoginProcedures(uuid string) (*view.GetLoginProceduresView, error) {
	var resp view.GetLoginProceduresView
	if err := cli.Get("v1/login/procedures", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
