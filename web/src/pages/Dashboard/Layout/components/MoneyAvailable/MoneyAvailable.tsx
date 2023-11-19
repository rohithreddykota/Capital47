import React, { ReactNode } from 'react';

import { Container } from './MoneyAvailable.styled';

interface MoneyAvailableProps {
  children: ReactNode;
}

const MoneyAvailable: React.FC<MoneyAvailableProps> = ({ children }) => {
  return <Container>{children}</Container>;
};

export default MoneyAvailable;
