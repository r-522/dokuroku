import { useState } from 'react';
import MarkdownPreview from './MarkdownPreview.jsx';

export default function NoteEditor({ initialValue = '', onSave, onCancel }) {
  const [text, setText] = useState(initialValue);
  const [saving, setSaving] = useState(false);
  async function submit(e) {
    e.preventDefault();
    if (!text.trim()) return;
    setSaving(true);
    await onSave(text);
    setText('');
    setSaving(false);
  }
  return <form className="editor" onSubmit={submit}>
    <textarea value={text} onChange={(e) => setText(e.target.value)} placeholder="# タイトル（任意）&#10;読んだものから残ったことを書く" />
    <MarkdownPreview content={text} />
    <div className="actions"><button disabled={saving || !text.trim()}>{saving ? '保存中' : '保存'}</button>{onCancel && <button type="button" onClick={onCancel}>キャンセル</button>}</div>
  </form>;
}
