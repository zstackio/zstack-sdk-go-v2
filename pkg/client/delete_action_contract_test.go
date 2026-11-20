// Copyright (c) ZStack.io, Inc.

package client

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
)

// Cloud 5.5.48 maps DELETE collections from repeated query keys. Check the
// requests received on the wire, including scalar selectors and optional fields.
func TestDeleteActionContracts(t *testing.T) {
	optional := "option"
	port := 8443
	tests := []deleteActionContract{
		{
			name:  "RemoveResourcesFromDirectory",
			path:  "/zstack/v1/remove/resources/directory",
			query: url.Values{"directoryUuid": {"selector"}, "resourceUuids": {"item-a", "item-b"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveResourcesFromDirectoryWithParams(param.DeleteModePermissive, param.RemoveResourcesFromDirectoryParamDetail{DirectoryUuid: "selector", ResourceUuids: []string{"item-a", "item-b"}})
			},
		},
		{
			name:  "DetachPoliciesFromUser",
			path:  "/zstack/v1/accounts/users/parent-id/policies",
			query: url.Values{"policyUuids": {"item-a", "item-b"}},
			call: func(cli *ZSClient) error {
				return cli.DetachPoliciesFromUserWithParams("parent-id", param.DeleteModePermissive, param.DetachPoliciesFromUserParamDetail{PolicyUuids: []string{"item-a", "item-b"}})
			},
		},
		{
			name:  "RemoveIAM2VirtualIDsFromGroup",
			path:  "/zstack/v1/iam2/projects/groups/parent-id/virtual-ids",
			query: url.Values{"virtualIDUuids": {"item-a", "item-b"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveIAM2VirtualIDsFromGroupWithParams("parent-id", param.DeleteModePermissive, param.RemoveIAM2VirtualIDsFromGroupParamDetail{VirtualIDUuids: []string{"item-a", "item-b"}})
			},
		},
		{
			name:  "RemoveIAM2VirtualIDsFromOrganization",
			path:  "/zstack/v1/iam2/organizations/parent-id/virtual-ids",
			query: url.Values{"virtualIDUuids": {"item-a", "item-b"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveIAM2VirtualIDsFromOrganizationWithParams("parent-id", param.DeleteModePermissive, param.RemoveIAM2VirtualIDsFromOrganizationParamDetail{VirtualIDUuids: []string{"item-a", "item-b"}})
			},
		},
		{
			name:  "RemoveIAM2VirtualIDsFromProject",
			path:  "/zstack/v1/iam2/projects/parent-id/virtual-ids",
			query: url.Values{"virtualIDUuids": {"item-a", "item-b"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveIAM2VirtualIDsFromProjectWithParams("parent-id", param.DeleteModePermissive, param.RemoveIAM2VirtualIDsFromProjectParamDetail{VirtualIDUuids: []string{"item-a", "item-b"}})
			},
		},
		{
			name:  "RemoveIAM2VirtualIDsFromProjects",
			path:  "/zstack/v1/iam2/projects/virtual-ids",
			query: url.Values{"projectUuids": {"item-a", "item-b"}, "virtualIDUuids": {"item-a", "item-b"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveIAM2VirtualIDsFromProjectsWithParams(param.DeleteModePermissive, param.RemoveIAM2VirtualIDsFromProjectsParamDetail{ProjectUuids: []string{"item-a", "item-b"}, VirtualIDUuids: []string{"item-a", "item-b"}})
			},
		},
		{
			name:  "RemoveRolesFromIAM2VirtualIDGroup",
			path:  "/zstack/v1/iam2/projects/groups/parent-id/roles",
			query: url.Values{"roleUuids": {"item-a", "item-b"}, "projectUuid": {"option"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveRolesFromIAM2VirtualIDGroupWithParams("parent-id", param.DeleteModePermissive, param.RemoveRolesFromIAM2VirtualIDGroupParamDetail{RoleUuids: []string{"item-a", "item-b"}, ProjectUuid: &optional})
			},
		},
		{
			name:  "RemoveRolesFromIAM2VirtualID",
			path:  "/zstack/v1/iam2/projects/virtual-ids/parent-id/roles",
			query: url.Values{"roleUuids": {"item-a", "item-b"}, "projectUuid": {"option"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveRolesFromIAM2VirtualIDWithParams("parent-id", param.DeleteModePermissive, param.RemoveRolesFromIAM2VirtualIDParamDetail{RoleUuids: []string{"item-a", "item-b"}, ProjectUuid: &optional})
			},
		},
		{
			name:  "DetachL3NetworksFromIPsecConnection",
			path:  "/zstack/v1/ipsec/parent-id/l3networks",
			query: url.Values{"l3NetworkUuids": {"item-a", "item-b"}},
			call: func(cli *ZSClient) error {
				return cli.DetachL3NetworksFromIPsecConnectionWithParams("parent-id", param.DeleteModePermissive, param.DetachL3NetworksFromIPsecConnectionParamDetail{L3NetworkUuids: []string{"item-a", "item-b"}})
			},
		},
		{
			name:  "RemoveRemoteCidrsFromIPsecConnection",
			path:  "/zstack/v1/ipsec/parent-id/remote-cidrs",
			query: url.Values{"peerCidrs": {"192.0.2.0/24", "2001:db8::/64"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveRemoteCidrsFromIPsecConnectionWithParams("parent-id", param.DeleteModePermissive, param.RemoveRemoteCidrsFromIPsecConnectionParamDetail{PeerCidrs: []string{"192.0.2.0/24", "2001:db8::/64"}})
			},
		},
		{
			name:  "DeleteVmNicFromSecurityGroup",
			path:  "/zstack/v1/security-groups/parent-id/vm-instances/nics",
			query: url.Values{"vmNicUuids": {"item-a", "item-b"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteVmNicFromSecurityGroupWithParams("parent-id", param.DeleteModePermissive, param.DeleteVmNicFromSecurityGroupParamDetail{VmNicUuids: []string{"item-a", "item-b"}})
			},
		},
		{
			name:  "RemoveVmNicFromLoadBalancer",
			path:  "/zstack/v1/load-balancers/listeners/parent-id/vm-instances/nics",
			query: url.Values{"vmNicUuids": {"item-a", "item-b"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveVmNicFromLoadBalancerWithParams("parent-id", param.DeleteModePermissive, param.RemoveVmNicFromLoadBalancerParamDetail{VmNicUuids: []string{"item-a", "item-b"}})
			},
		},
		{
			name:  "RemoveMonFromCephBackupStorage",
			path:  "/zstack/v1/backup-storage/ceph/parent-id/mons",
			query: url.Values{"monHostnames": {"item-a", "item-b"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveMonFromCephBackupStorageWithParams("parent-id", param.DeleteModePermissive, param.RemoveMonFromCephBackupStorageParamDetail{MonHostnames: []string{"item-a", "item-b"}})
			},
		},
		{
			name:  "RemoveMonFromCephPrimaryStorage",
			path:  "/zstack/v1/primary-storage/ceph/parent-id/mons",
			query: url.Values{"monHostnames": {"item-a", "item-b"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveMonFromCephPrimaryStorageWithParams("parent-id", param.DeleteModePermissive, param.RemoveMonFromCephPrimaryStorageParamDetail{MonHostnames: []string{"item-a", "item-b"}})
			},
		},
		{
			name:  "DeleteSecurityGroupRule",
			path:  "/zstack/v1/security-groups/rules",
			query: url.Values{"ruleUuids": {"item-a", "item-b"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteSecurityGroupRuleWithParams(param.DeleteModePermissive, param.DeleteSecurityGroupRuleParamDetail{RuleUuids: []string{"item-a", "item-b"}})
			},
		},
		{
			name:  "DeleteIpAddress",
			path:  "/zstack/v1/l3-networks/parent-id/ip-address",
			query: url.Values{"usedIpUuids": {"item-a", "item-b"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteIpAddressWithParams("parent-id", param.DeleteModePermissive, param.DeleteIpAddressParamDetail{UsedIpUuids: []string{"item-a", "item-b"}})
			},
		},
		{
			name:  "DeleteNicQos",
			path:  "/zstack/v1/vm-instances/parent-id/nic-qos",
			query: url.Values{"direction": {"out"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteNicQosWithParams("parent-id", param.DeleteModePermissive, param.DeleteNicQosParamDetail{Direction: "out"})
			},
		},
		{
			name:  "DeleteVmStaticIp",
			path:  "/zstack/v1/vm-instances/parent-id/static-ips",
			query: url.Values{"l3NetworkUuid": {"selector"}, "staticIp": {"option"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteVmStaticIpWithParams("parent-id", param.DeleteModePermissive, param.DeleteVmStaticIpParamDetail{L3NetworkUuid: "selector", StaticIp: &optional})
			},
		},
		{
			name:  "DetachIsoFromVmInstance",
			path:  "/zstack/v1/vm-instances/parent-id/iso",
			query: url.Values{"isoUuid": {"option"}},
			call: func(cli *ZSClient) error {
				return cli.DetachIsoFromVmInstanceWithParams("parent-id", param.DeleteModePermissive, param.DetachIsoFromVmInstanceParamDetail{IsoUuid: &optional})
			},
		},
		{
			name:  "DeleteVolumeQos",
			path:  "/zstack/v1/volumes/parent-id/qos",
			query: url.Values{"mode": {"option"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteVolumeQosWithParams("parent-id", param.DeleteModePermissive, param.DeleteVolumeQosParamDetail{Mode: &optional})
			},
		},
		{
			name:  "DeleteVipQos",
			path:  "/zstack/v1/vips/parent-id/vip-qos",
			query: url.Values{"port": {"8443"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteVipQosWithParams("parent-id", param.DeleteModePermissive, param.DeleteVipQosParamDetail{Port: &port})
			},
		},
		{
			name:  "RemovePolicyStatementsFromRole",
			path:  "/zstack/v1/identities/roles/parent-id/policy-statements",
			query: url.Values{"policyStatementUuids": {"item-a", "item-b"}},
			call: func(cli *ZSClient) error {
				return cli.RemovePolicyStatementsFromRoleWithParams("parent-id", param.DeleteModePermissive, param.RemovePolicyStatementsFromRoleParamDetail{PolicyStatementUuids: []string{"item-a", "item-b"}})
			},
		},
		{
			name:  "RemoveHostRouteFromL3Network",
			path:  "/zstack/v1/l3-networks/parent-id/hostroute",
			query: url.Values{"prefix": {"192.0.2.0/24"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveHostRouteFromL3NetworkWithParams("parent-id", param.DeleteModePermissive, param.RemoveHostRouteFromL3NetworkParamDetail{Prefix: "192.0.2.0/24"})
			},
		},
		{
			name:  "RemoveCertificateFromLoadBalancerListener",
			path:  "/zstack/v1/load-balancers/listeners/parent-id/certificate",
			query: url.Values{"certificateUuid": {"selector"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveCertificateFromLoadBalancerListenerWithParams("parent-id", param.DeleteModePermissive, param.RemoveCertificateFromLoadBalancerListenerParamDetail{CertificateUuid: "selector"})
			},
		},
		{
			name:  "RemoveServerGroupFromLoadBalancerListener",
			path:  "/zstack/v1/load-balancers/listeners/parent-id/servergroups",
			query: url.Values{"serverGroupUuid": {"selector"}},
			call: func(cli *ZSClient) error {
				return cli.RemoveServerGroupFromLoadBalancerListenerWithParams("parent-id", param.DeleteModePermissive, param.RemoveServerGroupFromLoadBalancerListenerParamDetail{ServerGroupUuid: "selector"})
			},
		},
		{
			name:  "single security group rule",
			path:  "/zstack/v1/security-groups/rules",
			query: url.Values{"ruleUuids": {"rule-id"}},
			call: func(cli *ZSClient) error {
				return cli.DeleteSecurityGroupRule("rule-id", param.DeleteModePermissive)
			},
		},
		{
			name:  "legacy ISO detach",
			path:  "/zstack/v1/vm-instances/parent-id/iso",
			query: url.Values{},
			call: func(cli *ZSClient) error {
				return cli.DetachIsoFromVmInstance("parent-id", param.DeleteModePermissive)
			},
		},
		{
			name:  "legacy volume QoS",
			path:  "/zstack/v1/volumes/parent-id/qos",
			query: url.Values{},
			call: func(cli *ZSClient) error {
				return cli.DeleteVolumeQos("parent-id", param.DeleteModePermissive)
			},
		},
		{
			name:  "legacy VIP QoS",
			path:  "/zstack/v1/vips/parent-id/vip-qos",
			query: url.Values{},
			call: func(cli *ZSClient) error {
				return cli.DeleteVipQos("parent-id", param.DeleteModePermissive)
			},
		},
	}
	for _, tc := range tests {
		tc.query.Set("deleteMode", "Permissive")
	}
	assertDeleteActionContracts(t, tests)
}

type deleteActionContract struct {
	name  string
	path  string
	query url.Values
	call  func(*ZSClient) error
}

func assertDeleteActionContracts(t *testing.T, tests []deleteActionContract) {
	t.Helper()
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			received := false
			cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
				received = true
				if r.Method != http.MethodDelete || r.URL.Path != tc.path {
					t.Errorf("got %s %s, want DELETE %s", r.Method, r.URL.Path, tc.path)
				}
				if got := r.URL.Query(); !reflect.DeepEqual(got, tc.query) {
					t.Errorf("got query %#v, want %#v", got, tc.query)
				}
				w.WriteHeader(http.StatusNoContent)
			})
			if err := tc.call(cli); err != nil {
				t.Fatal(err)
			}
			if !received {
				t.Fatal("expected a DELETE request")
			}
		})
	}
}

func TestDeleteActionsRejectIncompleteRequests(t *testing.T) {
	tests := map[string]func(*ZSClient) error{
		"legacy NIC detach": func(cli *ZSClient) error {
			return cli.DeleteVmNicFromSecurityGroup("security-group", "")
		},
		"legacy load balancer NIC detach": func(cli *ZSClient) error {
			return cli.RemoveVmNicFromLoadBalancer("listener", "")
		},
		"legacy NIC QoS delete": func(cli *ZSClient) error {
			return cli.DeleteNicQos("vm", "")
		},
		"missing parent UUID": func(cli *ZSClient) error {
			return cli.DetachPoliciesFromUserWithParams("", "", param.DetachPoliciesFromUserParamDetail{PolicyUuids: []string{"policy"}})
		},
		"missing directory UUID": func(cli *ZSClient) error {
			return cli.RemoveResourcesFromDirectoryWithParams("", param.RemoveResourcesFromDirectoryParamDetail{ResourceUuids: []string{"resource"}})
		},
		"missing required list": func(cli *ZSClient) error {
			return cli.RemoveIAM2VirtualIDsFromProjectsWithParams("", param.RemoveIAM2VirtualIDsFromProjectsParamDetail{
				ProjectUuids: []string{"project"},
			})
		},
		"blank list entry": func(cli *ZSClient) error {
			return cli.DeleteSecurityGroupRuleWithParams("", param.DeleteSecurityGroupRuleParamDetail{RuleUuids: []string{"rule", " "}})
		},
		"unfiltered billing delete": func(cli *ZSClient) error {
			return cli.DeleteBillingWithParams("", param.DeleteBillingParamDetail{})
		},
		"license without selector": func(cli *ZSClient) error {
			return cli.DeleteLicenseWithParams("node", "", param.DeleteLicenseParamDetail{})
		},
		"invalid network service map": func(cli *ZSClient) error {
			return cli.DetachNetworkServiceFromL3NetworkWithParams("network", "", param.DetachNetworkServiceFromL3NetworkParamDetail{
				NetworkServices: map[string]any{"provider": []any{"DNS", false}},
			})
		},
		"missing VTEP cluster": func(cli *ZSClient) error {
			return cli.DeleteVxlanPoolRemoteVtepWithParams("pool", "", "", param.DeleteVxlanPoolRemoteVtepParamDetail{RemoteVtepIp: "192.0.2.10"})
		},
		"missing path UUID": func(cli *ZSClient) error {
			return cli.DeleteImageWithParams("", "", param.DeleteImageParamDetail{})
		},
		"path separator in UUID": func(cli *ZSClient) error {
			return cli.DeleteImageWithParams("../volumes/id", "", param.DeleteImageParamDetail{})
		},
		"missing IAM2 projects": func(cli *ZSClient) error {
			return cli.RemoveIAM2VirtualIDGroupFromProjectsWithParams("", param.RemoveIAM2VirtualIDGroupFromProjectsParamDetail{
				GroupUuids: []string{"group"},
			})
		},
		"empty VIP selection": func(cli *ZSClient) error {
			return cli.DetachVipFromVpcSharedQosWithParams("qos", "", param.DetachVipFromVpcSharedQosParamDetail{})
		},
	}
	for name, call := range tests {
		t.Run(name, func(t *testing.T) {
			cli := newTestZSClient(t, func(w http.ResponseWriter, r *http.Request) {
				t.Errorf("invalid arguments must not send a request: %s", r.URL)
				w.WriteHeader(http.StatusNoContent)
			})
			if err := call(cli); err == nil {
				t.Fatal("expected validation error")
			}
		})
	}
}
