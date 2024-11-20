import ReactMarkdown from 'react-markdown'

interface MdViewerProps {
  content: string
}

export default function MdViewer({ content }: MdViewerProps) {
  return (
    <div className="flex flex-col w-full max-w-3xl mx-auto p-4 sm:p-6">
      <div className="prose prose-slate dark:prose-invert max-w-none">
        <ReactMarkdown>{content}</ReactMarkdown>
      </div>
    </div>
  )
}
