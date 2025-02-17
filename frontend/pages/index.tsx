import React, { useState, useEffect } from 'react';
import MemoForm from '../components/MemoForm';
import MemoList from '../components/MemoList';
import { Memo } from '../types/memo';

const Home = () => {
  const [memos, setMemos] = useState<Memo[]>([]);

  // メモ一覧を取得する関数
  const fetchMemos = async () => {
    try {
      const response = await fetch('/api/memos');
      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Network response was not ok');
      }
      const data = await response.json();
      setMemos(data);
    } catch (error: any) {
      console.error('Error fetching memos:', error.message);
    }
  };

  // コンポーネントのマウント時にメモ一覧を取得
  useEffect(() => {
    fetchMemos();
  }, []);

  const handleSubmit = async (content: string) => {
    try {
      const response = await fetch('/api/memos', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({ content }),
      });

      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.error || 'Network response was not ok');
      }

      // メモ保存後に一覧を再取得
      fetchMemos();
    } catch (error: any) {
      console.error('Error submitting memo:', error.message);
    }
  };

  return (
    <div>
      <h1>Memoru</h1>
      <MemoForm onSubmit={handleSubmit} />
      <MemoList memos={memos} />
    </div>
  );
};

export default Home;