import NoteItem from './NoteItem.jsx';

export default function NoteList({ notes, onUpdate, onDelete }) {
  if (notes.length === 0) return <p className="empty">記録はまだありません。</p>;
  return <section className="list">{notes.map((note) => <NoteItem key={note.id} note={note} onUpdate={onUpdate} onDelete={onDelete} />)}</section>;
}
