CREATE TABLE IF NOT EXISTS BOOK_notes (
  id BIGSERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  content TEXT NOT NULL,
  urls TEXT[] NOT NULL DEFAULT '{}',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  CONSTRAINT book_notes_content_not_blank CHECK (length(btrim(content)) > 0)
);

CREATE INDEX IF NOT EXISTS idx_BOOK_notes_created_at ON BOOK_notes (created_at DESC);
