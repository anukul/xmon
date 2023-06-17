import {
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
} from "@mui/material";
import clsx from "clsx";

import { useMuiTheme } from "@/app/use-mui-theme";

const stickyClassNames = "sticky left-0 drop-shadow-lg";

export const SlotTable = () => {
  const theme = useMuiTheme();

  return (
    <TableContainer className="blur-sm" component={Paper}>
      <Table>
        <TableHead>
          <TableRow>
            <TableCell
              className={clsx(stickyClassNames)}
              style={{
                // TODO: use twind, twin-macro etc after TS 5.0 / Next support
                backgroundColor: theme.palette.background.default,
              }}
            >
              Slot
            </TableCell>
            <TableCell>Node A</TableCell>
            <TableCell>Node B</TableCell>
            <TableCell>Node C</TableCell>
          </TableRow>
        </TableHead>
        <TableBody>
          <TableRow>
            <TableCell className={clsx(stickyClassNames)}>12345</TableCell>
            <TableCell>0x12345</TableCell>
            <TableCell>0x12345</TableCell>
            <TableCell>0x12345</TableCell>
          </TableRow>
          <TableRow>
            <TableCell className={clsx(stickyClassNames)}>98765</TableCell>
            <TableCell>0x12345</TableCell>
            <TableCell>0x12345</TableCell>
            <TableCell>0x12345</TableCell>
          </TableRow>
        </TableBody>
      </Table>
    </TableContainer>
  );
};
