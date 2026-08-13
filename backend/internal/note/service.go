package note

import (
	"context"
	"strings"

	"dokuroku/backend/internal/ai"
	"dokuroku/backend/internal/urlutil"
)

type Service struct {
	Repo Repository
	AI   ai.Client
}

func (s Service) Create(ctx context.Context, raw string) (CreateResult, error) {
	title, content, explicit := splitExplicitTitle(raw)
	fallback := false
	if !explicit && s.AI != nil {
		result, err := s.AI.Structure(ctx, raw)
		if err == nil && strings.TrimSpace(result.Content) != "" && preservesInput(raw, result.Content) {
			title = strings.TrimSpace(result.Title)
			content = result.Content
		} else {
			fallback = true
		}
	}
	if strings.TrimSpace(title) == "" {
		title = "無題"
	}
	if strings.TrimSpace(content) == "" {
		content = raw
	}
	n, err := s.Repo.Create(ctx, title, content, urlutil.Extract(raw))
	return CreateResult{Note: n, AIFallback: fallback}, err
}

func (s Service) Update(ctx context.Context, id int64, raw string) (Note, error) {
	title, content, _ := splitExplicitTitle(raw)
	if strings.TrimSpace(title) == "" {
		title = "無題"
	}
	if strings.TrimSpace(content) == "" {
		content = raw
	}
	return s.Repo.Update(ctx, id, title, content, urlutil.Extract(raw))
}

func (s Service) List(ctx context.Context, query string) ([]Note, error) {
	return s.Repo.List(ctx, query)
}
func (s Service) Delete(ctx context.Context, id int64) error { return s.Repo.Delete(ctx, id) }

func splitExplicitTitle(raw string) (string, string, bool) {
	text := strings.ReplaceAll(raw, "\r\n", "\n")
	lines := strings.Split(text, "\n")
	if len(lines) == 0 {
		return "", raw, false
	}
	first := strings.TrimSpace(lines[0])
	if strings.HasPrefix(first, "# ") {
		return strings.TrimSpace(strings.TrimPrefix(first, "# ")), strings.Join(lines[1:], "\n"), true
	}
	return "", raw, false
}

func preservesInput(raw, content string) bool {
	normalize := func(s string) string {
		return strings.Join(strings.Fields(s), " ")
	}
	return strings.Contains(normalize(content), normalize(raw))
}
