// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ValidateSession 操作ValidateSession
func (cli *ZSClient) ValidateSession(params param.ValidateSessionParam) (*view.ValidateSessionView, error) {
	var resp view.ValidateSessionView
	if err := cli.Get("v1/accounts/sessions/{sessionUuid}/valid", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

