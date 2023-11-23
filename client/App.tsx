import { Timestamp } from '@bufbuild/protobuf';
import { Center, Stack } from '@mantine/core';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { subHours } from 'date-fns';
import { useMemo, useState } from 'react';

import { container } from './App.css';
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
import { ListLogLinesRequest } from './generated/proto/logger/v1/logger_pb';

export function App() {
  const queryClient = useQueryClient();
  const now = useMemo(() => new Date(), []);
  const [filter, setFilter] = useState<Partial<ListLogLinesRequest>>({
    start: Timestamp.fromDate(subHours(now, 12)),
  });
  const { data: tags } = useQuery(listTags.useQuery());
  const { data: lines } = useQuery(listLogLines.useQuery(filter));
  const createNewTag = useMutation(createTag.useMutation({}));
  const createLog = useMutation(createLogLine.useMutation({}));

  return (
    <Center>
      <Stack className={container}>
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
        {lines && (
          <LogLineTable
            lines={lines.logLines}
            onTagClick={(tag) => {
              if (filter.tagId) {
                setFilter({
                  ...filter,
                  tagId: undefined,
                });
              } else {
                setFilter({
                  ...filter,
                  tagId: tag.tagId,
                });
              }
            }}
          />
        )}
      </Stack>
    </Center>
  );
}
