"use client";

import { CircularProgress } from "@mui/material";

import { useListNodes } from "@/api/hooks";
import { BlockTable } from "@/app/partials/BlockTable";
import { NodeTable } from "@/app/partials/NodeTable";
import { SlotTable } from "@/app/partials/SlotTable";

export default function Home() {
  const { isLoading } = useListNodes();

  if (isLoading) {
    return (
      <div className="h-screen w-screen grid place-items-center">
        <CircularProgress />
      </div>
    );
  }

  return (
    <div className="p-6">
      <div className="text-xl mb-4">Nodes</div>
      <NodeTable />
      <div className="text-xl mt-8 mb-4">Blocks</div>
      <BlockTable />
      <div className="text-xl mt-8 mb-4">Slots</div>
      <SlotTable />
      <div className="text-xl mt-8 mb-4">Bad Blocks</div>
      <SlotTable />
    </div>
  );
}
