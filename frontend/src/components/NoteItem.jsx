import { useState } from 'react';
import NoteEditor from './NoteEditor.jsx';
import MarkdownPreview from './MarkdownPreview.jsx';

export default function NoteItem({ note, onUpdate, onDelete }) {
  const [editing, setEditing] = useState(false);
  if (editing) return <article className="note"><NoteEditor initialValue={`# ${note.title}\n${note.content}`} onSave={async (text) => { await onUpdate(note.id, text); setEditing(false); }} onCancel={() => setEditing(false)} /></article>;
  const preview = note.content.length > 140 ? `${note.content.slice(0, 140)}…` : note.content;
  return <article className="note"><div className="note-head"><h2>{note.title}</h2><time>{new Date(note.created_at).toLocaleString('ja-JP')}</time></div><MarkdownPreview content={preview} /><p className="urls">URL: {note.urls.length > 0 ? `${note.urls.length}件` : 'なし'}</p><div className="actions"><button onClick={() => setEditing(true)}>編集</button><button className="danger" onClick={() => onDelete(note.id)}>削除</button></div></article>;
}
