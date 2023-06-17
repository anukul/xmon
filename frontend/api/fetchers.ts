import { ListBlocksRequest, ListNodesRequest } from "@/proto/monitor_pb";

import { api } from "./client";

export const listNodes = async () => {
  const request = new ListNodesRequest();
  return api.listNodes(request);
};

export const listBlocks = async () => {
  const request = new ListBlocksRequest();
  return api.listBlocks(request);
};
