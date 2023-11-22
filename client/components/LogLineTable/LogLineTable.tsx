import { Table, TableTbody, TableTd, TableTr } from '@mantine/core';
import { format } from 'date-fns';

import { LogLine } from '../../generated/proto/logger/v1/logger_pb';
import { table } from './LogLineTable.css';

type Props = {
  lines: LogLine[];
};

export function LogLineTable({ lines }: Props) {
  return (
    <Table classNames={{ table }}>
      <TableTbody>
        {lines.map((line) => (
          <TableTr key={line.logLineId}>
            <TableTd>#{line.tag?.name}</TableTd>
            <TableTd>
              {line.createdAt
                ? format(line.createdAt.toDate(), 'MM-dd-yyyy hh:mm:ss aa')
                : '-'}
            </TableTd>
            <TableTd>{line.value}</TableTd>
            <TableTd>{line.description}</TableTd>
          </TableTr>
        ))}
      </TableTbody>
    </Table>
  );
}
