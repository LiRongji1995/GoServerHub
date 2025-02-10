# GoServerHub
Go 服务器开发的 "Hub"，聚合各种服务器项目

[English version available here](./README_EN.md)

## 1. Web 服务 / API 服务器

适合入门 & 进阶开发。

### 目标

使用 Go 搭建 RESTful API 或 gRPC 服务，提供数据接口。

### 推荐框架

- **Gin**：高性能 Web 框架。
- **Fiber**：基于 Fasthttp 的轻量级框架。
- **Echo**：适用于小型 API 服务。
- **gRPC**：适用于高性能微服务通信。

### 项目案例

- 用户管理系统（注册 / 登录 / 权限管理）。
- 任务管理 API（待办事项、笔记、项目管理）。
- 天气 / 股票 / 新闻数据接口。
- 微服务架构 API。

### 进阶方向

- 结合数据库（MySQL / PostgreSQL / MongoDB）。
- 实现 JWT 认证。
- 限流（Rate Limiting）。

### 目录结构

```
GoProjects/
├── WebServices/
│   ├── UserManagement/
│   ├── TaskAPI/
│   ├── NewsAggregator/
│   ├── StockPriceAPI/
```

---

## 2. 爬虫 & 数据采集

适合练习网络编程与数据分析。

### 目标

使用 Colly / Goquery 爬取网页数据，并将数据存入数据库或 JSON 文件。

### 推荐框架

- **Colly**：轻量级爬虫框架。
- **Goquery**：类 jQuery 的 HTML 解析工具。
- **Rod / Chromedp**：支持 JavaScript 渲染的爬虫工具。

### 项目案例

- 新闻 & 论坛爬取（如知乎、掘金、豆瓣）。
- 电商数据爬取（商品价格监控、比价工具）。
- 招聘信息采集（如拉钩、BOSS 直聘）。
- 企业信息爬取（工商、商标、专利数据）。

### 进阶方向

- 分布式爬虫。
- 代理池与反爬策略。
- 结合 NLP 进行数据分析。

### 目录结构

```
GoProjects/
├── WebScraping/
│   ├── NewsScraper/
│   ├── ECommerceScraper/
│   ├── JobScraper/
│   ├── CompanyDataScraper/
```

---

## 3. 服务器 & 运维工具

适合熟悉 Linux 与网络开发的开发者。

### 目标

使用 Go 开发运维工具或服务器后台服务。

### 推荐框架

- **SSH / Net**：用于远程执行命令。
- **gopsutil**：系统监控工具。
- **Go SNMP**：网络设备监控工具。

### 项目案例

- 端口扫描器。
- 日志收集与监控工具。
- 远程任务执行工具（类似 Ansible）。
- 带 Web 界面的运维面板。

### 进阶方向

- 结合 Prometheus / Grafana 实现监控。
- 自动化 DevOps 工具开发。

### 目录结构

```
GoProjects/
├── DevOpsTools/
│   ├── PortScanner/
│   ├── LogMonitor/
│   ├── SSHAutomation/
│   ├── MetricsCollector/
```

---

## 4. 高并发 & 消息队列

适合提升 Go 并发编程能力。

### 目标

实现高并发任务处理、队列消费与分布式任务调度。

### 推荐框架

- **Kafka / RabbitMQ / NATS**：消息队列工具。
- **Redis Stream / Pub-Sub**：基于 Redis 的消息队列。
- **Go Worker Pools**：Go 并发任务池。

### 项目案例

- 短链接服务。
- 实时日志分析系统。
- 异步任务调度（任务队列）。
- WebSocket 聊天室。

### 进阶方向

- 事件驱动架构。
- 基于 WebSocket 的实时推送服务。

### 目录结构

```
GoProjects/
├── HighConcurrency/
│   ├── ShortURLService/
│   ├── RealTimeLogs/
│   ├── TaskQueue/
│   ├── ChatServer/
```

---

## 5. 代理 & 网络编程

适合熟悉 TCP/UDP/WebSocket 的开发者。

### 目标

使用 Go 开发代理、流量转发或加速器工具。

### 推荐框架

- **goproxy / mitmproxy**：代理工具。
- **fasthttp / net/http**：HTTP 库。
- **ssh / socks5**：代理协议支持。

### 项目案例

- HTTP 代理服务器。
- SOCKS5 代理。
- 负载均衡器。
- VPN / Shadowsocks 工具。

### 进阶方向

- 流量统计与日志记录。
- 安全与加密代理开发。

### 目录结构

```
GoProjects/
├── ProxyNetworking/
│   ├── HTTPProxy/
│   ├── SOCKS5Proxy/
│   ├── LoadBalancer/
│   ├── ShadowsocksGo/
```

---



