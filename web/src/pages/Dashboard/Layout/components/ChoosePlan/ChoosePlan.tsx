import React, { HTMLAttributes, useContext } from 'react';
import { ThemeContext } from 'styled-components';
import { useTranslation } from 'react-i18next';
import moment from 'moment';

import { useSelector } from 'react-redux';
import { DashboardState } from 'store/modules/dashboard/types';

import { DatePicker } from 'antd';

import { FiCalendar, FiCheck } from 'react-icons/fi';

import {
  Container,
  ContainerItem,
  Header,
  List,
  Content,
  Title,
  Description,
  Price,
  Checkmark,
  ButtonDate,
} from './ChoosePlan.styled';

interface Props extends HTMLAttributes<HTMLDivElement> {
  onUpdateDateCalendar: (date: string, typeDelivery: string) => void;
}

const ChoosePlan: React.FC<Props> = ({ onUpdateDateCalendar, ...rest }) => {
  const { t } = useTranslation();

  const themeContext = useContext(ThemeContext);

  const formatExpress = '[Get] D MMMM YYYY [till 12pm]';
  const formatStandard = '[Get] D MMMM YYYY [till 6pm]';

  const { delivery, typeDelivery } = useSelector(
    (state: { dashboard: DashboardState }) => state.dashboard,
  );

  return (
    <Container {...rest}>
      <Header>
        <h2>{t('dashboard:chooseAPlan')}</h2>

        <ButtonDate>
          {t('common:selects.date.placeholder')}
          <FiCalendar size={18} color={themeContext?.colors.onSecondary} />
        </ButtonDate>

        <DatePicker
          placeholder={t('common:selects.date.placeholder')}
          suffixIcon={<FiCalendar size={18} color={themeContext?.colors.onSecondary} />}
          // defaultValue={moment()}
          size="large"
          inputReadOnly
          bordered={false}
          showNow={false}
          allowClear={false}
          showToday={false}
          onChange={(_, dateString) =>
            onUpdateDateCalendar(dateString, typeDelivery)
          }
        />
      </Header>

      <List>
        <form>
          <li>
            <label htmlFor="opt-express">
              <input
                type="radio"
                id="opt-express"
                name="plan"
                value={moment(delivery).set('hour', 12).toString()}
                onChange={event =>
                  onUpdateDateCalendar(event.target.value, 'Express')
                }
                defaultChecked
              />

              <ContainerItem>
                <Checkmark>
                  <FiCheck />
                </Checkmark>

                <div>
                  <Content>
                    <Title>{moment(delivery).format(formatExpress)}</Title>
                    <Description>{t('dashboard:express')}</Description>
                  </Content>

                  <Price>$0.99</Price>
                </div>
              </ContainerItem>
            </label>
          </li>

          <li>
            <label htmlFor="opt-standard">
              <input
                type="radio"
                id="opt-standard"
                name="plan"
                value={moment(delivery).set('hour', 6).toString()}
                onChange={event =>
                  onUpdateDateCalendar(event.target.value, 'Standard')
                }
              />

              <ContainerItem>
                <Checkmark>
                  <FiCheck />
                </Checkmark>

                <div>
                  <Content>
                    <Title>{moment(delivery).format(formatStandard)}</Title>
                    <Description>{t('dashboard:standard')}</Description>
                  </Content>

                  <Price>$1.00</Price>
                </div>
              </ContainerItem>
            </label>
          </li>

          <li>
            <label htmlFor="opt-today">
              <input
                type="radio"
                id="opt-today"
                name="plan"
                value={moment(delivery).set('hour', 8).toString()}
                onChange={event =>
                  onUpdateDateCalendar(event.target.value, 'Today')
                }
              />

              <ContainerItem>
                <Checkmark>
                  <FiCheck />
                </Checkmark>

                <div>
                  <Content>
                    <Title>{t('dashboard:getToday')}</Title>
                    <Description>{t('dashboard:onlyOnWorking')}</Description>
                  </Content>

                  <Price>$1.00</Price>
                </div>
              </ContainerItem>
            </label>
          </li>
        </form>
      </List>
    </Container>
  );
};

export default ChoosePlan;
