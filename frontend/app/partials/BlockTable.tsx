import {
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TablePagination,
  TableRow,
} from "@mui/material";
import clsx from "clsx";
import { useMemo, useState } from "react";

import { useListBlocks, useListNodes } from "@/api/hooks";
import { useMuiTheme } from "@/app/use-mui-theme";

// TODO: use twin.macro
const stickyClassNames = "sticky left-0 drop-shadow-lg";

export const BlockTable = () => {
  const theme = useMuiTheme();
  const { data: nodes } = useListNodes();
  const { data: blocks } = useListBlocks();
  const [page, setPage] = useState(0);
  const [rowsPerPage, setRowsPerPage] = useState(10);

  const handleChangePage = (event: unknown, newPage: number) => {
    setPage(newPage);
  };

  const handleChangeRowsPerPage = (
    event: React.ChangeEvent<HTMLInputElement>
  ) => {
    setRowsPerPage(parseInt(event.target.value, 10));
    setPage(0);
  };

  const visibleBlocks = useMemo(
    () => blocks.slice(page * rowsPerPage, page * rowsPerPage + rowsPerPage),
    [blocks, page, rowsPerPage]
  );

  const redBackground =
    theme.palette.mode === "dark" ? "bg-red-700" : "bg-red-200";

  const greenBackground =
    theme.palette.mode === "dark" ? "bg-green-700" : "bg-green-200";

  return (
    <TableContainer component={Paper}>
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
              Block
            </TableCell>
            {nodes.map((node) => (
              <TableCell key={node.name} className="underline">
                {node.name}
              </TableCell>
            ))}
          </TableRow>
        </TableHead>
        <TableBody>
          {visibleBlocks.map((b) => (
            <TableRow
              key={Number(b.blockNumber)}
              className={clsx(b.isSynced ? greenBackground : redBackground)}
            >
              <TableCell
                className={clsx(stickyClassNames, "bg-inherit font-semibold")}
              >
                {Number(b.blockNumber)}
              </TableCell>
              {nodes.map((node) => (
                <TableCell key={node.name}>
                  <a
                    className="text-inherit no-underline"
                    href={`https://etherscan.io/block/${b.hashes[node.name]}`}
                    target="_blank"
                  >
                    {b.hashes[node.name].substring(2, 2 + 6)}
                  </a>
                </TableCell>
              ))}
            </TableRow>
          ))}
        </TableBody>
      </Table>
      <TablePagination
        component="div"
        count={blocks.length}
        onPageChange={handleChangePage}
        onRowsPerPageChange={handleChangeRowsPerPage}
        page={page}
        rowsPerPage={rowsPerPage}
        showFirstButton
        showLastButton
      />
    </TableContainer>
  );
};
