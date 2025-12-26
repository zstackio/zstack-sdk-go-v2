// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveCertificateFromLoadBalancerListener removes CertificateFromLoadBalancerListener
func (cli *ZSClient) RemoveCertificateFromLoadBalancerListener(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/listeners/{listenerUuid}/certificate", uuid, string(deleteMode))
}
