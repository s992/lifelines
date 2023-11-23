import { TransportProvider } from '@connectrpc/connect-query';
import { createConnectTransport } from '@connectrpc/connect-web';
import { MantineProvider } from '@mantine/core';
import '@mantine/core/styles.css';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import React from 'react';
import ReactDOM from 'react-dom/client';

import { App } from './App';
import './hacks';
import { theme } from './theme';

const transport = createConnectTransport({ baseUrl: 'http://localhost:8080' });
const queryClient = new QueryClient();

// eslint-disable-next-line @typescript-eslint/no-non-null-assertion
ReactDOM.createRoot(document.getElementById('root')!).render(
  <TransportProvider transport={transport}>
    <QueryClientProvider client={queryClient}>
      <MantineProvider theme={theme}>
        <React.StrictMode>
          <App />
        </React.StrictMode>
      </MantineProvider>
    </QueryClientProvider>
  </TransportProvider>,
);
