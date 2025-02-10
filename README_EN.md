# GoServerHub
A "Hub" for Go server development, aggregating various server projects.

[Chinese version available here](./README.md)

## 1. Web Services / API Servers

Suitable for beginners & advanced development.

### Objectives

Build RESTful APIs or gRPC services using Go to provide data interfaces.

### Recommended Frameworks

- **Gin**: High-performance web framework.
- **Fiber**: Lightweight framework based on Fasthttp.
- **Echo**: Suitable for small API services.
- **gRPC**: High-performance microservice communication.

### Project Examples

- User management system (registration / login / permission management).
- Task management API (to-do lists, notes, project management).
- Weather / Stock / News data APIs.
- Microservice architecture APIs.

### Advanced Topics

- Integrating databases (MySQL / PostgreSQL / MongoDB).
- Implementing JWT authentication.
- Rate limiting.

### Directory Structure

```
GoProjects/
├── WebServices/
│   ├── UserManagement/
│   ├── TaskAPI/
│   ├── NewsAggregator/
│   ├── StockPriceAPI/
```

---

## 2. Web Scraping & Data Collection

Great for practicing web programming and data analysis.

### Objectives

Use Colly / Goquery to scrape web data and store it in a database or JSON files.

### Recommended Frameworks

- **Colly**: Lightweight web scraping framework.
- **Goquery**: jQuery-like HTML parsing tool.
- **Rod / Chromedp**: Tools for scraping JavaScript-rendered content.

### Project Examples

- News & forum scraping (e.g., Zhihu, Juejin, Douban).
- E-commerce data scraping (price tracking, comparison tools).
- Job listing scraping (e.g., LinkedIn, Indeed).
- Business data scraping (company registry, trademarks, patents).

### Advanced Topics

- Distributed web scrapers.
- Proxy pools & anti-scraping strategies.
- NLP-based data analysis.

### Directory Structure

```
GoProjects/
├── WebScraping/
│   ├── NewsScraper/
│   ├── ECommerceScraper/
│   ├── JobScraper/
│   ├── CompanyDataScraper/
```

---

## 3. Servers & DevOps Tools

Great for developers familiar with Linux & networking.

### Objectives

Develop DevOps tools or backend server services using Go.

### Recommended Frameworks

- **SSH / Net**: For remote command execution.
- **gopsutil**: System monitoring tools.
- **Go SNMP**: Network monitoring tools.

### Project Examples

- Port scanners.
- Log collection & monitoring tools.
- Remote task execution tools (similar to Ansible).
- Web-based DevOps dashboards.

### Advanced Topics

- Integrating Prometheus / Grafana for monitoring.
- Automating DevOps processes.

### Directory Structure

```
GoProjects/
├── DevOpsTools/
│   ├── PortScanner/
│   ├── LogMonitor/
│   ├── SSHAutomation/
│   ├── MetricsCollector/
```

---

## 4. High Concurrency & Message Queues

Great for enhancing Go concurrency programming skills.

### Objectives

Implement high-concurrency task processing, queue consumption, and distributed task scheduling.

### Recommended Frameworks

- **Kafka / RabbitMQ / NATS**: Message queue tools.
- **Redis Stream / Pub-Sub**: Redis-based messaging queues.
- **Go Worker Pools**: Go concurrency task pools.

### Project Examples

- URL shortening service.
- Real-time log analysis system.
- Asynchronous task scheduling (task queues).
- WebSocket chat server.

### Advanced Topics

- Event-driven architecture.
- Real-time push services using WebSocket.

### Directory Structure

```
GoProjects/
├── HighConcurrency/
│   ├── ShortURLService/
│   ├── RealTimeLogs/
│   ├── TaskQueue/
│   ├── ChatServer/
```

---

## 5. Proxies & Network Programming

Suitable for developers familiar with TCP/UDP/WebSocket.

### Objectives

Develop proxies, traffic forwarding, or accelerator tools using Go.

### Recommended Frameworks

- **goproxy / mitmproxy**: Proxy tools.
- **fasthttp / net/http**: HTTP libraries.
- **ssh / socks5**: Proxy protocol support.

### Project Examples

- HTTP proxy server.
- SOCKS5 proxy.
- Load balancer.
- VPN / Shadowsocks tool.

### Advanced Topics

- Traffic statistics & logging.
- Secure & encrypted proxy development.

### Directory Structure

```
GoProjects/
├── ProxyNetworking/
│   ├── HTTPProxy/
│   ├── SOCKS5Proxy/
│   ├── LoadBalancer/
│   ├── ShadowsocksGo/
```

---

