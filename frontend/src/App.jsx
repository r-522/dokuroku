import { useEffect, useState } from 'react';
import { api } from './api.js';
import PasswordGate from './components/PasswordGate.jsx';
import NoteEditor from './components/NoteEditor.jsx';
import SearchBox from './components/SearchBox.jsx';
import NoteList from './components/NoteList.jsx';

export default function App() {
  const [authed, setAuthed] = useState(false);
  const [notes, setNotes] = useState([]);
  const [query, setQuery] = useState('');
  const [message, setMessage] = useState('');

  async function load(q = query) {
    const data = await api.listNotes(q);
    setNotes(data);
  }

  useEffect(() => { if (authed) load().catch((e) => setMessage(e.message)); }, [authed]);

  if (!authed) return <PasswordGate onLogin={() => setAuthed(true)} />;

  return <main className="container">
    <header><h1>読録</h1><p>読んで残った感想や考えを、そのまま書いて保存します。</p></header>
    {message && <p className="message">{message}</p>}
    <NoteEditor onSave={async (text) => { await api.createNote(text); setMessage('保存しました'); await load(); }} />
    <SearchBox value={query} onChange={setQuery} onSearch={() => load(query)} />
    <NoteList notes={notes} onUpdate={async (id, text) => { await api.updateNote(id, text); await load(); }} onDelete={async (id) => { await api.deleteNote(id); await load(); }} />
  </main>;
}
