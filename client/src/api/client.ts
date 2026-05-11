import axios from 'axios';
import type { ArticlesResponse } from './types';

const api = axios.create({
  baseURL: 'http://localhost:8080',
});

export async function fetchArticles(page: number = 1, limit: number = 20): Promise<ArticlesResponse> {
  const res = await api.get<ArticlesResponse>('/api/v1/articles', {
    params: { page, limit }
  });
  console.log(res.data);
  return res.data;
}

export async function ping(): Promise<{ status: string }> {
  const res = await api.get<{ status: string }>('/ping');
  return res.data;
}