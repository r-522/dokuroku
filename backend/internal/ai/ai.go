package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Result struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Client interface {
	Structure(ctx context.Context, raw string) (Result, error)
}

type ClaudeClient struct {
	APIKey     string
	HTTPClient *http.Client
	Timeout    time.Duration
}

func (c ClaudeClient) Structure(ctx context.Context, raw string) (Result, error) {
	if c.APIKey == "" {
		return Result{}, errors.New("anthropic api key is empty")
	}
	if c.Timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, c.Timeout)
		defer cancel()
	}
	body := map[string]any{
		"model":      "claude-haiku-4.5",
		"max_tokens": 2048,
		"system":     systemPrompt,
		"messages": []map[string]string{{
			"role":    "user",
			"content": raw,
		}},
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return Result{}, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "https://api.anthropic.com/v1/messages", bytes.NewReader(payload))
	if err != nil {
		return Result{}, err
	}
	req.Header.Set("content-type", "application/json")
	req.Header.Set("x-api-key", c.APIKey)
	req.Header.Set("anthropic-version", "2023-06-01")
	client := c.HTTPClient
	if client == nil {
		client = http.DefaultClient
	}
	res, err := client.Do(req)
	if err != nil {
		return Result{}, err
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return Result{}, fmt.Errorf("anthropic status: %d", res.StatusCode)
	}
	var decoded struct {
		Content []struct {
			Text string `json:"text"`
		} `json:"content"`
	}
	if err := json.NewDecoder(res.Body).Decode(&decoded); err != nil {
		return Result{}, err
	}
	if len(decoded.Content) == 0 {
		return Result{}, errors.New("empty ai response")
	}
	var result Result
	if err := json.Unmarshal([]byte(strings.TrimSpace(decoded.Content[0].Text)), &result); err != nil {
		return Result{}, err
	}
	if strings.TrimSpace(result.Content) == "" {
		return Result{}, errors.New("empty ai content")
	}
	return result, nil
}

const systemPrompt = `あなたは読録の入力整形補助です。JSONだけを返してください。本文を要約、省略、添削、言い換えしてはいけません。入力に存在しないURLや情報を追加してはいけません。形式は{"title":"一覧識別用の短いタイトル","content":"Markdown本文"}です。`
