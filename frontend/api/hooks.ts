import { useEffect, useState } from "react";
import useSWR from "swr";

import { refreshInterval } from "@/api/client";
import { listBlocks, listNodes } from "@/api/fetchers";
import { RequestKeys } from "@/api/keys";
import {
  ListBlocksResponse_BlockData,
  ListNodesResponse_Node,
} from "@/proto/monitor_pb";

export const useListNodes = () => {
  const [response, setResponse] = useState<Array<ListNodesResponse_Node>>([]);
  const { data, error, isLoading } = useSWR(RequestKeys.LIST_NODES, listNodes, {
    refreshInterval: refreshInterval,
  });

  useEffect(() => {
    setResponse(data?.nodes || []);
  }, [data, isLoading]);

  return { data: response, error, isLoading };
};

export const useListBlocks = () => {
  const [response, setResponse] = useState<Array<ListBlocksResponse_BlockData>>(
    []
  );
  const { data, error, isLoading } = useSWR(
    RequestKeys.LIST_BLOCKS,
    listBlocks,
    {
      refreshInterval: refreshInterval,
    }
  );

  useEffect(() => {
    setResponse(data?.blocks || []);
  }, [data, isLoading]);

  return { data: response, error, isLoading };
};
