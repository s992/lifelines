import { style } from '@vanilla-extract/css';

export const wrapper = style({
  position: 'relative',
});

const inputHeight = '48px';

const base = style({
  fontFamily: 'monospace',
  fontSize: '1rem',
  lineHeight: inputHeight,
});

export const input = style([base]);

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
    display: 'flex',
    height: '100%',
    opacity: 0.3,
    paddingLeft: '34px',
    position: 'absolute',
    top: 0,
    whiteSpace: 'pre',
    width: '100%',
  },
]);
