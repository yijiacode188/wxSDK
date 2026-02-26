package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type RequestParams struct {
	Url         string          `json:"url"`
	Params      url.Values      `json:"params"`
	Header      http.Header     `json:"header"`
	ContentType string          `json:"contentType"`
	Body        []byte          `json:"body"`
	Timeout     time.Duration   `json:"timeout"`
	Context     context.Context `json:"-"`
}

// Response HTTP 响应
type Response struct {
	StatusCode int
	Header     http.Header
	Body       []byte
	Raw        *http.Response
}

func HttpGet[T any](params *RequestParams) (T, *Response, error) {

	var result T
	httpResp, err := doRequest("GET", params)
	if err != nil {
		return result, nil, err
	}

	if strings.Contains("application/json", httpResp.Header.Get("Content-Type")) {
		// 解析响应体
		if err := json.Unmarshal(httpResp.Body, &result); err != nil {
			return result, httpResp, err
		}
	}
	return result, httpResp, nil
}

// HttpPost 发送 POST 请求并解析响应到指定类型
func HttpPost[T any](params *RequestParams) (T, *Response, error) {

	var result T

	httpResp, err := doRequest("POST", params)
	if err != nil {
		return result, nil, err
	}
	if strings.Contains("application/json", httpResp.Header.Get("Content-Type")) {
		// 解析响应体
		if err := json.Unmarshal(httpResp.Body, &result); err != nil {
			return result, httpResp, err
		}
	}
	return result, httpResp, nil
}

// doRequest 执行 HTTP 请求
func doRequest(method string, params *RequestParams) (*Response, error) {
	if params == nil {
		return nil, errors.New("params is nil")
	}

	// 准备请求
	req, err := prepareRequest(method, params)
	if err != nil {
		return nil, err
	}

	return executeRequest(req, params)
}

// prepareRequest 准备 HTTP 请求
func prepareRequest(method string, params *RequestParams) (*http.Request, error) {
	var req *http.Request
	var err error

	// 处理 URL 参数
	urlStr := params.Url
	if params.Params != nil {
		if strings.Contains(urlStr, "?") {
			urlStr += "&" + params.Params.Encode()
		} else {
			urlStr += "?" + params.Params.Encode()
		}
	}

	// 创建请求
	if method == "POST" || method == "PUT" || method == "PATCH" {
		var bodyReader io.Reader

		// 如果有 Body，使用 Body
		if len(params.Body) > 0 {
			bodyReader = bytes.NewReader(params.Body)
		} else if params.Params != nil {
			// 否则使用 Params
			contentType := params.ContentType
			if contentType == "" {
				contentType = "application/x-www-form-urlencoded"
			}

			if contentType == "application/json" {
				// JSON 格式
				jsonData, err := json.Marshal(params.Params)
				if err != nil {
					return nil, err
				}
				bodyReader = bytes.NewReader(jsonData)
			} else {
				// Form 格式
				bodyReader = strings.NewReader(params.Params.Encode())
			}
		}

		req, err = http.NewRequest(method, urlStr, bodyReader)
	} else {
		req, err = http.NewRequest(method, urlStr, nil)
	}

	if err != nil {
		return nil, errors.New("create request failed")
	}

	// 设置请求头
	if params.Header != nil {
		req.Header = params.Header
	}

	// 设置 Content-Type
	if params.ContentType != "" {
		req.Header.Set("Content-Type", params.ContentType)
	} else if method == "POST" && req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	return req, nil
}

// DefaultTimeout 默认超时时间
const DefaultTimeout = 30 * time.Second

// executeRequest 执行 HTTP 请求
func executeRequest(req *http.Request, params *RequestParams) (*Response, error) {

	// 设置超时
	timeout := params.Timeout
	if timeout == 0 {
		timeout = DefaultTimeout
	}

	// 创建 HTTP 客户端
	client := &http.Client{
		Timeout: timeout,
		// 可以在这里添加其他配置，如 Transport、CheckRedirect 等
	}

	// 使用上下文
	ctx := params.Context
	if ctx == nil {
		ctx = context.Background()
	}
	req = req.WithContext(ctx)

	// 执行请求
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// 读取响应体
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// 检查 HTTP 状态码
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return &Response{
			StatusCode: resp.StatusCode,
			Header:     resp.Header,
			Body:       body,
			Raw:        resp,
		}, errors.New(string(body))
	}

	return &Response{
		StatusCode: resp.StatusCode,
		Header:     resp.Header,
		Body:       body,
		Raw:        resp,
	}, nil
}
