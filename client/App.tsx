import { Timestamp } from '@bufbuild/protobuf';
import { Stack } from '@mantine/core';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { subHours } from 'date-fns';
import { useState } from 'react';

import { LogLineCreator } from './components/LogLineCreator';
import { LogLineTable } from './components/LogLineTable';
import {
  createLogLine,
  listLogLines,
} from './generated/proto/logger/v1/logger-LogLineService_connectquery';
import {
  createTag,
  listTags,
} from './generated/proto/logger/v1/logger-TagService_connectquery';

export function App() {
  const queryClient = useQueryClient();
  const [now] = useState(new Date());
  const { data: tags } = useQuery(listTags.useQuery());
  const { data: lines } = useQuery(
    listLogLines.useQuery({
      start: Timestamp.fromDate(subHours(now, 12)),
    }),
  );
  const createNewTag = useMutation(createTag.useMutation({}));
  const createLog = useMutation(createLogLine.useMutation({}));

  return (
    <Stack>
      {tags && (
        <LogLineCreator
          tags={tags.tags}
          onCreate={async (tagName, logLinePayload) => {
            let tagId = logLinePayload.tagId;

            if (!tagId) {
              const newTag = await createNewTag.mutateAsync({
                name: tagName,
              });

              tagId = newTag.tag?.tagId;
            }

            await createLog.mutateAsync({
              ...logLinePayload,
              tagId,
            });
            await queryClient.invalidateQueries({
              queryKey: listTags.getQueryKey(),
            });
            await queryClient.invalidateQueries({
              queryKey: listLogLines.getQueryKey(),
            });
          }}
        />
      )}
      {lines && <LogLineTable lines={lines.logLines} />}
    </Stack>
  );
}
