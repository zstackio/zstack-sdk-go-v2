// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachAppBuildSystemToZone operates on AppBuildSystemToZone
func (cli *ZSClient) AttachAppBuildSystemToZone(params param.AttachAppBuildSystemToZoneParam) (*view.AttachAppBuildSystemToZoneEventView, error) {
	resp := view.AttachAppBuildSystemToZoneEventView{}
	if err := cli.Post("v1/zones/{zoneUuid}/buildsystem/{buildSystemUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
