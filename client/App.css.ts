import { style } from '@vanilla-extract/css';

import { vars } from './theme';

export const container = style({
  width: '70%',
  padding: vars.spacing.lg,
  '@media': {
    [`(max-width: ${vars.breakpoints.xs})`]: {
      width: '100%',
      padding: vars.spacing.xs,
    },
  },
});
