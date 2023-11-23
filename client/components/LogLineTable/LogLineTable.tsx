import { Table } from '@mantine/core';
import { format } from 'date-fns';

import { LogLine } from '../../generated/proto/logger/v1/logger_pb';
import { TagDisplay } from '../TagDisplay';
import { autoWidthTd, table } from './LogLineTable.css';

type Props = {
  lines: LogLine[];
};

export function LogLineTable({ lines }: Props) {
  return (
    <Table classNames={{ table }}>
      <Table.Tbody>
        {lines.map((line) => (
          <Table.Tr key={line.logLineId}>
            <Table.Td className={autoWidthTd}>
              {line.tag && <TagDisplay tag={line.tag} />}
            </Table.Td>
            <Table.Td className={autoWidthTd}>
              {line.createdAt
                ? format(line.createdAt.toDate(), 'MM-dd-yyyy hh:mm:ss aa')
                : '-'}
            </Table.Td>
            <Table.Td className={autoWidthTd}>{line.value}</Table.Td>
            <Table.Td>{line.description}</Table.Td>
          </Table.Tr>
        ))}
      </Table.Tbody>
    </Table>
  );
}
