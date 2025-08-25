package ip

import (
	"context"
	"io"
	"net/http"
	"time"
)

var services = []string{
	"https://checkip.amazonaws.com/",
	"http://myexternalip.com/raw",
	"https://ifconfig.me/ip",
	"https://ipinfo.io/ip",
	"https://icanhazip.com",
	"https://api.ipify.org",
	"https://ifconfig.co/ip",

	// ipv6?
	// https://api64.ipify.org
}

var defaultServices = func() *StringScorer {
	ss := NewStringScorer(8)
	for _, s := range services {
		ss.Set(s)
	}
	return ss
}()

var DefaultClient = http.Client{
	Timeout: 5 * time.Second,
}

func ExternalIP() (string, error) {
	return multipleRequests()
}

type value struct {
	ip  string
	err error
}

func multipleRequests() (string, error) {
	// 创建可取消的context用于取消其他请求
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
	defer cancel()

	ch := make(chan *value, 1) // buffered channel，只接收第一个结果

	// 启动所有请求，每个请求都使用带超时的context
	var i int
	for link := range defaultServices.All() {
		i++
		select {
		case <-ctx.Done():
			goto end
		default:
			go func(link string) {
				// 使用带超时的context创建请求
				req, err := http.NewRequestWithContext(ctx, "GET", link, nil)
				if err != nil {
					return
				}

				resp, err := DefaultClient.Do(req)
				if err != nil {
					return
				}
				defer resp.Body.Close()

				if resp.StatusCode != 200 {
					return
				}

				body, err := io.ReadAll(io.LimitReader(resp.Body, 128))

				// 尝试发送结果，如果channel已满或其他请求先发送了，select会处理
				v := value{ip: string(body), err: err}
				select {
				case ch <- &v:
					// 发送成功，取消其他所有请求
					cancel()
					if err == nil && v.ip != "" {
						defaultServices.AddScore(link)
					}
				case <-ctx.Done():
					// 已经被取消或超时
					return
				}
			}(link)
		}
	}
end:

	// 等待第一个结果或超时
	select {
	case result := <-ch:
		return result.ip, result.err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}
