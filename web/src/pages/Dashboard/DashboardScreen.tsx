import React, { useCallback, useEffect } from 'react';
import moment from 'moment';
import axios from 'axios';

import { useDispatch, useSelector } from 'react-redux';
import { Country, DashboardState } from '../../store/modules/dashboard/types';
import {
  countriesRequest,
  updateFromCountry,
  updateToCountry,
  updateYouSend,
  updateDeliveryDate,
  endAnimations,
  reverseCurrency,
} from '../../store/modules/dashboard/actions';

import Dashboard from './Layout/Dashboard.layout';

const DashboardScreen: React.FC = () => {
  const dispatch = useDispatch();

  const { youSend, recipientGets, fromCountry, toCountry, typeDelivery } =
    useSelector((state: { dashboard: DashboardState }) => state.dashboard);

  useEffect(() => {
    setTimeout(() => {
      dispatch(countriesRequest());
      dispatch(endAnimations());
    }, 1500);
  }, [dispatch]);

  useEffect(() => {
    dispatch(updateYouSend(youSend));
  }, [dispatch, fromCountry, toCountry, youSend]);

  const handleClickCountry = useCallback(
    (selectorName: string, country: Country) => {
      if (selectorName === 'from') {
        return dispatch(updateFromCountry(country));
      }

      dispatch(updateToCountry(country));
    },
    [dispatch]
  );

  const handleChangeYouSend = useCallback(
    (value: string | number) => {
      dispatch(updateYouSend(Number(value)));
    },
    [dispatch]
  );

  const handleUpdateDateCalendar = useCallback(
    (date: string, delivery: string) => {
      dispatch(updateDeliveryDate(date, delivery));
    },
    [dispatch]
  );

  const handleClickReverseCurrency = useCallback(
    (from: Country, to: Country) => {
      dispatch(reverseCurrency(from, to));
    },
    [dispatch]
  );

  const handleSubmitConfirm = useCallback(() => {
    const payload = {
      medium: 'balance',
      payee_id: '65599a9f9683f20dd5188a43',
      transaction_date: '2023-11-19',
      status: 'pending',
      description: 'Sending money to my friend',
      amount: youSend,
    };
    axios
      .post(
        'http://api.nessieisreal.com/accounts/65599bf59683f20dd5188a44/transfers?key=e157f3ddfb74e929d3bb40ed199ad3ba',
        payload
      )
      .then((response) => {
        alert(`response.data.message. Transfered: ${youSend}`);
      })
      .catch((error) => {
        console.error('Error:', error);
      });
  }, [youSend, recipientGets, fromCountry, toCountry, typeDelivery]);

  return (
    <Dashboard
      onClickCountry={handleClickCountry}
      onChangeYouSend={handleChangeYouSend}
      onReverseCurrency={handleClickReverseCurrency}
      onUpdateDateCalendar={handleUpdateDateCalendar}
      onSubmitConfirm={handleSubmitConfirm}
    />
  );
};

export default DashboardScreen;
