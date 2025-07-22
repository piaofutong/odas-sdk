# Basic Analyze 测试用例

本目录包含了对 `dev/odas-sdk/odas/request/v5/basic_analyze` 包装接口的测试用例。

## 测试接口列表

### 订单相关接口

1. **OrderStatistics** - 订单统计
   - 测试函数: `TestOrderStatistics`
   - 功能: 获取指定时间范围内的订单统计数据，支持同比和环比分析

2. **OrderSummaryByYMD** - 按年月日订单汇总
   - 测试函数: `TestOrderSummaryByYMD`
   - 功能: 按年月日维度汇总订单数据

3. **OrderSummaryByProduct** - 按产品订单汇总
   - 测试函数: `TestOrderSummaryByProduct`
   - 功能: 按产品维度汇总订单数据

4. **OrderSummaryByTicket** - 按票订单汇总
   - 测试函数: `TestOrderSummaryByTicket`
   - 功能: 按票维度汇总订单数据

5. **BookingOrderList** - 预订订单列表
   - 测试函数: `TestBookingOrderList`
   - 功能: 获取预订订单列表数据

### 汇总分析接口

6. **ReportSummary** - 报表汇总
   - 测试函数: `TestReportSummary`
   - 功能: 获取报表汇总数据

7. **AnnualCardSummary** - 年卡汇总
   - 测试函数: `TestAnnualCardSummary`
   - 功能: 获取年卡相关的汇总数据

8. **DistributorSummary** - 分销商汇总
   - 测试函数: `TestDistributorSummary`
   - 功能: 获取分销商相关的汇总数据

9. **TerminalSummary** - 终端汇总
   - 测试函数: `TestTerminalSummary`
   - 功能: 获取终端相关的汇总数据

10. **Summary** - 通用汇总
    - 测试函数: `TestSummary`
    - 功能: 获取通用汇总数据，支持同比环比

11. **Rank** - 排名统计
    - 测试函数: `TestRank`
    - 功能: 获取排名统计数据

12. **SummaryByHour** - 按小时汇总
    - 测试函数: `TestSummaryByHour`
    - 功能: 按小时维度汇总数据

13. **SummaryByTicket** - 按票汇总
    - 测试函数: `TestSummaryByTicket`
    - 功能: 按票维度汇总数据

14. **SummaryByLevel1** - 按一级分类汇总
    - 测试函数: `TestSummaryByLevel1`
    - 功能: 按一级分类维度汇总数据

15. **SummaryByLevel2** - 按二级分类汇总
    - 测试函数: `TestSummaryByLevel2`
    - 功能: 按二级分类维度汇总数据

16. **SummaryByLevel1AndLevel1Name** - 按一级分类和名称汇总
    - 测试函数: `TestSummaryByLevel1AndLevel1Name`
    - 功能: 按一级分类和名称维度汇总数据

### TOI相关接口

17. **ToiSummary** - TOI汇总
    - 测试函数: `TestToiSummary`
    - 功能: 获取TOI汇总数据

18. **ToiStatistics** - TOI统计
    - 测试函数: `TestToiStatistics`
    - 功能: 获取TOI统计数据

### 销售相关接口

19. **SalesDetail** - 销售明细
    - 测试函数: `TestSalesDetail`
    - 功能: 获取销售明细数据

20. **RefundSummaryByLevel2** - 按二级分类退款汇总
    - 测试函数: `TestRefundSummaryByLevel2`
    - 功能: 按二级分类维度汇总退款数据

21. **SummaryByLevel2AndTicket** - 按二级分类和票汇总
    - 测试函数: `TestSummaryByLevel2AndTicket`
    - 功能: 按二级分类和票维度汇总数据

22. **SummaryByTicketAndChannel** - 按票和渠道汇总
    - 测试函数: `TestSummaryByTicketAndChannel`
    - 功能: 按票和渠道维度汇总数据

23. **SummaryByTicketAndDay** - 按票和日期汇总
    - 测试函数: `TestSummaryByTicketAndDay`
    - 功能: 按票和日期维度汇总数据

### 票务相关接口

24. **TicketList** - 票列表
    - 测试函数: `TestTicketList`
    - 功能: 获取票列表数据

25. **TodayTicketingDetail** - 今日票务明细
    - 测试函数: `TestTodayTicketingDetail`
    - 功能: 获取今日票务明细数据

## 运行测试

### 运行所有测试
```bash
cd /mnt/source_code/dev/odas-sdk
go test ./test/v5/basic_analyze/... -v
```

### 运行单个测试
```bash
cd /mnt/source_code/dev/odas-sdk
go test ./test/v5/basic_analyze -run TestOrderStatistics -v
```

## 测试配置

测试用例中使用的配置变量：

- `accessId`: 访问ID
- `accessKey`: 访问密钥
- `token`: 访问令牌
- `sid`: 景区ID
- `lid`: 位置ID列表
- `orderType`: 订单类型
- `ticketId`: 票ID

## 注意事项

1. 测试前请确保配置了正确的访问凭证
2. 测试使用的是本地模式，通过 `odas.SetLocalMode()` 设置
3. 所有测试都使用了过去7天的时间范围作为测试数据
4. 测试结果会通过 `t.Logf()` 输出到控制台
5. 如果接口返回错误，测试会通过 `t.Fatal()` 终止

## 文件结构

```
basic_analyze/
├── README.md              # 本说明文件
└── basic_analyze_test.go  # 统一测试文件，包含所有basic_analyze模块的测试用例
```