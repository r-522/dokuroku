import DOMPurify from 'dompurify';
import { marked } from 'marked';

export default function MarkdownPreview({ content }) {
  const html = DOMPurify.sanitize(marked.parse(content || ''));
  return <div className="preview" dangerouslySetInnerHTML={{ __html: html }} />;
}
