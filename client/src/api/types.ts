export interface Category {
  id: string;
  name: string;
}

export interface Article {
  id: string;
  title: string;
  publishedAt: string;
  link: string;
  author?: string;
  source: string;
  description?: string;
  category: Category[];
}

export type ArticlesResponse = Article[];
