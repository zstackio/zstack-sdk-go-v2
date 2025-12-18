// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveCertificateFromLoadBalancerListener 操作RemoveCertificateFromLoadBalancerListener
func (cli *ZSClient) RemoveCertificateFromLoadBalancerListener(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/listeners/{listenerUuid}/certificate", uuid, string(deleteMode))
}

