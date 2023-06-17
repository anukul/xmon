import { create } from "zustand";

interface UIStore {
  loadingCount: number;
  increaseLoadingCount: () => void;
  decreaseLoadingCount: () => void;
}

export const useUIStore = create<UIStore>((set) => ({
  loadingCount: 0,
  increaseLoadingCount: () =>
    set((state) => ({ loadingCount: state.loadingCount + 1 })),
  decreaseLoadingCount: () =>
    set((state) => ({ loadingCount: state.loadingCount - 1 })),
}));
