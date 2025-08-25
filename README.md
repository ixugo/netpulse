[🇺🇸 English](README.md) | [🇨🇳 中文](README_CN.md)

# NetPulse

🚀 Go library for intelligent IP information retrieval with fault tolerance and multi-provider support

## 🌟 Project Overview

NetPulse is a lightweight and efficient Go library designed for retrieving and parsing IP address information. It obtains information by requesting third-party APIs and provides two core modules:

- **ip module**: Get local external and internal IP addresses
- **geoip module**: Query geographic information based on IP addresses

## ⚠️ Important Usage Guidelines

### Proper Use of Free APIs
1. **Do not abuse**: These free APIs have usage limits, please comply with the service provider's terms
2. **Add caching**: Upper-level callers must implement caching mechanisms to avoid frequent repeated calls for the same data
3. **Enterprise recommendation**: For enterprise projects, please migrate to paid APIs as soon as possible to ensure service stability and reliability

### Built-in Caching Mechanism
- geoip module enables 1-hour memory cache by default
- Supports custom cache implementation
- Automatically handles cache hits and expiration

## 📦 Installation

```bash
go get github.com/ixugo/netpulse
```

## 🚀 Quick Start

### Getting Local IP Addresses

```go
package main

import (
    "fmt"
    "log"

    "github.com/ixugo/netpulse/ip"
)

func main() {
    // Get external IP address
    externalIP, err := ip.ExternalIP()
    if err != nil {
        log.Printf("Failed to get external IP: %v", err)
    } else {
        fmt.Printf("External IP: %s\n", externalIP)
    }

    // Get internal IP address
    internalIP := ip.InternalIP()
    fmt.Printf("Internal IP: %s\n", internalIP)
}
```

### Querying IP Geolocation Information

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/ixugo/netpulse/geoip"
)

func main() {
    // Create English geolocation query engine
    engine := geoip.New(geoip.English)

    // Query IP geolocation information
    ctx := context.Background()
    info, err := engine.Lookup(ctx, "8.8.8.8")
    if err != nil {
        log.Printf("Query failed: %v", err)
        return
    }

    fmt.Printf("IP: %s\n", info.IP)
    fmt.Printf("Country: %s\n", info.Country)
    fmt.Printf("Region: %s\n", info.Region)
    fmt.Printf("City: %s\n", info.City)
    fmt.Printf("ISP: %s\n", info.ISP)
}
```

#### geoip Module Details

**Supported Language Modes:**

```go
// English mode - uses international providers
engine := geoip.New(geoip.English)

// Chinese mode - uses Chinese provider for more accurate Chinese regional information
engine := geoip.New(geoip.Chinese)
```

**Intelligent Failover Mechanism:**
- 🔄 Automatic provider switching: Automatically tries the next provider when one times out or returns an error
- 🏆 Best choice: Returns the first successful result to ensure fastest response

### Custom Providers

You can specify which providers to use through the `WithHandlers` option:

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/ixugo/netpulse/geoip"
)

func main() {
    // Use only specified providers
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
        log.Printf("Query failed: %v", err)
        return
    }

    fmt.Printf("IP: %s\n", info.IP)
    fmt.Printf("Country: %s\n", info.Country)
    fmt.Printf("City: %s\n", info.City)
}
```

**Available Providers:**
- `NewFreeIPAPI()` - freeipapi.com
- `NewIfconfigco()` - ifconfig.co
- `NewIPapi()` - ipapi.com
- `NewIPwho()` - ipwho.io

## ⚙️ Advanced Usage

### Custom Configuration

```go
package main

import (
    "github.com/ixugo/netpulse/geoip"
)

// Disable cache completely
engine := geoip.New(
    geoip.English,
    geoip.WithCache(nil),
)
```

### Custom Cache Implementation

```go
package main

import (
    "github.com/ixugo/netpulse/geoip"
)

// Implement custom cache
type MyCache struct {
    // Your cache implementation
}

func (c *MyCache) Get(key string) (*geoip.Info, error) {
    // Implement cache retrieval logic
}

func (c *MyCache) Set(key string, info *geoip.Info) {
    // Implement cache storage logic
}

// Use custom cache
engine := geoip.New(
    geoip.English,
    geoip.WithCache(&MyCache{}),
)
```

## 📊 Data Structure

### geoip.Info Structure

```go
type Info struct {
    IP         string  // IP address
    Country    string  // Country
    Region     string  // Province/State
    RegionCode string  // Province/State code
    City       string  // City
    CityCode   string  // City code
    ISP        string  // Internet Service Provider
    Address    string  // Full address description
}
```

## 🙏 Acknowledgments

First and foremost, we would like to express our sincere gratitude to the following free API service providers. Their selfless dedication enables individual developers to learn and use these valuable services:

### English Providers
- **ipapi.com**
- **freeipapi.com**
- **ifconfig.co**
- **ipwho.io**

### Chinese Provider
- **whois.pconline.com.cn**

## 🤝 Contributing

Welcome to submit Issues and Pull Requests! Before contributing code, please ensure:

1. Code passes all tests
2. Add necessary test cases
3. Update relevant documentation
4. Follow the project's code style

## 📄 License

This project is open sourced under the MIT license - see [LICENSE](LICENSE) file for details

## 🔗 Related Links

- [GitHub Repository](https://github.com/ixugo/netpulse)
- [Go Module Proxy](https://pkg.go.dev/github.com/ixugo/netpulse)
- [goddd](https://github.com/ixugo/goddd)
- [nsqite](https://github.com/ixugo/nsqite)

---

**NetPulse** - Making IP information retrieval simple, reliable, and efficient! 🌍
