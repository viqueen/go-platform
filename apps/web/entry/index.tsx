import React, { useMemo } from "react";

import { createTheme, CssBaseline, ThemeProvider } from "@mui/material";
import { createRoot, hydrateRoot } from "react-dom/client";

import { AppRoutes } from "../routes";

const App = () => {
  const theme = useMemo(
    () =>
      createTheme({
        palette: {
          mode: "light",
        },
      }),
    [],
  );

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <AppRoutes />
    </ThemeProvider>
  );
};

const rootContainer = document.querySelector("#root");
if (rootContainer) {
  if (rootContainer.hasChildNodes()) {
    hydrateRoot(rootContainer, <App />);
  } else {
    createRoot(rootContainer).render(<App />);
  }
}
