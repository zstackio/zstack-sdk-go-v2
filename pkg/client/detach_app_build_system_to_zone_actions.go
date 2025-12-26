// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DetachAppBuildSystemToZone operates on AppBuildSystemToZone
func (cli *ZSClient) DetachAppBuildSystemToZone(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zones/{zoneUuid}/buildsystem/{buildSystemUuid}", uuid, string(deleteMode))
}
