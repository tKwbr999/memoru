import React from 'react';
import MemoForm from '../components/MemoForm';
import MemoList from '../components/MemoList';

const Home = () => {
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

      // TODO: レスポンスを処理する（例：メモ一覧を更新する）
    } catch (error: any) {
      console.error('Error submitting memo:', error.message);
    }
  };

  return (
    <div>
      <h1>Memoru</h1>
      <MemoForm onSubmit={handleSubmit} />
      <MemoList memos={[{ id: 1, content: '仮のメモ1' }, { id: 2, content: '仮のメモ2' }]} />
    </div>
  );
};

export default Home;