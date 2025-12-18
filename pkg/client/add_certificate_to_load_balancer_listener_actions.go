// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddCertificateToLoadBalancerListener 操作AddCertificateToLoadBalancerListener
func (cli *ZSClient) AddCertificateToLoadBalancerListener(params param.AddCertificateToLoadBalancerListenerParam) (*view.AddCertificateToLoadBalancerListenerEventView, error) {
	resp := view.AddCertificateToLoadBalancerListenerEventView{}
	if err := cli.Post("v1/load-balancers/listeners/{listenerUuid}/certificate", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

