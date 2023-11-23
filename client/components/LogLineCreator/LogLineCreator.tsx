import { Input } from '@mantine/core';
import { useFocusWithin, useHotkeys } from '@mantine/hooks';
import { IconHash, IconPlus } from '@tabler/icons-react';
import { useRef } from 'react';

import {
  CreateLogLineRequest,
  Tag,
} from '../../generated/proto/lifelines/v1/lifelines_pb';
import { completion, input, inputWrapper, wrapper } from './LogLineCreator.css';
import { useSuggestTag } from './useSuggestTag';

type Props = {
  tags: Tag[];
  onCreate: (
    tagName: string,
    logLine: Partial<CreateLogLineRequest>,
  ) => Promise<void>;
};

export function LogLineCreator({ tags, onCreate }: Props) {
  const { ref: wrapperRef, focused } = useFocusWithin();
  const inputRef = useRef<HTMLInputElement>(null);
  const focus = () => inputRef.current?.focus();
  const {
    inputValue,
    matchedTag,
    suggestionText,
    tagColor,
    reset,
    inputEventHandlers,
  } = useSuggestTag(inputRef, tags);
  const PrefixIconCmp = !inputValue.length || matchedTag ? IconHash : IconPlus;

  useHotkeys([
    ['shift+Digit3', focus],
    ['i', focus],
    ['shift+i', focus],
  ]);

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();

        const [tagName, value, ...description] = inputValue.split(' ');
        const tag = tags.find(({ name }) => name === tagName);

        onCreate(tagName, {
          tagId: tag?.tagId,
          value: parseFloat(value),
          description: description.join(' '),
        })
          .then(() => {
            reset();
          })
          .catch((err) => {
            console.error(err);
          });
      }}
    >
      <div ref={wrapperRef} className={wrapper}>
        <Input
          ref={inputRef}
          variant="unstyled"
          leftSection={<PrefixIconCmp size={16} />}
          type="text"
          autoComplete="off"
          value={inputValue}
          placeholder={focused ? 'tag value (optional description)' : undefined}
          styles={{
            input: {
              color: tagColor,
            },
          }}
          classNames={{
            input,
            wrapper: inputWrapper,
          }}
          {...inputEventHandlers}
        />
        <div className={completion}>{suggestionText}</div>
      </div>
    </form>
  );
}
