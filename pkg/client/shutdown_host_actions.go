// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ShutdownHost operates on ShutdownHost
func (cli *ZSClient) ShutdownHost(uuid string, params param.ShutdownHostParam) (*view.ShutdownHostEventView, error) {
	resp := view.ShutdownHostEventView{}
	if err := cli.Put("v1/hosts/power/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
