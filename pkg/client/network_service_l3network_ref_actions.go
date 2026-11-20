// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryNetworkServiceL3NetworkRef queries NetworkServiceL3NetworkRef list
func (cli *ZSClient) QueryNetworkServiceL3NetworkRef(params *param.QueryParam) ([]view.NetworkServiceL3NetworkRefInventoryView, error) {
	var resp []view.NetworkServiceL3NetworkRefInventoryView
	return resp, cli.List("v1/l3-networks/network-services/refs", params, &resp)
}

func (cli *ZSClient) GetNetworkServiceL3NetworkRef(uuid string) (*view.NetworkServiceL3NetworkRefInventoryView, error) {
	var resp view.NetworkServiceL3NetworkRefInventoryView
	if err := cli.Get("v1/l3-networks/network-services/refs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageNetworkServiceL3NetworkRef Pagination
func (cli *ZSClient) PageNetworkServiceL3NetworkRef(params *param.QueryParam) ([]view.NetworkServiceL3NetworkRefInventoryView, int, error) {
	var networkServiceL3NetworkRefs []view.NetworkServiceL3NetworkRefInventoryView
	total, err := cli.Page("v1/l3-networks/network-services/refs", params, &networkServiceL3NetworkRefs)
	return networkServiceL3NetworkRefs, total, err
}

// DetachNetworkServiceFromL3Network requires a service selector.
//
// Deprecated: use DetachNetworkServiceFromL3NetworkWithParams to specify the services.
func (cli *ZSClient) DetachNetworkServiceFromL3Network(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("DetachNetworkServiceFromL3Network requires a service selector; use DetachNetworkServiceFromL3NetworkWithParams")
}

// DetachNetworkServiceFromL3NetworkWithParams detaches a named service or the
// service lists specified for each provider. Cloud uses Service when supplied.
func (cli *ZSClient) DetachNetworkServiceFromL3NetworkWithParams(l3NetworkUuid string, deleteMode param.DeleteMode, params param.DetachNetworkServiceFromL3NetworkParamDetail) error {
	const resource = "v1/l3-networks"
	if params.Service != nil {
		service, err := actionPathSegment("service", *params.Service)
		if err != nil {
			return err
		}
		return cli.deleteActionWithQuery(resource, l3NetworkUuid, "network-services/"+service, deleteMode, url.Values{})
	}
	if len(params.NetworkServices) == 0 {
		return fmt.Errorf("service or networkServices is required")
	}

	query := url.Values{}
	required := make([]string, 0, len(params.NetworkServices))
	for provider, raw := range params.NetworkServices {
		if strings.TrimSpace(provider) == "" || strings.Contains(provider, ".") {
			return fmt.Errorf("networkServices provider keys must be nonempty and must not contain dots")
		}
		key := "networkServices." + provider
		switch values := raw.(type) {
		case []string:
			query[key] = values
		case []any:
			for _, value := range values {
				service, ok := value.(string)
				if !ok {
					return fmt.Errorf("%s must contain only service names", key)
				}
				query.Add(key, service)
			}
		default:
			return fmt.Errorf("%s must be a list of service names", key)
		}
		required = append(required, key)
	}
	return cli.deleteActionWithQuery(resource, l3NetworkUuid, "network-services", deleteMode, query, required...)
}

// AttachNetworkServiceToL3Network operates on NetworkServiceToL3Network
func (cli *ZSClient) AttachNetworkServiceToL3Network(l3NetworkUuid string, params param.AttachNetworkServiceToL3NetworkParam) (*view.L3NetworkInventoryView, error) {
	resp := view.L3NetworkInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/l3-networks/%s/network-services", l3NetworkUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
