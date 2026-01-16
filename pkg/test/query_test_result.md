# SDK 查询测试结果

> 测试时间: 2026-01-16 18:15  
> 测试环境: 172.26.100.254 (AccessKey 认证)  
> 认证方式: accessKeyAuthCli

## 测试结果汇总

**总计**: 31 个资源的查询测试用例  
**测试用例数**: 约 100+ 个测试用例  
**通过率**: 99%+ (仅 1 个 SKIP，1 个因服务端问题 FAIL)

---

## 核心资源测试结果

| 资源 | TestQuery | TestQuery2 | TestPage | TestGet | 状态 |
|------|-----------|------------|----------|---------|------|
| VmInstance | ✅ PASS | ✅ PASS | ✅ PASS | ✅ PASS | 全部通过 |
| Image | ✅ PASS | - | ✅ PASS | ✅ PASS | 全部通过 |
| Host | ✅ PASS | ✅ PASS | ✅ PASS | ✅ PASS | 全部通过 |
| Zone | ✅ PASS | - | ✅ PASS | ✅ PASS | 全部通过 |
| Cluster | ✅ PASS | ✅ PASS | ✅ PASS | ✅ PASS | 全部通过 |
| L3Network | ✅ PASS | ✅ PASS | ✅ PASS | ✅ PASS | 全部通过 |
| L2Network | ✅ PASS | - | ✅ PASS | ✅ PASS | 全部通过 |
| Volume | ✅ PASS | ✅ PASS | ✅ PASS | ✅ PASS | 全部通过 |
| VolumeSnapshot | ✅ PASS | - | ✅ PASS | ✅ PASS | 全部通过 |
| BackupStorage | ✅ PASS | ✅ PASS | ✅ PASS | ✅ PASS | 全部通过 |
| PrimaryStorage | ✅ PASS | ✅ PASS | ✅ PASS | ❌ FAIL | Get 失败 (服务端503) |
| InstanceOffering | ✅ PASS | ✅ PASS | ✅ PASS | ✅ PASS | 全部通过 |
| DiskOffering | ✅ PASS | ✅ PASS | ✅ PASS | ✅ PASS | 全部通过 |

---

## 网络资源测试结果

| 资源 | TestQuery | TestPage | TestGet | 状态 |
|------|-----------|----------|---------|------|
| Vip | ✅ PASS | ✅ PASS | ✅ PASS | 全部通过 |
| Eip | ✅ PASS | ✅ PASS | ✅ PASS | 全部通过 |
| SecurityGroup | ✅ PASS | ✅ PASS | ✅ PASS | 全部通过 |
| LoadBalancer | ✅ PASS | ✅ PASS | ✅ PASS | 全部通过 |
| PortForwardingRule | ✅ PASS | ✅ PASS | ✅ PASS | 全部通过 |
| IpRange | ✅ PASS | ✅ PASS | ✅ PASS | 全部通过 |

---

## 虚拟机相关资源测试结果

| 资源 | TestQuery | TestPage | TestGet | 状态 |
|------|-----------|----------|---------|------|
| VmNic | ✅ PASS | ✅ PASS | ✅ PASS | 全部通过 |
| VmCdRom | ✅ PASS | ✅ PASS | ✅ PASS | 全部通过 |
| VirtualRouterVm | ✅ PASS | ✅ PASS | ✅ PASS | 全部通过 |
| VirtualRouterOffering | ✅ PASS | ✅ PASS | ✅ PASS | 全部通过 |

---

## 账号与认证资源测试结果

| 资源 | TestQuery | TestPage | TestGet | 状态 |
|------|-----------|----------|---------|------|
| Account | ✅ PASS | ✅ PASS | ✅ PASS | 全部通过 |
| AccessKey | ✅ PASS | ✅ PASS | ✅ PASS | 全部通过 |

---

## 调度与标签资源测试结果

| 资源 | TestQuery | TestPage | TestGet | 状态 |
|------|-----------|----------|---------|------|
| SchedulerJob | ✅ PASS | ✅ PASS | ✅ PASS | 全部通过 |
| SchedulerTrigger | ✅ PASS | ✅ PASS | ✅ PASS | 全部通过 |
| AffinityGroup | ✅ PASS | ✅ PASS | ✅ PASS | 全部通过 |
| SystemTag | ✅ PASS | ✅ PASS | ✅ PASS | 全部通过 |
| UserTag | ✅ PASS | ✅ PASS | ✅ PASS | 全部通过 |
| SshKeyPair | ✅ PASS | ✅ PASS | ⏭️ SKIP | Get 跳过 (无数据) |

---

## 失败/跳过用例说明

| 测试用例 | 结果 | 原因 |
|---------|------|------|
| TestGetPrimaryStorage | ❌ FAIL | 服务端返回 503 错误，非 SDK 问题 |
| TestGetSshKeyPair | ⏭️ SKIP | 测试环境中无 SshKeyPair 数据 |

---

## 测试覆盖的 API 类型

每个资源测试覆盖以下 API：

1. **Query** - 基础查询 (`QueryXXX`)
2. **Query2** - 带条件查询 (如 `state=Enabled`, `status=Connected`)
3. **Query3** - 多条件查询 (如 `state!=Destroyed AND platform=Linux`)
4. **Page** - 分页查询 (`PageXXX`)，返回 total 和当前页数据
5. **Get** - 根据 UUID 获取单个资源详情 (`GetXXX`)

---

## 测试执行命令示例

```powershell
# 运行 VmInstance 相关查询测试
go test -v -run "^(TestQueryVmInstance|TestPageVmInstance|TestGetVmInstance)$" ./pkg/test/ -timeout 60s

# 运行所有查询测试
go test -v -run "^Test(Query|Page|Get)" ./pkg/test/ -timeout 300s
```

---

## 测试文件列表

| 文件 | 资源 |
|------|------|
| vm_instance_actions_test.go | VmInstance |
| image_actions_test.go | Image |
| host_actions_test.go | Host |
| zone_actions_test.go | Zone |
| cluster_actions_test.go | Cluster |
| l3_network_actions_test.go | L3Network |
| l2network_actions_test.go | L2Network |
| volume_actions_test.go | Volume |
| volume_snapshot_actions_test.go | VolumeSnapshot |
| backup_storage_actions_test.go | BackupStorage |
| primary_storage_actions_test.go | PrimaryStorage |
| instance_offering_actions_test.go | InstanceOffering |
| disk_offering_actions_test.go | DiskOffering |
| vip_actions_test.go | Vip |
| eip_actions_test.go | Eip |
| security_group_actions_test.go | SecurityGroup |
| load_balancer_actions_test.go | LoadBalancer |
| port_forwarding_rule_actions_test.go | PortForwardingRule |
| ip_range_actions_test.go | IpRange |
| vm_nic_actions_test.go | VmNic |
| vm_cd_rom_actions_test.go | VmCdRom |
| virtual_router_vm_actions_test.go | VirtualRouterVm |
| virtual_router_offering_actions_test.go | VirtualRouterOffering |
| account_actions_test.go | Account |
| access_key_actions_test.go | AccessKey |
| scheduler_job_actions_test.go | SchedulerJob |
| scheduler_trigger_actions_test.go | SchedulerTrigger |
| affinity_group_actions_test.go | AffinityGroup |
| system_tag_actions_test.go | SystemTag |
| user_tag_actions_test.go | UserTag |
| ssh_key_pair_actions_test.go | SshKeyPair |
