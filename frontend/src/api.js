async function request(path, options = {}) {
  const res = await fetch(path, { credentials: 'include', headers: { 'content-type': 'application/json', ...(options.headers || {}) }, ...options });
  if (!res.ok) {
    const body = await res.json().catch(() => ({ error: 'request failed' }));
    throw new Error(body.error || 'request failed');
  }
  if (res.status === 204) return null;
  return res.json();
}

export const api = {
  login: (password) => request('/api/login', { method: 'POST', body: JSON.stringify({ password }) }),
  listNotes: (q = '') => request(`/api/notes?q=${encodeURIComponent(q)}`),
  createNote: (rawText) => request('/api/notes', { method: 'POST', body: JSON.stringify({ raw_text: rawText }) }),
  updateNote: (id, rawText) => request(`/api/notes/${id}`, { method: 'PUT', body: JSON.stringify({ raw_text: rawText }) }),
  deleteNote: (id) => request(`/api/notes/${id}`, { method: 'DELETE' }),
};
