import { Input } from '@mantine/core';
import { getHotkeyHandler, useFocusWithin, useHotkeys } from '@mantine/hooks';
import { useRef, useState } from 'react';

import {
  CreateLogLineRequest,
  Tag,
} from '../../generated/proto/logger/v1/logger_pb';
import { completion, input, inputWrapper, wrapper } from './LogLineCreator.css';

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
  const [inputValue, setInputValue] = useState('');
  const [suggestedTag, setSuggestedTag] = useState<Tag | null>(null);
  const focus = () => inputRef.current?.focus();

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
            setInputValue('');
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
          leftSection="#"
          type="text"
          autoComplete="off"
          value={inputValue}
          placeholder={focused ? 'tag value (optional description)' : undefined}
          onChange={(e) => {
            const value = e.target.value;
            const suggestion = value
              ? tags.find(({ name }) => name.startsWith(value))
              : null;

            setInputValue(value);
            setSuggestedTag(suggestion ?? null);
          }}
          onKeyDown={getHotkeyHandler([
            [
              'Escape',
              () => {
                setInputValue('');
                setSuggestedTag(null);
                inputRef.current?.blur();
              },
            ],
            [
              'Tab',
              () => {
                if (!suggestedTag) {
                  return;
                }

                const nextValue = `${suggestedTag.name} `;

                setInputValue(nextValue);
                setSuggestedTag(null);
              },
            ],
          ])}
          classNames={{
            input,
            wrapper: inputWrapper,
          }}
        />
        <div className={completion}>
          {suggestedTag?.name
            .slice(inputValue.length)
            .padStart(suggestedTag.name.length, ' ')}
        </div>
      </div>
    </form>
  );
}
