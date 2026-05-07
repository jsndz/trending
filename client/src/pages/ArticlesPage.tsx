import { useMemo, useState } from 'react';
import { AlertCircle, TrendingUp, Frown } from 'lucide-react';
import { useArticles } from '../hooks/useArticles';
import { ArticleCard } from '../components/ArticleCard';
import { Header } from '../components/Header';
import { FilterBar } from '../components/FilterBar';
import { SkeletonCard } from '../components/SkeletonCard';

export function ArticlesPage() {
  const { articles, loading, error, refetch } = useArticles();
  const [query, setQuery] = useState('');
  const [activeSource, setActiveSource] = useState('');

  const sources = useMemo(
    () => Array.from(new Set(articles.map((a) => a.source))).sort(),
    [articles]
  );

  const filtered = useMemo(() => {
    let result = articles;
    if (activeSource) {
      result = result.filter((a) => a.source === activeSource);
    }
    if (query.trim()) {
      const q = query.toLowerCase();
      result = result.filter(
        (a) =>
          a.title.toLowerCase().includes(q) ||
          a.description?.toLowerCase().includes(q) ||
          a.author?.toLowerCase().includes(q) ||
          a.category.some((c) => c.name.toLowerCase().includes(q))
      );
    }
    return result;
  }, [articles, query, activeSource]);

  return (
    <div className="min-h-screen bg-gray-50">
      <Header total={articles.length} loading={loading} onRefresh={refetch} />

      <main className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {/* Hero strip */}
        <div className="mb-8">
          <div className="flex items-center gap-2 mb-1">
            <TrendingUp size={16} className="text-blue-600" />
            <span className="text-xs font-semibold uppercase tracking-widest text-blue-600">
              Live Feed
            </span>
          </div>
          <h2 className="text-2xl sm:text-3xl font-bold text-gray-900 leading-tight">
            What's trending right now
          </h2>
          <p className="text-gray-500 text-sm mt-1">
            Normalized articles aggregated from multiple sources in real time.
          </p>
        </div>

        {/* Filters */}
        {!loading && !error && articles.length > 0 && (
          <div className="mb-6">
            <FilterBar
              query={query}
              onQueryChange={setQuery}
              sources={sources}
              activeSource={activeSource}
              onSourceChange={setActiveSource}
            />
          </div>
        )}

        {/* Error state */}
        {error && (
          <div className="flex flex-col items-center justify-center py-24 gap-4 text-center">
            <div className="w-14 h-14 bg-red-50 rounded-2xl flex items-center justify-center">
              <AlertCircle size={28} className="text-red-500" />
            </div>
            <div>
              <p className="font-semibold text-gray-800 text-lg">Could not load articles</p>
              <p className="text-gray-500 text-sm mt-1 max-w-xs">{error}</p>
            </div>
            <button
              onClick={refetch}
              className="mt-2 px-5 py-2 bg-blue-600 text-white text-sm font-medium rounded-xl hover:bg-blue-700 transition-colors shadow-sm"
            >
              Try again
            </button>
          </div>
        )}

        {/* Loading skeleton */}
        {loading && (
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-5">
            {Array.from({ length: 9 }).map((_, i) => (
              <SkeletonCard key={i} />
            ))}
          </div>
        )}

        {/* Empty state after filtering */}
        {!loading && !error && filtered.length === 0 && articles.length > 0 && (
          <div className="flex flex-col items-center justify-center py-24 gap-3 text-center">
            <Frown size={36} className="text-gray-300" />
            <p className="font-semibold text-gray-700">No articles match your filters</p>
            <button
              onClick={() => { setQuery(''); setActiveSource(''); }}
              className="text-sm text-blue-600 hover:underline"
            >
              Clear filters
            </button>
          </div>
        )}

        {/* Empty state — no data at all */}
        {!loading && !error && articles.length === 0 && (
          <div className="flex flex-col items-center justify-center py-24 gap-3 text-center">
            <TrendingUp size={36} className="text-gray-200" />
            <p className="font-semibold text-gray-700">No articles yet</p>
            <p className="text-gray-400 text-sm">Check back soon or refresh the feed.</p>
          </div>
        )}

        {/* Article grid */}
        {!loading && !error && filtered.length > 0 && (
          <>
            <p className="text-xs text-gray-400 mb-4">
              Showing <span className="font-medium text-gray-600">{filtered.length}</span>
              {filtered.length !== articles.length && (
                <> of <span className="font-medium text-gray-600">{articles.length}</span></>
              )}{' '}
              articles
            </p>
            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-5">
              {filtered.map((article, i) => (
                <ArticleCard key={article.id} article={article} index={i} />
              ))}
            </div>
          </>
        )}
      </main>
    </div>
  );
}
