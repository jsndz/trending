import { useState, useEffect } from 'react';
import { fetchArticles } from '../api/client';
import type { Article } from '../api/types';

interface UseArticlesResult {
  articles: Article[];
  loading: boolean;
  error: string | null;
  refetch: () => void;
}

export function useArticles(): UseArticlesResult {
  const [articles, setArticles] = useState<Article[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [tick, setTick] = useState(0);

  useEffect(() => {
    let cancelled = false;
    setLoading(true);
    setError(null);

    fetchArticles()
      .then((data:any) => {
        if (!cancelled) {
          setArticles(data);
        }
      })
      .catch((err: Error) => {
        if (!cancelled) {
          setError(err.message);
        }
      })
      .finally(() => {
        if (!cancelled) setLoading(false);
      });

       
    return () => { cancelled = true; };
  }, [tick]);

  return { articles, loading, error, refetch: () => setTick((t) => t + 1) };
}
