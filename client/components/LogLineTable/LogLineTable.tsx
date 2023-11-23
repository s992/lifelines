import { Table } from '@mantine/core';
import { format } from 'date-fns';

import { LogLine, Tag } from '../../generated/proto/lifelines/v1/lifelines_pb';
import { TagDisplay } from '../TagDisplay';
import { autoWidthTd, table } from './LogLineTable.css';

type Props = {
  lines: LogLine[];
  onTagClick?: (tag: Tag) => void;
};

export function LogLineTable({ lines, onTagClick }: Props) {
  return (
    <Table classNames={{ table }}>
      <Table.Tbody>
        {lines.map((line) => (
          <Table.Tr key={line.logLineId}>
            <Table.Td className={autoWidthTd}>
              {line.tag && <TagDisplay tag={line.tag} onClick={onTagClick} />}
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
