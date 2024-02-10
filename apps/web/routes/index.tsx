import React from "react";
import { HashRouter, Navigate, Route, Routes } from "react-router-dom";

import { HomePage } from "./home-page";

const AppRoutes = () => {
  return (
    <HashRouter>
      <Routes>
        <Route path="/" element={<Navigate to="/home" />} />
        <Route path="/home" element={<HomePage />} />
      </Routes>
    </HashRouter>
  );
};

export { AppRoutes };
