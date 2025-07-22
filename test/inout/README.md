# Inout V5 测试用例

本目录包含了出入园相关接口的测试用例。

## 测试覆盖范围

### 出入园授权接口
- `TestAuthorizationList` - 查询授权列表
- `TestAuthorizationCreate` - 创建授权
- `TestAuthorizationUpdate` - 更新授权
- `TestAuthorizationDelete` - 删除授权

### 出入园统计组管理接口
- `TestCreate` - 创建出入园统计组
- `TestDelete` - 删除出入园统计组
- `TestGet` - 获取出入园统计组详情
- `TestGroupGet` - 获取出入园统计组详情
- `TestGroupList` - 获取出入园统计组列表
- `TestList` - 获取出入园统计组列表
- `TestSave` - 保存出入园统计组
- `TestUpdate` - 更新出入园统计组

### 出入园统计分析接口
- `TestDayCompareSummaryByHour` - 按小时对比日期出入园数据
- `TestHourSummaryByDevice` - 按设备小时统计出入园数据
- `TestSummaryByDate` - 按日期维度统计出入园数据
- `TestSummaryByTime` - 按时间维度统计出入园数据
- `TestTodaySummaryByGroupHour` - 按分组小时统计今日出入园数据

## 运行测试

### 运行所有测试
```bash
go test -v ./test/v5/inout/
```

### 运行单个测试
```bash
go test -v ./test/v5/inout/ -run TestSummaryByTime
```

## 配置说明

在运行测试前，请确保在 `inout_test.go` 中配置正确的测试参数：

- `accessId` - 访问ID
- `accessKey` - 访问密钥
- `token` - 认证令牌
- `sid` - 景区ID
- `gid` - 统计组ID列表
- `devices` - 设备ID列表

## 注意事项

1. 测试用例使用的是模拟数据，实际运行时请根据真实环境调整参数
2. 部分测试用例（如创建、更新、删除操作）可能会影响实际数据，请在测试环境中运行
3. 确保测试环境网络连接正常，能够访问 ODAS API 服务