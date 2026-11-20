// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateLoadBalancerListener updates LoadBalancerListener
func (cli *ZSClient) UpdateLoadBalancerListener(uuid string, params param.UpdateLoadBalancerListenerParam) (*view.LoadBalancerListenerInventoryView, error) {
	resp := view.LoadBalancerListenerInventoryView{}
	if err := cli.PutWithRespKey("v1/load-balancers/listeners", uuid, "", map[string]interface{}{
		"updateLoadBalancerListener": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateLoadBalancerListener creates LoadBalancerListener
func (cli *ZSClient) CreateLoadBalancerListener(loadBalancerUuid string, params param.CreateLoadBalancerListenerParam) (*view.LoadBalancerListenerInventoryView, error) {
	resp := view.LoadBalancerListenerInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/load-balancers/%s/listeners", loadBalancerUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryLoadBalancerListener queries LoadBalancerListener list
func (cli *ZSClient) QueryLoadBalancerListener(params *param.QueryParam) ([]view.LoadBalancerListenerInventoryView, error) {
	var resp []view.LoadBalancerListenerInventoryView
	return resp, cli.List("v1/load-balancers/listeners", params, &resp)
}

func (cli *ZSClient) GetLoadBalancerListener(uuid string) (*view.LoadBalancerListenerInventoryView, error) {
	var resp view.LoadBalancerListenerInventoryView
	if err := cli.Get("v1/load-balancers/listeners", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageLoadBalancerListener Pagination
func (cli *ZSClient) PageLoadBalancerListener(params *param.QueryParam) ([]view.LoadBalancerListenerInventoryView, int, error) {
	var loadBalancerListeners []view.LoadBalancerListenerInventoryView
	total, err := cli.Page("v1/load-balancers/listeners", params, &loadBalancerListeners)
	return loadBalancerListeners, total, err
}
// ChangeLoadBalancerListener changes LoadBalancerListener
func (cli *ZSClient) ChangeLoadBalancerListener(uuid string, params param.ChangeLoadBalancerListenerParam) (*view.LoadBalancerListenerInventoryView, error) {
	resp := view.LoadBalancerListenerInventoryView{}
	if err := cli.PutWithRespKey("v1/load-balancers/listeners", uuid, "", map[string]interface{}{
		"changeLoadBalancerListener": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteLoadBalancerListener deletes LoadBalancerListener
func (cli *ZSClient) DeleteLoadBalancerListener(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/listeners", uuid, string(deleteMode))
}

// AddAccessControlListToLoadBalancer adds AccessControlListToLoadBalancer
func (cli *ZSClient) AddAccessControlListToLoadBalancer(listenerUuid string, params param.AddAccessControlListToLoadBalancerParam) (*view.LoadBalancerListenerInventoryView, error) {
	resp := view.LoadBalancerListenerInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/load-balancers/listeners/%s/access-control-lists", listenerUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveServerGroupFromLoadBalancerListener requires additional arguments.
//
// Deprecated: use RemoveServerGroupFromLoadBalancerListenerWithParams to supply the required fields.
func (cli *ZSClient) RemoveServerGroupFromLoadBalancerListener(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("RemoveServerGroupFromLoadBalancerListener requires additional arguments; use RemoveServerGroupFromLoadBalancerListenerWithParams")
}

// RemoveServerGroupFromLoadBalancerListenerWithParams sends the required route and query parameters for RemoveServerGroupFromLoadBalancerListener.
func (cli *ZSClient) RemoveServerGroupFromLoadBalancerListenerWithParams(listenerUuid string, deleteMode param.DeleteMode, params param.RemoveServerGroupFromLoadBalancerListenerParamDetail) error {
	query := url.Values{
		"serverGroupUuid": {params.ServerGroupUuid},
	}
	return cli.deleteActionWithQuery("v1/load-balancers/listeners", listenerUuid, "servergroups", deleteMode, query, "serverGroupUuid")
}

// RemoveCertificateFromLoadBalancerListener requires additional arguments.
//
// Deprecated: use RemoveCertificateFromLoadBalancerListenerWithParams to supply the required fields.
func (cli *ZSClient) RemoveCertificateFromLoadBalancerListener(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("RemoveCertificateFromLoadBalancerListener requires additional arguments; use RemoveCertificateFromLoadBalancerListenerWithParams")
}

// RemoveCertificateFromLoadBalancerListenerWithParams sends the required route and query parameters for RemoveCertificateFromLoadBalancerListener.
func (cli *ZSClient) RemoveCertificateFromLoadBalancerListenerWithParams(listenerUuid string, deleteMode param.DeleteMode, params param.RemoveCertificateFromLoadBalancerListenerParamDetail) error {
	query := url.Values{
		"certificateUuid": {params.CertificateUuid},
	}
	return cli.deleteActionWithQuery("v1/load-balancers/listeners", listenerUuid, "certificate", deleteMode, query, "certificateUuid")
}

// GetLoadBalancerListenerACLEntries gets LoadBalancerListenerACLEntries by uuid
func (cli *ZSClient) GetLoadBalancerListenerACLEntries() (*view.StringView, error) {
	var resp view.StringView
	if err := cli.GetWithRespKey("v1/load-balancers/listeners/access-control-lists/entries", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddCertificateToLoadBalancerListener adds CertificateToLoadBalancerListener
func (cli *ZSClient) AddCertificateToLoadBalancerListener(listenerUuid string, params param.AddCertificateToLoadBalancerListenerParam) (*view.LoadBalancerListenerInventoryView, error) {
	resp := view.LoadBalancerListenerInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/load-balancers/listeners/%s/certificate", listenerUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddServerGroupToLoadBalancerListener adds ServerGroupToLoadBalancerListener
func (cli *ZSClient) AddServerGroupToLoadBalancerListener(listenerUuid string, params param.AddServerGroupToLoadBalancerListenerParam) (*view.LoadBalancerListenerInventoryView, error) {
	resp := view.LoadBalancerListenerInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/load-balancers/listeners/%s/servergroups", listenerUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCandidateVmNicsForLoadBalancer gets CandidateVmNicsForLoadBalancer by uuid
func (cli *ZSClient) GetCandidateVmNicsForLoadBalancer(uuid string) (*view.VmNicInventoryView, error) {
	var resp view.VmNicInventoryView
	if err := cli.GetWithRespKey("v1/load-balancers/listeners", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveVmNicFromLoadBalancer requires additional arguments.
//
// Deprecated: use RemoveVmNicFromLoadBalancerWithParams to supply the required fields.
func (cli *ZSClient) RemoveVmNicFromLoadBalancer(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("RemoveVmNicFromLoadBalancer requires additional arguments; use RemoveVmNicFromLoadBalancerWithParams")
}

// RemoveVmNicFromLoadBalancerWithParams sends the required route and query parameters for RemoveVmNicFromLoadBalancer.
func (cli *ZSClient) RemoveVmNicFromLoadBalancerWithParams(listenerUuid string, deleteMode param.DeleteMode, params param.RemoveVmNicFromLoadBalancerParamDetail) error {
	query := url.Values{
		"vmNicUuids": params.VmNicUuids,
	}
	return cli.deleteActionWithQuery("v1/load-balancers/listeners", listenerUuid, "vm-instances/nics", deleteMode, query, "vmNicUuids")
}

// AddVmNicToLoadBalancer adds VmNicToLoadBalancer
func (cli *ZSClient) AddVmNicToLoadBalancer(listenerUuid string, params param.AddVmNicToLoadBalancerParam) (*view.LoadBalancerListenerInventoryView, error) {
	resp := view.LoadBalancerListenerInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/load-balancers/listeners/%s/vm-instances/nics", listenerUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveAccessControlListFromLoadBalancer requires additional arguments.
//
// Deprecated: use RemoveAccessControlListFromLoadBalancerWithParams to supply the required fields.
func (cli *ZSClient) RemoveAccessControlListFromLoadBalancer(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("RemoveAccessControlListFromLoadBalancer requires additional arguments; use RemoveAccessControlListFromLoadBalancerWithParams")
}

// RemoveAccessControlListFromLoadBalancerWithParams sends the DELETE request with its operation parameters.
func (cli *ZSClient) RemoveAccessControlListFromLoadBalancerWithParams(listenerUuid string, deleteMode param.DeleteMode, params param.RemoveAccessControlListFromLoadBalancerParamDetail) error {
	query := url.Values{
		"aclUuids":         params.AclUuids,
		"serverGroupUuids": params.ServerGroupUuids,
	}
	return cli.deleteActionWithQuery("v1/load-balancers/listeners", listenerUuid, "access-control-lists", deleteMode, query, "aclUuids")
}
