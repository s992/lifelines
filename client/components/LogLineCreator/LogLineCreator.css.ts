import { rem } from '@mantine/core';
import { style } from '@vanilla-extract/css';

import { vars } from '../../theme';

export const wrapper = style({
  position: 'relative',
});

const inputHeight = rem(48);

const base = style({
  fontFamily: 'monospace',
  fontSize: rem(16),
  lineHeight: inputHeight,
});

export const input = style([
  base,
  {
    border: `1px solid ${vars.colors.gray[3]}`,
    borderRadius: vars.radius.md,
    zIndex: 1,
    ':focus': {
      borderColor: vars.colors.gray.outline,
    },
    ':hover': {
      borderColor: vars.colors.gray.outline,
    },
  },
]);

export const inputWrapper = style({
  alignItems: 'center',
  display: 'flex',
  height: inputHeight,
  position: 'relative',
});

export const completion = style([
  base,
  {
    alignItems: 'center',
    border: '1px solid transparent',
    color: vars.colors.placeholder,
    display: 'flex',
    height: '100%',
    paddingLeft: rem(34),
    position: 'absolute',
    top: 0,
    whiteSpace: 'pre',
    width: '100%',
  },
]);
