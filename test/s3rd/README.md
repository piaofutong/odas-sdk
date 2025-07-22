# S3rd V5 测试用例

本目录包含了第三方服务相关接口的测试用例。

## 测试覆盖范围

### 停车场相关接口
- `TestInoutSummaryByHour` - 按小时统计停车场出入数据
- `TestLocationInSummary` - 位置入园汇总统计
- `TestSpace` - 停车位信息查询

### 酒店相关接口
- `TestOccupancy` - 酒店入住率统计
- `TestRevenueSummaryByCodeCategory` - 按代码分类收入汇总
- `TestRmOrderSummary` - 客房订单汇总
- `TestRmSaleSummary` - 客房销售汇总
- `TestRmSaleSummaryByBind` - 按绑定关系客房销售汇总

## 运行测试

### 运行所有测试
```bash
go test -v ./test/v5/s3rd/
```

### 运行单个测试
```bash
go test -v ./test/v5/s3rd/ -run TestOccupancy
```

## 配置说明

在运行测试前，请确保在 `s3rd_test.go` 中配置正确的测试参数：

- `accessId` - 访问ID
- `accessKey` - 访问密钥
- `token` - 认证令牌
- `sid` - 景区ID

## 注意事项

1. 测试用例使用的是模拟数据，实际运行时请根据真实环境调整参数
2. 确保测试环境网络连接正常，能够访问 ODAS API 服务
3. 第三方服务接口可能需要额外的权限配置