// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachPoliciesToUser operates on PoliciesToUser
func (cli *ZSClient) AttachPoliciesToUser(params param.AttachPoliciesToUserParam) (*view.AttachPoliciesToUserEventView, error) {
	resp := view.AttachPoliciesToUserEventView{}
	if err := cli.Post("v1/accounts/users/{userUuid}/policy-collection", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
