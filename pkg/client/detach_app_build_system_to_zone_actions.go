// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachAppBuildSystemToZone operates on AppBuildSystemToZone
func (cli *ZSClient) DetachAppBuildSystemToZone(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zones/{zoneUuid}/buildsystem/{buildSystemUuid}", uuid, string(deleteMode))
}
