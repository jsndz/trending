import { TrendingUp, RefreshCw } from 'lucide-react';

interface Props {
  total: number;
  loading: boolean;
  onRefresh: () => void;
}

export function Header({ total, loading, onRefresh }: Props) {
  return (
    <header className="sticky top-0 z-20 bg-white/90 backdrop-blur-md border-b border-gray-100 shadow-sm">
      <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 h-16 flex items-center justify-between gap-4">
        <div className="flex items-center gap-3">
          <div className="w-8 h-8 bg-blue-600 rounded-lg flex items-center justify-center shadow-sm">
            <TrendingUp size={18} className="text-white" />
          </div>
          <div>
            <h1 className="text-gray-900 font-bold text-lg leading-none">
              Trending
            </h1>
            <p className="text-gray-400 text-xs leading-none mt-0.5">
              Content Aggregator
            </p>
          </div>
        </div>

        <div className="flex items-center gap-4">
          {total > 0 && (
            <span className="hidden sm:block text-sm text-gray-400">
              <span className="font-semibold text-gray-700">{total}</span> articles
            </span>
          )}
          <button
            onClick={onRefresh}
            disabled={loading}
            className="inline-flex items-center gap-2 text-sm font-medium text-gray-600 hover:text-blue-600 disabled:opacity-40 transition-colors px-3 py-1.5 rounded-lg hover:bg-blue-50"
            aria-label="Refresh articles"
          >
            <RefreshCw size={15} className={loading ? 'animate-spin' : ''} />
            <span className="hidden sm:inline">Refresh</span>
          </button>
        </div>
      </div>
    </header>
  );
}
