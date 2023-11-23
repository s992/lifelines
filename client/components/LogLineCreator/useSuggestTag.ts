import { getHotkeyHandler } from '@mantine/hooks';
import { ChangeEventHandler, RefObject, useState } from 'react';

import { Tag } from '../../generated/proto/lifelines/v1/lifelines_pb';
import { vars } from '../../theme';
import { getTagColor } from '../../util/getTagColor';

export function useSuggestTag(
  inputRef: RefObject<HTMLInputElement>,
  tags: Tag[],
) {
  const [inputValue, setInputValue] = useState('');
  const [tagName, ...rest] = inputValue.split(' ');
  let matchingTag: Tag | undefined;
  let suggestedTag: Tag | undefined;

  if (rest.length) {
    matchingTag = tags.find(({ name }) => name === tagName);
  } else if (tagName.length) {
    suggestedTag = tags.find(({ name }) => name.startsWith(tagName));
  }

  const matchedTag = matchingTag ?? suggestedTag;
  const tagColor = matchedTag
    ? getTagColor(matchedTag)
    : vars.colors.blue.filled;
  const suggestionText = suggestedTag?.name
    .slice(inputValue.length)
    .padStart(suggestedTag.name.length, ' ');

  const reset = () => {
    setInputValue('');
  };

  const maybeAcceptSuggestion = () => {
    if (!suggestedTag || inputValue.length >= suggestedTag.name.length) {
      return;
    }

    setInputValue(`${suggestedTag.name} `);
  };

  const onChange: ChangeEventHandler<HTMLInputElement> = (e) => {
    setInputValue(e.target.value);
  };

  const onKeyDown = getHotkeyHandler([
    [
      'Escape',
      () => {
        reset();
        inputRef.current?.blur();
      },
    ],
    ['Tab', maybeAcceptSuggestion],
  ]);

  return {
    inputValue,
    matchedTag,
    suggestionText,
    tagColor,
    reset,
    inputEventHandlers: {
      onChange,
      onKeyDown,
      onTouchEnd: maybeAcceptSuggestion,
    },
  };
}
