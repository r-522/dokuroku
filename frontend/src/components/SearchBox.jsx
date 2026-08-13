export default function SearchBox({ value, onChange, onSearch }) {
  return <form className="search" onSubmit={(e) => { e.preventDefault(); onSearch(); }}><input value={value} onChange={(e) => onChange(e.target.value)} placeholder="検索" /><button>探す</button></form>;
}
