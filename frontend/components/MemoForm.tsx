import React, { useState } from 'react';

interface Props {
  onSubmit: (content: string) => void;
}

const MemoForm: React.FC<Props> = ({ onSubmit }) => {
  const [content, setContent] = useState('');

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    if (content.trim() === '') return; // 空の場合は送信しない
    onSubmit(content);
    setContent(''); // 送信後、フォームをクリア
  };

  return (
    <form onSubmit={handleSubmit}>
      <textarea
        value={content}
        onChange={(e) => setContent(e.target.value)}
        placeholder="メモを入力..."
        rows={4}
        cols={50}
      />
      <button type="submit">保存</button>
    </form>
  );
};

export default MemoForm;