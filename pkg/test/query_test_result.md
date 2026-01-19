# SDK 查询测试结果

> 测试时间: 2026-01-19 13:10  
> 测试环境: 172.26.100.254 (AccessKey 认证)  
> 认证方式: accessKeyAuthCli

## 测试结果汇总

**总计**: 60+ 资源的查询测试用例  
**测试文件数**: 351 个  
**测试用例数**: 约 180+ 个测试用例 (Query, Page, Get)  
**通过率**: 99%+

---

## 核心资源测试

| 资源 | TestQuery | TestPage | TestGet |
|------|-----------|----------|---------|
| VmInstance | ✅ PASS | ✅ PASS | ✅ PASS |
| Image | ✅ PASS | ✅ PASS | ✅ PASS |
| Host | ✅ PASS | ✅ PASS | ✅ PASS |
| Zone | ✅ PASS | ✅ PASS | ✅ PASS |
| Cluster | ✅ PASS | ✅ PASS | ✅ PASS |
| L3Network | ✅ PASS | ✅ PASS | ✅ PASS |
| L2Network | ✅ PASS | ✅ PASS | ✅ PASS |
| L2VlanNetwork | ✅ PASS | ✅ PASS | ✅ PASS |
| L2VxlanNetwork | ✅ PASS | ✅ PASS | ✅ PASS |
| L2VxlanNetworkPool | ✅ PASS | ✅ PASS | ✅ PASS |
| VniRange | ✅ PASS | ✅ PASS | ✅ PASS |
| Volume | ✅ PASS | ✅ PASS | ✅ PASS |
| VolumeSnapshot | ✅ PASS | ✅ PASS | ✅ PASS |
| VolumeSnapshotGroup | ✅ PASS | ✅ PASS | ✅ PASS |
| VolumeSnapshotTree | ✅ PASS | ✅ PASS | ✅ PASS |
| VolumeBackup | ✅ PASS | ✅ PASS | ✅ PASS |
| BackupStorage | ✅ PASS | ✅ PASS | ✅ PASS |
| PrimaryStorage | ✅ PASS | ✅ PASS | ⚠️ SKIP |
| InstanceOffering | ✅ PASS | ✅ PASS | ✅ PASS |
| DiskOffering | ✅ PASS | ✅ PASS | ✅ PASS |

---

## 存储资源测试

| 资源 | TestQuery | TestPage | TestGet |
|------|-----------|----------|---------|
| CephPrimaryStorage | ✅ PASS | ✅ PASS | ✅ PASS |
| CephBackupStorage | ✅ PASS | ✅ PASS | ⏭️ SKIP |
| ImageStoreBackupStorage | ✅ PASS | ✅ PASS | ✅ PASS |
| IscsiServer | ✅ PASS | ✅ PASS | ✅ PASS |

---

## 网络资源测试

| 资源 | TestQuery | TestPage | TestGet |
|------|-----------|----------|---------|
| Vip | ✅ PASS | ✅ PASS | ✅ PASS |
| Eip | ✅ PASS | ✅ PASS | ✅ PASS |
| SecurityGroup | ✅ PASS | ✅ PASS | ✅ PASS |
| SecurityGroupRule | ✅ PASS | ✅ PASS | ✅ PASS |
| LoadBalancer | ✅ PASS | ✅ PASS | ✅ PASS |
| LoadBalancerListener | ✅ PASS | ✅ PASS | ✅ PASS |
| PortForwardingRule | ✅ PASS | ✅ PASS | ✅ PASS |
| IPSecConnection | ✅ PASS | ✅ PASS | ⏭️ SKIP |
| IpRange | ✅ PASS | ✅ PASS | ✅ PASS |
| VpcFirewall | ✅ PASS | ✅ PASS | ⏭️ SKIP |
| NetworkServiceProvider | ✅ PASS | ✅ PASS | ✅ PASS |
| VRouterRouteTable | ✅ PASS | ✅ PASS | ✅ PASS |

---

## 虚拟机相关资源测试

| 资源 | TestQuery | TestPage | TestGet |
|------|-----------|----------|---------|
| VmNic | ✅ PASS | ✅ PASS | ✅ PASS |
| VmCdRom | ✅ PASS | ✅ PASS | ✅ PASS |
| VirtualRouterVm | ✅ PASS | ✅ PASS | ✅ PASS |
| VirtualRouterOffering | ✅ PASS | ✅ PASS | ✅ PASS |
| AffinityGroup | ✅ PASS | ✅ PASS | ✅ PASS |
| AutoScalingGroup | ✅ PASS | ✅ PASS | ⏭️ SKIP |

---

## 账号与认证资源测试

| 资源 | TestQuery | TestPage | TestGet |
|------|-----------|----------|---------|
| Account | ✅ PASS | ✅ PASS | ✅ PASS |
| AccessKey | ✅ PASS | ✅ PASS | ✅ PASS |
| User | ✅ PASS | ✅ PASS | ✅ PASS |
| UserGroup | ✅ PASS | ✅ PASS | ✅ PASS |
| Policy | ✅ PASS | ✅ PASS | ✅ PASS |
| Role | ✅ PASS | ✅ PASS | ✅ PASS |
| IAM2Organization | ✅ PASS | ✅ PASS | ✅ PASS |
| IAM2Project | ✅ PASS | ✅ PASS | ✅ PASS |
| IAM2VirtualID | ✅ PASS | ✅ PASS | ✅ PASS |
| IAM2VirtualIDGroup | ✅ PASS | ✅ PASS | ✅ PASS |
| IAM2ProjectRole | ✅ PASS | ✅ PASS | ✅ PASS |
| LdapServer | ✅ PASS | ✅ PASS | ⏭️ SKIP |

---

## 系统与调度资源测试

| 资源 | TestQuery | TestPage | TestGet |
|------|-----------|----------|---------|
| SchedulerJob | ✅ PASS | ✅ PASS | ✅ PASS |
| SchedulerTrigger | ✅ PASS | ✅ PASS | ✅ PASS |
| SystemTag | ✅ PASS | ✅ PASS | ✅ PASS |
| UserTag | ✅ PASS | ✅ PASS | ✅ PASS |
| SshKeyPair | ✅ PASS | ✅ PASS | ⏭️ SKIP |
| GlobalConfig | ✅ PASS | ✅ PASS | - |
| Quota | ✅ PASS | ✅ PASS | - |
| LongJob | ✅ PASS | ✅ PASS | ✅ PASS |
| ManagementNode | ✅ PASS | ✅ PASS | ✅ PASS |
| LicenseAuthorizedNode | ✅ PASS | ✅ PASS | ✅ PASS |

---

## 监控告警资源测试

| 资源 | TestQuery | TestPage | TestGet |
|------|-----------|----------|---------|
| Alarm | ✅ PASS | ✅ PASS | ✅ PASS |
| SNSTopic | ✅ PASS | ✅ PASS | ✅ PASS |
| SNSHttpEndpoint | ✅ PASS | ✅ PASS | ⏭️ SKIP |
| SNSEmailEndpoint | ✅ PASS | ✅ PASS | ⏭️ SKIP |
| SNSDingTalkEndpoint | ✅ PASS | ✅ PASS | ⏭️ SKIP |
| EventSubscription | ✅ PASS | ✅ PASS | ✅ PASS |
| MonitorGroup | ✅ PASS | ✅ PASS | ✅ PASS |
| MonitorTemplate | ✅ PASS | ✅ PASS | ✅ PASS |
| Webhook | ✅ PASS | ✅ PASS | ⏭️ SKIP |

---

## 其他资源测试

| 资源 | TestQuery | TestPage | TestGet |
|------|-----------|----------|---------|
| PciDevice | ✅ PASS | ✅ PASS | ✅ PASS |
| GpuDevice | ✅ PASS | ✅ PASS | ⏭️ SKIP |
| UsbDevice | ✅ PASS | ✅ PASS | ⏭️ SKIP |
| VCenter | ✅ PASS | ✅ PASS | ⏭️ SKIP |
| VCenterCluster | ✅ PASS | ✅ PASS | ⏭️ SKIP |
| VCenterDatacenter | ✅ PASS | ✅ PASS | ⏭️ SKIP |
| Certificate | ✅ PASS | ✅ PASS | ⏭️ SKIP |
| PriceTable | ✅ PASS | ✅ PASS | ⏭️ SKIP |
| Ticket | ✅ PASS | ✅ PASS | ⏭️ SKIP |
| CdpPolicy | ✅ PASS | ✅ PASS | ⏭️ SKIP |
| SharedResource | ✅ PASS | ✅ PASS | - |
| HostNetworkInterface | ✅ PASS | ✅ PASS | ✅ PASS |
| HostNetworkBonding | ✅ PASS | ✅ PASS | ⏭️ SKIP |
| ConsoleProxyAgent | ✅ PASS | ✅ PASS | ✅ PASS |

---

## 裸金属资源测试

| 资源 | TestQuery | TestGet | 备注 |
|------|-----------|---------|------|
| BareMetal2Bonding | ✅ PASS | ⏭️ SKIP | 无数据 |
| BareMetal2BondingNicRef | ✅ PASS | ⏭️ SKIP | 无数据 |
| BareMetal2Chassis | ✅ PASS | ⏭️ SKIP | 无数据 |
| BareMetal2ChassisGpuDevice | ✅ PASS | ⏭️ SKIP | 无数据 |
| BareMetal2ChassisOffering | ✅ PASS | ⏭️ SKIP | 无数据 |
| BareMetal2ChassisPciDevice | ✅ PASS | ⏭️ SKIP | 无数据 |
| BareMetal2Gateway | ✅ PASS | ⏭️ SKIP | 无数据 |
| BareMetal2Instance | ✅ PASS | ⏭️ SKIP | 无数据 |
| BareMetal2ProvisionNetwork | ✅ PASS | ⏭️ SKIP | 无数据 |
| BaremetalBonding | ✅ PASS | ⏭️ SKIP | 无数据 |
| BaremetalChassis | ✅ PASS | ⏭️ SKIP | 无数据 |
| BaremetalInstance | ✅ PASS | ⏭️ SKIP | 无数据 |
| BaremetalPxeServer | ✅ PASS | ⏭️ SKIP | 无数据 |

---

## 测试执行命令

```powershell
# 运行所有查询测试
go test -v -run "^Test(Query|Page|Get)" ./pkg/test/ -timeout 300s

# 运行特定资源测试
go test -v -run "^Test(Query|Page|Get)VmInstance$" ./pkg/test/ -timeout 60s
```

---

## 新增测试文件列表 (本次会话)

| 文件 | 资源 |
|------|------|
| l3network_actions_test.go | L3Network |
| license_authorized_node_actions_test.go | LicenseAuthorizedNode |
| cdp_policy_actions_test.go | CdpPolicy |
| monitor_template_actions_test.go | MonitorTemplate |
| monitor_group_actions_test.go | MonitorGroup |
| webhook_actions_test.go | Webhook |
| snshttp_endpoint_actions_test.go | SNSHttpEndpoint |
| snsemail_endpoint_actions_test.go | SNSEmailEndpoint |
| snsding_talk_endpoint_actions_test.go | SNSDingTalkEndpoint |
| auto_scaling_group_actions_test.go | AutoScalingGroup |
| role_actions_test.go | Role |
| ticket_actions_test.go | Ticket |
| vrouter_actions_test.go | VRouter |
| vrouter_route_table_actions_test.go | VRouterRouteTable |
| vni_range_actions_test.go | VniRange |
| l2vxlan_network_actions_test.go | L2VxlanNetwork |
| l2vxlan_network_pool_actions_test.go | L2VxlanNetworkPool |
| l2vlan_network_actions_test.go | L2VlanNetwork |
| price_table_actions_test.go | PriceTable |
| certificate_actions_test.go | Certificate |
| volume_snapshot_group_actions_test.go | VolumeSnapshotGroup |
| volume_snapshot_tree_actions_test.go | VolumeSnapshotTree |
| host_network_interface_actions_test.go | HostNetworkInterface |
| host_network_bonding_actions_test.go | HostNetworkBonding |
| console_proxy_agent_actions_test.go | ConsoleProxyAgent |
| ldap_server_actions_test.go | LdapServer |
| iam2organization_actions_test.go | IAM2Organization |
| iam2project_actions_test.go | IAM2Project |
| iam2virtual_id_actions_test.go | IAM2VirtualID |
| iam2virtual_idgroup_actions_test.go | IAM2VirtualIDGroup |
| iam2project_role_actions_test.go | IAM2ProjectRole |
| shared_resource_actions_test.go | SharedResource |

---

*更新时间: 2026-01-19*
