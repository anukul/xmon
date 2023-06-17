import { createTheme, useMediaQuery } from "@mui/material";
import { useMemo } from "react";

export const useMuiTheme = () => {
  const prefersDarkMode = useMediaQuery("(prefers-color-scheme: dark)");

  return useMemo(
    () =>
      createTheme({
        palette: {
          mode: prefersDarkMode ? "dark" : "light",
        },
      }),
    [prefersDarkMode]
  );
};
