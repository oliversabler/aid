package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type LLMClient struct {
	cfg        *Config
	httpClient *http.Client
}

type chatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chatRequest struct {
	Model       string        `json:"model"`
	Messages    []chatMessage `json:"messages"`
	MaxTokens   int           `json:"max_tokens,omitempty"`
	Temperature float64       `json:"temperature,omitempty"`
}

type chatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
	Error *struct {
		Message string `json:"message"`
	} `json:"error,omitempty"`
}

func NewLLMClient(cfg *Config) *LLMClient {
	return &LLMClient{
		cfg: cfg,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

func (c *LLMClient) chat(systemPrompt, userMessage string) (string, error) {
	reqBody := chatRequest{
		Model: c.cfg.API.Model,
		Messages: []chatMessage{
			{Role: "system", Content: systemPrompt},
			{Role: "user", Content: userMessage},
		},
		MaxTokens:   256,
		Temperature: 0.3,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", c.cfg.API.Endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("API request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
	}

	var chatResp chatResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}

	if chatResp.Error != nil {
		return "", fmt.Errorf("API error: %s", chatResp.Error.Message)
	}

	if len(chatResp.Choices) == 0 {
		return "", fmt.Errorf("no response from API")
	}

	return chatResp.Choices[0].Message.Content, nil
}

const prevSystemPrompt = `You are a concise terminal assistant. The user will provide a shell command that likely failed or produced unexpected results.

Your task:
1. Analyze what might have gone wrong
2. Explain the issue in 1-2 sentences maximum
3. If applicable, provide the corrected command

Keep your response extremely brief. No introductions or pleasantries. Focus only on the problem and solution.`

const questionSystemPrompt = `You are a concise terminal assistant. The user will ask how to accomplish a task using the command line.

Your task:
1. Provide the exact command(s) needed
2. Add a brief explanation only if the command is complex

Keep your response to 1-3 lines maximum. No introductions or pleasantries. Just the command and minimal explanation.`

func (c *LLMClient) AnalyzePreviousCommand(command string) (string, error) {
	return c.chat(prevSystemPrompt, "The previous command was:\n"+command)
}

func (c *LLMClient) Question(q string) (string, error) {
	return c.chat(questionSystemPrompt, q)
}
