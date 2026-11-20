# ZSTAC-91245 DELETE 接口复核

以 SDK 主分支 `a79c1dd6` 和 Cloud 5.5.48 的 API 快照为基线，复查原来的 81 个候选。SDK 当前保留 67 项修复，第二批包含 26 项路径偏差、2 项必填参数缺失和 13 项可选参数未暴露。

## 首批 26 项

- 标签使用此前的三参数 `DetachTagFromResources(tagUuid, deleteMode, resourceUuids)`，与 Terraform 的 ZSTAC-91245 分支一致。
- `DeleteSecurityGroupRule` 将单个 UUID 发送为 `ruleUuids`，批量使用 `DeleteSecurityGroupRuleWithParams`。
- `DetachIsoFromVmInstance`、`DeleteVolumeQos` 和 `DeleteVipQos` 的旧入口调用正确的子资源路径。有额外筛选条件时可使用对应的 `WithParams`。
- 其余已修接口使用对应的 `WithParams` 方法传入必填字段。旧方法保留签名，在请求前返回迁移错误。

  - `RemoveResourcesFromDirectory` → `RemoveResourcesFromDirectoryWithParams`
  - `DetachPoliciesFromUser` → `DetachPoliciesFromUserWithParams`
  - `RemoveIAM2VirtualIDsFromGroup` → `RemoveIAM2VirtualIDsFromGroupWithParams`
  - `RemoveIAM2VirtualIDsFromOrganization` → `RemoveIAM2VirtualIDsFromOrganizationWithParams`
  - `RemoveIAM2VirtualIDsFromProject` → `RemoveIAM2VirtualIDsFromProjectWithParams`
  - `RemoveIAM2VirtualIDsFromProjects` → `RemoveIAM2VirtualIDsFromProjectsWithParams`
  - `RemoveRolesFromIAM2VirtualIDGroup` → `RemoveRolesFromIAM2VirtualIDGroupWithParams`
  - `RemoveRolesFromIAM2VirtualID` → `RemoveRolesFromIAM2VirtualIDWithParams`
  - `DetachL3NetworksFromIPsecConnection` → `DetachL3NetworksFromIPsecConnectionWithParams`
  - `RemoveRemoteCidrsFromIPsecConnection` → `RemoveRemoteCidrsFromIPsecConnectionWithParams`
  - `DeleteVmNicFromSecurityGroup` → `DeleteVmNicFromSecurityGroupWithParams`
  - `RemoveVmNicFromLoadBalancer` → `RemoveVmNicFromLoadBalancerWithParams`
  - `RemoveMonFromCephBackupStorage` → `RemoveMonFromCephBackupStorageWithParams`
  - `RemoveMonFromCephPrimaryStorage` → `RemoveMonFromCephPrimaryStorageWithParams`
  - `DeleteIpAddress` → `DeleteIpAddressWithParams`
  - `DeleteNicQos` → `DeleteNicQosWithParams`
  - `DeleteVmStaticIp` → `DeleteVmStaticIpWithParams`
  - `RemovePolicyStatementsFromRole` → `RemovePolicyStatementsFromRoleWithParams`
  - `RemoveHostRouteFromL3Network` → `RemoveHostRouteFromL3NetworkWithParams`
  - `RemoveCertificateFromLoadBalancerListener` → `RemoveCertificateFromLoadBalancerListenerWithParams`
  - `RemoveServerGroupFromLoadBalancerListener` → `RemoveServerGroupFromLoadBalancerListenerWithParams`

HTTP 契约测试覆盖实际发送的路径、重复查询参数、兼容入口和输入校验。Cloud API 快照用于核实契约。

## 第二批 41 项

同一路径的旧方法继续使用默认参数。只有需要调用方明确提供目标时，旧方法才会在发请求前返回迁移错误。集合接口已有明确的单个 ID 时，旧方法将它作为列表的唯一元素发送。

- `DeleteBillingWithParams` 需要账号或时间筛选条件。整个账单表的清除使用显式的 `DeleteAllBilling`。
- `DeleteLicenseWithParams` 需要管理节点及许可证 UUID 或模块。Terraform 销毁许可证时传入 state 中的许可证 UUID。
- `DetachNetworkServiceFromL3NetworkWithParams` 支持按服务名称解绑或按提供者传入多个服务。
- `GetVipQosByPort` 从 Cloud 的 QoS 列表定位规则，Terraform 刷新和删除使用同一端口。

### 已修复的路径偏差（26 项）

- `DeleteAlert`
- `DeleteBilling`
- `DeleteLicense`
- `DeleteLogConfiguration`
- `DeleteMetricData`
- `DeleteResourceStackVmPortMonitor`
- `DetachHostFromHostSchedulingRuleGroup`
- `DetachNetworkServiceFromL3Network`
- `DetachServiceFromObservabilityServer`
- `DetachUserDefinedXmlHookScriptFromVm`
- `DetachVipFromVpcSharedQos`
- `DetachVmFromVmSchedulingRuleGroup`
- `RemoveAccessControlListFromLoadBalancer`
- `RemoveAttributesFromIAM2Organization`
- `RemoveAttributesFromIAM2Project`
- `RemoveAttributesFromIAM2VirtualID`
- `RemoveAttributesFromIAM2VirtualIDGroup`
- `RemoveDnsFromVpcRouter`
- `RemoveIAM2VirtualIDGroupFromProjects`
- `RemoveRendezvousPointFromMulticastRouter`
- `RemoveSchedulerJobsFromSchedulerJobGroup`
- `RemoveTicketTypesFromTicketFlowCollection`
- `RemoveVRouterNetworksFromFlowMeter`
- `RemoveVRouterNetworksFromOspfArea`
- `UnregisterLicenseServer`
- `WithdrawLicenseCapacityApplication`

### 已补齐的必填参数（2 项）

- `DeleteVxlanPoolRemoteVtep`
- `DetachProvisionNicFromBonding`

### 已暴露的可选参数（13 项）

- `DeleteBuildAppExportHistory`
- `DeleteCbtTask`
- `DeleteCdpTask`
- `DeleteDatabaseBackup`
- `DeleteImage`
- `DeleteResourcePrice`
- `DeleteVmBackup`
- `DeleteVolumeBackup`
- `DeleteVolumeSnapshot`
- `DeleteVolumeSnapshotGroup`
- `LogOut`
- `RemoveSNSSmsReceiver`
- `SdnControllerRemoveHost`
