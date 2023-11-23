import { rem } from '@mantine/core';
import { style } from '@vanilla-extract/css';

import { vars } from './theme';

export const container = style({
  minWidth: rem(960),
  padding: vars.spacing.lg,
});
