import React from 'react';
import { Routes as _Routes, Route } from 'react-router-dom';

import Dashboard from '../pages/Dashboard/DashboardScreen';

const Routes: React.FC = () => (
  <_Routes>
    <Route path="/" element={<Dashboard/>} />
  </_Routes>
);

export default Routes;
