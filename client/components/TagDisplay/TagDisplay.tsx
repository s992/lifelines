import { Anchor } from '@mantine/core';

import { Tag } from '../../generated/proto/logger/v1/logger_pb';
import { getTagColor } from './getTagColor';

type Props = {
  tag: Tag;
  onClick?: (tag: Tag) => void;
};

export function TagDisplay({ tag, onClick }: Props) {
  return (
    <Anchor
      component={onClick ? 'a' : 'span'}
      underline={onClick ? 'hover' : 'never'}
      style={{
        color: getTagColor(tag),
        cursor: onClick ? 'pointer' : 'text',
      }}
      onClick={() => {
        onClick?.(tag);
      }}
    >
      #{tag.name}
    </Anchor>
  );
}
