import { useQuery } from '@tanstack/react-query';

import { listTags } from './generated/proto/logger/v1/logger-TagService_connectquery';

export function App() {
  const { data } = useQuery(listTags.useQuery({}));

  return (
    <div>{data?.tags.map((tag) => <p key={tag.tagId}>{tag.name}</p>)}</div>
  );
}
