import { Input } from '@mantine/core';
import { useState } from 'react';

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
  const [inputValue, setInputValue] = useState('');
  const [suggestedTag, setSuggestedTag] = useState<Tag | null>(null);

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
      <div className={wrapper}>
        <Input
          variant="unstyled"
          leftSection="#"
          type="text"
          autoComplete="off"
          value={inputValue}
          onChange={(e) => {
            const value = e.target.value;
            const suggestion = value
              ? tags.find(({ name }) => name.startsWith(value))
              : null;

            setInputValue(value);
            setSuggestedTag(suggestion ?? null);
          }}
          onKeyDown={(e) => {
            if (e.key !== 'Tab' || !suggestedTag) {
              return;
            }

            e.preventDefault();
            const nextValue = `${suggestedTag.name} `;

            setInputValue(nextValue);
            setSuggestedTag(null);
          }}
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
