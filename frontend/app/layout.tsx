"use client";

import "./globals.css";

import createCache from "@emotion/cache";
import { CacheProvider } from "@emotion/react";
import {
  CssBaseline,
  LinearProgress,
  StyledEngineProvider,
  ThemeProvider,
} from "@mui/material";
import { Inter } from "next/font/google";
import { ReactNode } from "react";

import { useUIStore } from "@/app/store";
import { useMuiTheme } from "@/app/use-mui-theme";

const inter = Inter({ subsets: ["latin"] });

const cache = createCache({
  key: "css",
  prepend: true,
});

export default function RootLayout({ children }: { children: ReactNode }) {
  const theme = useMuiTheme();
  const { loadingCount } = useUIStore();

  return (
    <html lang="en">
      <body className={inter.className} id="root">
        <CacheProvider value={cache}>
          <StyledEngineProvider injectFirst>
            <ThemeProvider theme={theme}>
              <CssBaseline />
              {loadingCount > 0 && (
                <LinearProgress className="fixed top-0 w-full" />
              )}
              {children}
            </ThemeProvider>
          </StyledEngineProvider>
        </CacheProvider>
      </body>
    </html>
  );
}
