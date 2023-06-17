import {
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from "@mui/material";
import _ from "lodash";
import { DateTime } from "luxon";
import { Fragment } from "react";

import { useListNodes } from "@/api/hooks";
import { ListNodesResponse_Client_Status } from "@/proto/monitor_pb";

export const NodeTable = () => {
  const { data: nodes } = useListNodes();

  function getRelativeTime(from: bigint): string {
    return DateTime.fromSeconds(Number(from)).toRelative()!;
  }

  return (
    <TableContainer component={Paper}>
      <Table>
        <TableHead>
          <TableRow>
            <TableCell>Name</TableCell>
            <TableCell>Version</TableCell>
            <TableCell>Status</TableCell>
            <TableCell>Last Progress</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          {nodes.map((node) => (
            <Fragment key={node.name}>
              <TableRow>
                <TableCell className="underline" colSpan={100}>
                  {node.name}
                </TableCell>
              </TableRow>
              <TableRow>
                <TableCell className="pl-10">Execution Client</TableCell>
                <TableCell>{node.executionClient?.version}</TableCell>
                <TableCell>
                  {_.startCase(
                    _.lowerCase(
                      ListNodesResponse_Client_Status[
                        node.executionClient!.status
                      ]
                    )
                  )}
                </TableCell>
                <TableCell>
                  {getRelativeTime(node.executionClient!.lastPing!.seconds)}
                </TableCell>
              </TableRow>
              <TableRow>
                <TableCell className="pl-10">Consensus Client</TableCell>
                <TableCell className="blur-sm">
                  besu/v23.4.1/linux-x86_64/openjdk-java-17
                </TableCell>
                <TableCell className="blur-sm">Healthy</TableCell>
                <TableCell className="blur-sm">10 minutes ago</TableCell>
              </TableRow>
            </Fragment>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
};
