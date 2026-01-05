// Copyright (c) ZStack.io, Inc.

package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"time"
)

// AuthType authentication type
type AuthType string

const (
	AuthTypeAccessKey AuthType = "accesskey"
	AuthTypeLogin    AuthType = "login"
)

const (
	defaultZStackPort        = 8080
)

// ZSConfig client configuration
type ZSConfig struct {
	hostname        string
	port            int
	contextPath     string
	accessKeyId     string
	accessKeySecret string
	username        string
	password        string
	authType        AuthType
	debug           bool
	timeout         time.Duration
}

// NewZSConfig creates a new configuration
func NewZSConfig(hostname string, port int, contextPath string) *ZSConfig {
	return &ZSConfig{
		hostname:    hostname,
		port:        port,
		contextPath: contextPath,
		timeout:     30 * time.Second,
	}
}

// DefaultZSConfig creates a default configuration
func DefaultZSConfig(hostname, contextPath string) *ZSConfig {
	return NewZSConfig(hostname, defaultZStackPort, contextPath)
}

// AccessKey sets access key authentication
func (config *ZSConfig) AccessKey(id, secret string) *ZSConfig {
	config.accessKeyId = id
	config.accessKeySecret = secret
	config.authType = AuthTypeAccessKey
	return config
}

// Login sets login authentication
func (config *ZSConfig) Login(username, password string) *ZSConfig {
	config.username = username
	config.password = password
	config.authType = AuthTypeLogin
	return config
}

// Debug enables debug mode
func (config *ZSConfig) Debug(debug bool) *ZSConfig {
	config.debug = debug
	return config
}

// ZSClient ZStack API client
type ZSClient struct {
	config     *ZSConfig
	httpClient *http.Client
	sessionId  string
}

// JobView job inventory view
type JobView struct {
	UUID       string      `json:"uuid"`
	State      string      `json:"state"`
	Result     interface{} `json:"result,omitempty"`
	Error      interface{} `json:"error,omitempty"`
	CreateDate string      `json:"createDate"`
}

const (
	JobStateProcessing = "Processing"
	JobStateSucceeded  = "Succeeded"
	JobStateFailed     = "Failed"
)

// NewZSClient creates a new ZStack client
func NewZSClient(config *ZSConfig) *ZSClient {
	return &ZSClient{
		config: config,
		httpClient: &http.Client{
			Timeout: config.timeout,
		},
	}
}

func (cli *ZSClient) baseURL() string {
	return fmt.Sprintf("http://%s:%d%s", cli.config.hostname, cli.config.port, cli.config.contextPath)
}

// Get performs a GET request
func (cli *ZSClient) Get(path string, uuid string, params interface{}, result interface{}) error {
	url := fmt.Sprintf("%s/%s", cli.baseURL(), path)
	if uuid != "" {
		url = fmt.Sprintf("%s/%s", url, uuid)
	}
	return cli.doRequest("GET", url, nil, result)
}

func (cli *ZSClient) QueryJob(uuid string) (*JobView, error) {
	var resp JobView
	url := fmt.Sprintf("%s/v1/api-jobs/%s", cli.baseURL(), uuid)
	err := cli.doRequest("GET", url, nil, &resp)
	return &resp, err
}

// List performs a list query
func (cli *ZSClient) List(path string, params interface{}, result interface{}) error {
	baseURL := cli.baseURL()
	requestURL := fmt.Sprintf("%s/%s", baseURL, path)

	if params != nil {
		if queryParam, ok := params.(*param.QueryParam); ok {
			queryString := cli.buildQueryString(queryParam)
			if queryString != "" {
				requestURL = fmt.Sprintf("%s?%s", requestURL, queryString)
			}
		}
	}

	return cli.doRequest("GET", requestURL, nil, result)
}

// Post performs a POST request
func (cli *ZSClient) Post(path string, params interface{}, result interface{}) error {
	url := fmt.Sprintf("%s/%s", cli.baseURL(), path)
	return cli.doRequest("POST", url, params, result)
}

// Put performs a PUT request
func (cli *ZSClient) Put(path string, uuid string, params interface{}, result interface{}) error {
	url := fmt.Sprintf("%s/%s/%s", cli.baseURL(), path, uuid)
	return cli.doRequest("PUT", url, params, result)
}

// Delete performs a DELETE request
func (cli *ZSClient) Delete(path string, uuid string, deleteMode string) error {
	url := fmt.Sprintf("%s/%s/%s?deleteMode=%s", cli.baseURL(), path, uuid, deleteMode)
	return cli.doRequest("DELETE", url, nil, nil)
}

func (cli *ZSClient) doRequest(method, url string, body interface{}, result interface{}) error {
	var bodyReader io.Reader
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			return err
		}
		bodyReader = bytes.NewBuffer(jsonBody)
	}

	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	cli.addAuthHeaders(req)

	resp, err := cli.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == 202 {
		var location struct {
			Location string `json:"location"`
			Uuid     string `json:"org.zstack.header.rest.APIEvent/uuid"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&location); err != nil {
			return fmt.Errorf("failed to decode 202 response: %v", err)
		}
		jobUUID := location.Uuid
		if jobUUID == "" {
			parts := bytes.Split([]byte(location.Location), []byte("/"))
			if len(parts) > 0 {
				jobUUID = string(parts[len(parts)-1])
			}
		}

		if jobUUID == "" {
			return fmt.Errorf("failed to extract job uuid from 202 response")
		}

		return cli.waitForJob(jobUUID, result)
	}

	if resp.StatusCode >= 400 {
		return fmt.Errorf("API error: %s %s returned status code %d", method, url, resp.StatusCode)
	}

	if result != nil {
		return json.NewDecoder(resp.Body).Decode(result)
	}
	return nil
}

func (cli *ZSClient) waitForJob(jobUUID string, result interface{}) error {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	timeout := time.After(30 * time.Minute)

	for {
		select {
		case <-timeout:
			return fmt.Errorf("job %s timeout", jobUUID)
		case <-ticker.C:
			job, err := cli.QueryJob(jobUUID)
			if err != nil {
				continue
			}

			if job.State == JobStateSucceeded {
				if result != nil && job.Result != nil {
					data, err := json.Marshal(job.Result)
					if err != nil {
						return fmt.Errorf("failed to marshal job result: %v", err)
					}
					return json.Unmarshal(data, result)
				}
				return nil
			}

			if job.State == JobStateFailed {
				return fmt.Errorf("job failed: %v", job.Error)
			}
		}
	}
}

func (cli *ZSClient) buildQueryString(params *param.QueryParam) string {
	u := url.Values{}

	for _, q := range params.Conditions {
		if q.Value != "" {
			u.Add("q", q.Value)
		}
	}

	if params.LimitNum != nil {
		u.Set("limit", strconv.Itoa(*params.LimitNum))
	}
	if params.StartNum != nil {
		u.Set("start", strconv.Itoa(*params.StartNum))
	}
	if params.Count {
		u.Set("count", "true")
	}
	if params.ReplyWithCount {
		u.Set("replyWithCount", "true")
	}
	if params.GroupBy != "" {
		u.Set("groupBy", params.GroupBy)
	}
	if params.SortBy != "" {
		u.Set("sortBy", params.SortBy)
	}
	if params.SortDirection != "" {
		u.Set("sortDirection", params.SortDirection)
	}
	for _, f := range params.Fields {
		u.Add("fields", f)
	}

	return u.Encode()
}

func (cli *ZSClient) addAuthHeaders(req *http.Request) {
	if cli.config.authType == AuthTypeAccessKey {
		req.Header.Set("X-Access-Key-Id", cli.config.accessKeyId)
		req.Header.Set("X-Access-Key-Secret", cli.config.accessKeySecret)
	} else if cli.sessionId != "" {
		req.Header.Set("Authorization", "OAuth "+cli.sessionId)
	}
}
