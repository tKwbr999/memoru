export interface Memo {
  id: string;
  content: string;
  created_at?: string; // Optional because it might not be available initially
}