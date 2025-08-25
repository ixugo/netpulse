[🇺🇸 English](README.md) | [🇨🇳 中文](README_CN.md)

# NetPulse

🚀 Go IP 信息获取库，具备智能故障转移和多服务商支持

高德地图用户可以看这个库 [github.com/ixugo/amap](https://github.com/ixugo/amap)

## 项目简介

NetPulse 是一个轻量级、高效的 Go 语言库，专为获取和解析 IP 地址信息而设计。通过请求第三方 API 服务来获取这些信息，项目提供了两个核心模块：

- **ip 模块**：获取本机内外网 IP 地址
- **geoip 模块**：根据 IP 地址查询地理位置信息


## ⚠️ 重要使用建议

### 对于免费 API 的正确使用
1. **请勿滥用**：这些免费 API 都是有使用限制的，请遵守服务商的使用条款
2. **添加缓存**：上层调用者务必添加缓存机制，避免对相同数据频繁重复调用
3. **企业项目建议**：如果用于企业项目，请尽快迁移到付费 API，以确保服务稳定性和可靠性

### 本项目已实现的缓存机制
- geoip 模块默认开启 1 小时的内存缓存
- 支持自定义缓存实现
- 自动处理缓存命中和过期逻辑

## 安装

```bash
go get github.com/ixugo/netpulse
```

## 🚀 快速开始

### 获取本机 IP 地址

```go
package main

import (
    "fmt"
    "log"

    "github.com/ixugo/netpulse/ip"
)

func main() {
    // 获取外网 IP 地址
    externalIP, err := ip.ExternalIP()
    if err != nil {
        log.Printf("获取外网IP失败: %v", err)
    } else {
        fmt.Printf("外网IP: %s\n", externalIP)
    }

    // 获取内网 IP 地址
    internalIP := ip.InternalIP()
    fmt.Printf("内网IP: %s\n", internalIP)
}

### 查询 IP 地理位置信息

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/ixugo/netpulse/geoip"
)

func main() {
    // 创建中文地理位置查询引擎
    engine := geoip.New(geoip.Chinese)

    // 查询 IP 地理位置信息
    ctx := context.Background()
    info, err := engine.Lookup(ctx, "8.8.8.8")
    if err != nil {
        log.Printf("查询失败: %v", err)
        return
    }

    fmt.Printf("IP: %s\n", info.IP)
    fmt.Printf("地区: %s\n", info.Region)
    fmt.Printf("城市: %s\n", info.City)
    fmt.Printf("地址: %s\n", info.Address)
}
```

#### geoip 模块功能详解

**支持的语言模式：**

```go
// 英文模式 - 使用国际服务商
engine := geoip.New(geoip.English)

// 中文模式 - 使用中文服务商，提供更准确的中文地区信息
engine := geoip.New(geoip.Chinese)
```

**智能故障转移机制：**
- 自动切换服务商：当一个服务商超时或返回错误时，自动尝试下一个
- 最佳选择：返回第一个成功的结果，确保最快响应

### 自定义服务商

你可以通过 `WithHandlers` 选项来指定使用哪些服务商：

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/ixugo/netpulse/geoip"
)

func main() {
    // 只使用指定的服务商
    engine := geoip.New(
        geoip.English,
        geoip.WithHandlers(
            geoip.NewFreeIPAPI(),  // freeipapi.com
            geoip.NewIfconfigco(), // ifconfig.co
        ),
    )

    ctx := context.Background()
    info, err := engine.Lookup(ctx, "8.8.8.8")
    if err != nil {
        log.Printf("查询失败: %v", err)
        return
    }

    fmt.Printf("IP: %s\n", info.IP)
    fmt.Printf("国家: %s\n", info.Country)
    fmt.Printf("城市: %s\n", info.City)
}
```

**可用的服务商：**
- `NewFreeIPAPI()` - freeipapi.com
- `NewIfconfigco()` - ifconfig.co
- `NewIPapi()` - ipapi.com
- `NewIPwho()` - ipwho.io

## 高级用法

### 自定义配置

```go
package main

import (
    "github.com/ixugo/netpulse/geoip"
)

// 完全禁用缓存
engine := geoip.New(
    geoip.English,
    geoip.WithCache(nil),
)
```

### 自定义缓存实现

```go
package main

import (
    "github.com/ixugo/netpulse/geoip"
)

// 实现自定义缓存
type MyCache struct {
    // 你的缓存实现
}

func (c *MyCache) Get(key string) (*geoip.Info, error) {
    // 实现获取缓存逻辑
}

func (c *MyCache) Set(key string, info *geoip.Info) {
    // 实现设置缓存逻辑
}

// 使用自定义缓存
engine := geoip.New(
    geoip.English,
    geoip.WithCache(&MyCache{}),
)
```

## 📊 返回数据结构

### geoip.Info 结构体

```go
type Info struct {
    IP         string  // IP 地址
    Country    string  // 国家
    Region     string  // 省份/州
    RegionCode string  // 省份/州代码
    City       string  // 城市
    CityCode   string  // 城市代码
    ISP        string  // 互联网服务提供商
    Address    string  // 完整地址描述
}
```

## 🙏 致谢

首先，我们要衷心感谢以下免费 API 服务提供商，正是因为他们的无私奉献，才让个人开发者能够学习和使用这些宝贵的服务：

### 英文服务商
- **ipapi.com**
- **freeipapi.com**
- **ifconfig.co**
- **ipwho.io**

### 中文服务商
- **whois.pconline.com.cn**

## 贡献指南

欢迎提交 Issue 和 Pull Request！在贡献代码前，请确保：

1. 代码通过所有测试
2. 添加必要的测试用例
3. 更新相关文档
4. 遵循项目的代码风格

## 开源协议

本项目采用 MIT 协议开源 - 详见 [LICENSE](LICENSE) 文件

## 相关链接

- [goddd 模板](https://github.com/ixugo/goddd)
- [nsqite 事务消息队列](https://github.com/ixugo/nsqite)
- [amap 高德地图 API](https://github.com/ixugo/amap)

---

**NetPulse** - 让 IP 信息获取变得简单、可靠、高效！ 🌍
