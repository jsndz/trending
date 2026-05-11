import { useState, useEffect, useCallback } from 'react';
import { fetchArticles } from '../api/client';
import type { Article } from '../api/types';

interface UseArticlesResult {
  articles: Article[];
  loading: boolean;
  loadingMore: boolean;
  error: string | null;
  hasMore: boolean;
  refetch: () => void;
  loadMore: () => void;
}

const LIMIT = 20;

export function useArticles(): UseArticlesResult {
  const [articles, setArticles] = useState<Article[]>([]);
  const [loading, setLoading] = useState(true);
  const [loadingMore, setLoadingMore] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [page, setPage] = useState(1);
  const [hasMore, setHasMore] = useState(true);
  const [tick, setTick] = useState(0);

  const fetchInitial = useCallback(async () => {
    setLoading(true);
    setError(null);
    setPage(1);
    try {
      const data = await fetchArticles(1, LIMIT);
      setArticles(data);
      setHasMore(data.length === LIMIT);
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : 'An unknown error occurred');
    } finally {
      setLoading(false);
    }
  }, []);

  useEffect(() => {
    fetchInitial();
  }, [fetchInitial, tick]);

  const loadMore = useCallback(async () => {
    if (loading || loadingMore || !hasMore) return;

    setLoadingMore(true);
    const nextPage = page + 1;
    try {
      const data = await fetchArticles(nextPage, LIMIT);
      if (data.length > 0) {
        setArticles((prev) => [...prev, ...data]);
        setPage(nextPage);
      }
      setHasMore(data.length === LIMIT);
    } catch (err: unknown) {
      setError(err instanceof Error ? err.message : 'An unknown error occurred');
    } finally {
      setLoadingMore(false);
    }
  }, [loading, loadingMore, hasMore, page]);

  return {
    articles,
    loading,
    loadingMore,
    error,
    hasMore,
    refetch: () => setTick((t) => t + 1),
    loadMore
  };
}
