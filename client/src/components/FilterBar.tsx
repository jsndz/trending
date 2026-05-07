import { Search, X } from 'lucide-react';

interface Props {
  query: string;
  onQueryChange: (q: string) => void;
  sources: string[];
  activeSource: string;
  onSourceChange: (s: string) => void;
}

export function FilterBar({ query, onQueryChange, sources, activeSource, onSourceChange }: Props) {
  return (
    <div className="flex flex-col sm:flex-row gap-3">
      {/* Search */}
      <div className="relative flex-1 max-w-md">
        <Search size={15} className="absolute left-3 top-1/2 -translate-y-1/2 text-gray-400 pointer-events-none" />
        <input
          type="text"
          placeholder="Search articles..."
          value={query}
          onChange={(e) => onQueryChange(e.target.value)}
          className="w-full pl-9 pr-8 py-2 text-sm bg-white border border-gray-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-blue-500/30 focus:border-blue-400 transition placeholder-gray-400"
        />
        {query && (
          <button
            onClick={() => onQueryChange('')}
            className="absolute right-2.5 top-1/2 -translate-y-1/2 text-gray-400 hover:text-gray-600"
          >
            <X size={14} />
          </button>
        )}
      </div>

      {/* Source filter pills */}
      {sources.length > 1 && (
        <div className="flex items-center gap-2 flex-wrap">
          <button
            onClick={() => onSourceChange('')}
            className={`text-xs font-medium px-3 py-1.5 rounded-full border transition-all ${
              activeSource === ''
                ? 'bg-blue-600 text-white border-blue-600 shadow-sm'
                : 'bg-white text-gray-600 border-gray-200 hover:border-blue-300 hover:text-blue-600'
            }`}
          >
            All
          </button>
          {sources.map((src) => (
            <button
              key={src}
              onClick={() => onSourceChange(src === activeSource ? '' : src)}
              className={`text-xs font-medium px-3 py-1.5 rounded-full border transition-all ${
                activeSource === src
                  ? 'bg-blue-600 text-white border-blue-600 shadow-sm'
                  : 'bg-white text-gray-600 border-gray-200 hover:border-blue-300 hover:text-blue-600'
              }`}
            >
              {src}
            </button>
          ))}
        </div>
      )}
    </div>
  );
}
