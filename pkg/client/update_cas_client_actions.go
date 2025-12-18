// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateCasClient updates CasClient
func (cli *ZSClient) UpdateCasClient(uuid string, params param.UpdateCasClientParam) (*view.UpdateCasClientEventView, error) {
	resp := view.UpdateCasClientEventView{}
	if err := cli.Put("v1/update/cas/client", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
