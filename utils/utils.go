package utils

import (
	"fmt"
	"github.com/docker/docker/client"
	"golang.org/x/net/proxy"
	"net"
	"net/http"
	"net/url"
	"time"
)

func CreateDockerClient(address, clientversion, proxyaddr string) (*client.Client, error) {
	opts := []client.Opt{
		client.WithHost(address),
		client.WithVersion(clientversion),
		client.WithAPIVersionNegotiation(),
	}

	if proxyaddr != "" {
		proxyURL, err := url.Parse(proxyaddr)
		if err != nil {
			return nil, fmt.Errorf("Failed to parse the proxy address: %w", err)
		}

		dialer, err := proxy.FromURL(proxyURL, &net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		})
		if err != nil {
			return nil, fmt.Errorf("Failed to create the proxy dialer: %w", err)
		}

		contextDialer := dialer.(proxy.ContextDialer)

		httpTransport := &http.Transport{
			DialContext:         contextDialer.DialContext,
			TLSHandshakeTimeout: 10 * time.Second,
			IdleConnTimeout:     90 * time.Second,
		}

		httpClient := &http.Client{
			Transport: httpTransport,
			Timeout:   0,
		}

		opts = append(opts, client.WithHTTPClient(httpClient))
	}

	clt, err := client.NewClientWithOpts(opts...)
	if err != nil {
		return nil, err
	}

	return clt, nil
}
