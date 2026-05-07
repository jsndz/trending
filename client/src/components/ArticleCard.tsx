import { ExternalLink, Calendar, User, Globe, Tag } from 'lucide-react';
import type { Article } from '../api/types';

function formatDate(iso: string): string {
  try {
    return new Intl.DateTimeFormat('en-US', {
      month: 'short',
      day: 'numeric',
      year: 'numeric',
      hour: '2-digit',
      minute: '2-digit',
    }).format(new Date(iso));
  } catch {
    return iso;
  }
}

interface Props {
  article: Article;
  index: number;
}

export function ArticleCard({ article, index }: Props) {
  return (
    <article
      className="group bg-white rounded-2xl border border-gray-100 shadow-sm hover:shadow-md transition-all duration-300 overflow-hidden flex flex-col"
      style={{ animationDelay: `${index * 40}ms` }}
    >
      <div className="p-6 flex flex-col gap-3 flex-1">
        {/* Source badge + trending rank */}
        <div className="flex items-center justify-between gap-2">
          <span className="inline-flex items-center gap-1.5 text-xs font-semibold uppercase tracking-wide text-blue-700 bg-blue-50 px-2.5 py-1 rounded-full">
            <Globe size={11} />
            {article.source}
          </span>
          <span className="text-xs text-gray-400 font-mono">#{index + 1}</span>
        </div>

        {/* Title */}
        <a
          href={article.link}
          target="_blank"
          rel="noopener noreferrer"
          className="group/link"
        >
          <h2 className="text-gray-900 font-semibold text-base leading-snug group-hover/link:text-blue-600 transition-colors duration-200 line-clamp-3">
            {article.title}
          </h2>
        </a>

        {/* Description */}
        {article.description && (
          <p className="text-gray-500 text-sm leading-relaxed line-clamp-3">
            {article.description}
          </p>
        )}

        {/* Categories */}
        {article.category.length > 0 && (
          <div className="flex items-center gap-1.5 flex-wrap">
            <Tag size={12} className="text-gray-400 shrink-0" />
            {article.category.slice(0, 4).map((cat) => (
              <span
                key={cat.id}
                className="text-xs text-gray-500 bg-gray-100 px-2 py-0.5 rounded-full"
              >
                {cat.name}
              </span>
            ))}
            {article.category.length > 4 && (
              <span className="text-xs text-gray-400">
                +{article.category.length - 4}
              </span>
            )}
          </div>
        )}
      </div>

      {/* Footer */}
      <div className="px-6 py-4 border-t border-gray-50 bg-gray-50/50 flex items-center justify-between gap-3">
        <div className="flex items-center gap-3 min-w-0">
          {article.author && (
            <span className="flex items-center gap-1.5 text-xs text-gray-500 min-w-0">
              <User size={12} className="shrink-0" />
              <span className="truncate max-w-[120px]">{article.author}</span>
            </span>
          )}
          <span className="flex items-center gap-1.5 text-xs text-gray-400">
            <Calendar size={12} className="shrink-0" />
            {formatDate(article.publishedAt)}
          </span>
        </div>

        <a
          href={article.link}
          target="_blank"
          rel="noopener noreferrer"
          className="shrink-0 inline-flex items-center gap-1.5 text-xs font-medium text-blue-600 hover:text-blue-700 transition-colors"
        >
          Read
          <ExternalLink size={12} />
        </a>
      </div>
    </article>
  );
}
