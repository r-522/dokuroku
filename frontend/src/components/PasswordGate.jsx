import { useState } from 'react';
import { api } from '../api.js';

export default function PasswordGate({ onLogin }) {
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  async function submit(e) {
    e.preventDefault();
    setError('');
    try {
      await api.login(password);
      onLogin();
    } catch (err) {
      setError(err.message);
    }
  }
  return <main className="gate"><form onSubmit={submit}><h1>読録</h1><label>パスワード<input type="password" value={password} onChange={(e) => setPassword(e.target.value)} /></label><button>入る</button>{error && <p className="error">{error}</p>}</form></main>;
}
