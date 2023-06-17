import { createPromiseClient, Interceptor } from "@bufbuild/connect";
import { createConnectTransport } from "@bufbuild/connect-web";

import { useUIStore } from "@/app/store";
import { Monitor } from "@/proto/monitor_connect";

const showProgressBar: Interceptor = (next: Function) => async (req) => {
  const { increaseLoadingCount, decreaseLoadingCount } = useUIStore.getState();
  increaseLoadingCount();
  const response = await next(req);
  decreaseLoadingCount();
  return response;
};

const transport = createConnectTransport({
  baseUrl: process.env.NEXT_PUBLIC_API_URL!,
  interceptors: [showProgressBar],
});

export const api = createPromiseClient(Monitor, transport);

export const refreshInterval = 5 * 1000;
