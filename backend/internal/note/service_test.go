package note

import (
	"context"
	"errors"
	"testing"
	"time"

	"dokuroku/backend/internal/ai"
)

type fakeRepo struct{ saved Note }

func (r *fakeRepo) Create(ctx context.Context, title, content string, urls []string) (Note, error) {
	r.saved = Note{ID: 1, Title: title, Content: content, URLs: urls, CreatedAt: time.Now()}
	return r.saved, nil
}
func (r *fakeRepo) List(ctx context.Context, query string) ([]Note, error) { return nil, nil }
func (r *fakeRepo) Update(ctx context.Context, id int64, title, content string, urls []string) (Note, error) {
	return Note{ID: id, Title: title, Content: content, URLs: urls}, nil
}
func (r *fakeRepo) Delete(ctx context.Context, id int64) error { return nil }

type fakeAI struct {
	result ai.Result
	err    error
}

func (f fakeAI) Structure(ctx context.Context, raw string) (ai.Result, error) { return f.result, f.err }

func TestNoteServiceCreate_UT_NOTE_004(t *testing.T) {
	repo := &fakeRepo{}
	_, err := (Service{Repo: repo}).Create(context.Background(), "# 月人の野望\nとても面白かった。")
	if err != nil {
		t.Fatal(err)
	}
	if repo.saved.Title != "月人の野望" {
		t.Fatalf("title = %q", repo.saved.Title)
	}
	if repo.saved.Content != "とても面白かった。" {
		t.Fatalf("content = %q", repo.saved.Content)
	}
}

func TestNoteServiceCreate_UT_NOTE_006(t *testing.T) {
	repo := &fakeRepo{}
	res, err := (Service{Repo: repo, AI: fakeAI{err: errors.New("down")}}).Create(context.Background(), "本文だけ")
	if err != nil {
		t.Fatal(err)
	}
	if !res.AIFallback || repo.saved.Title != "無題" || repo.saved.Content != "本文だけ" {
		t.Fatalf("res=%+v saved=%+v", res, repo.saved)
	}
}

func TestNoteServiceCreate_UT_NOTE_007(t *testing.T) {
	repo := &fakeRepo{}
	raw := "GoのHTTPサーバーについて読んだ。\nHandlerについて理解できた。"
	_, err := (Service{Repo: repo, AI: fakeAI{result: ai.Result{Title: "Go HTTP", Content: raw}}}).Create(context.Background(), raw)
	if err != nil {
		t.Fatal(err)
	}
	if repo.saved.Content != raw {
		t.Fatalf("content changed: %q", repo.saved.Content)
	}
}
