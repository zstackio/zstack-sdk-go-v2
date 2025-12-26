// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachAppBuildSystemToZone operates on AppBuildSystemToZone
func (cli *ZSClient) AttachAppBuildSystemToZone(params param.AttachAppBuildSystemToZoneParam) (*view.AttachAppBuildSystemToZoneEventView, error) {
	resp := view.AttachAppBuildSystemToZoneEventView{}
	if err := cli.Post("v1/zones/{zoneUuid}/buildsystem/{buildSystemUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
