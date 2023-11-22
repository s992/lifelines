import { vanillaExtractPlugin } from '@vanilla-extract/vite-plugin';
import react from '@vitejs/plugin-react-swc';
import { defineConfig } from 'vite';

// eslint-disable-next-line import/no-default-export
export default defineConfig({
  root: './client',
  plugins: [react(), vanillaExtractPlugin()],
  build: {
    outDir: '../dist/client',
    emptyOutDir: true,
  },
});
