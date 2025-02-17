import React from 'react';

interface Memo {
  id: string;
  content: string;
}

interface Props {
  memos: Memo[];
}

const MemoList: React.FC<Props> = ({ memos }) => {
  return (
    <ul>
      {memos.map((memo) => (
        <li key={memo.id}>
          {memo.content}
        </li>
      ))}
    </ul>
  );
};

export default MemoList;