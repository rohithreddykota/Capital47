import React, {
  useState,
  useEffect,
  useCallback,
  useContext,
  useRef,
} from 'react';
import { useTranslation } from 'react-i18next';
import { ThemeContext } from "styled-components";


import { Country } from 'store/modules/dashboard/types';

import { CircleSpinner } from 'react-spinners-kit';
import { FiChevronDown, FiSearch } from 'react-icons/fi';

import {
  Container,
  Loader,
  Current,
  List,
  Button,
  Flag,
  Search,
  Selected,
} from './SelectField.styled';
import { Scrollbar } from 'smooth-scrollbar-react';

interface Props {
  selectorName: string;
  data: Country[];
  currentData: Country;
  selectPlaceholder?: string;
  onClickCountry: (selectorName: string, item: Country) => void;
}

const SelectField: React.FC<Props> = ({
  selectorName,
  data = [],
  currentData = {},
  selectPlaceholder,
  onClickCountry,
}) => {
  const { t } = useTranslation();
  const [open, setOpen] = useState(false);
  const [isSearch, setIsSearch] = useState(false);
  const [searchCountries, setSearchCountries] = useState<Country[]>([]);
  const _themeContext = useContext(ThemeContext);
  const selectRef = useRef(null);

  const handleClickOutside = useCallback((event: MouseEvent) => {
    if (selectRef.current) {
      if (!(selectRef.current! as any).contains(event.target)) {
        setOpen(false);
        setIsSearch(false);
        setSearchCountries([]);
      }
    }
  }, []);

  const handleClickSelect = useCallback(() => {
    setOpen(true);
  }, []);

  const handleInputFilter = useCallback(
    (value: string) => {
      if (value.length > 0) {
        setIsSearch(true);

        const searchCountry = data.filter((country: Country) =>
          country.label.toLowerCase().includes(value.toLowerCase()),
        );

        return setSearchCountries(searchCountry);
      }

      setSearchCountries([]);
      setIsSearch(false);
    },
    [data],
  );

  useEffect(() => {
    if (open) {
      document.addEventListener('click', handleClickOutside, true);
    }

    return () => {
      document.removeEventListener('click', handleClickOutside, true);
    };
  }, [open, handleClickOutside]);

  return (
    <Container ref={selectRef} isOpen={open}>
      <Current
        isActive={open}
        onClick={handleClickSelect}
        disabled={data.length === 0}
      >

          <>
            {open ? (
              <Search>
                <input
                  type="text"
                  placeholder={t('common:selects.country.placeholder')}
                  onChange={event => handleInputFilter(event.target.value)}
                />
                <FiSearch />
              </Search>
            ) : (
              <Selected>
                {selectPlaceholder}
                {currentData.flag && <Flag source={currentData.flag} />}
                <span>{currentData.label}</span>
                <FiChevronDown size={18} />
              </Selected>
            )}
          </>

      </Current>

      <List isOpen={open}>
        <Scrollbar alwaysShowTracks continuousScrolling={false}>
          {!isSearch ? (
            <>
              {data.map((item: Country) => (
                <li key={item.id}>
                  <Button
                    onClick={() => {
                      onClickCountry(selectorName, item);
                      setOpen(false);
                    }}
                  >
                    {item.flag && <Flag source={item.flag} />}
                    {item.label}
                  </Button>
                </li>
              ))}
            </>
          ) : (
            <>
              {searchCountries.map((item: Country) => (
                <li key={item.id}>
                  <Button
                    onClick={() => {
                      onClickCountry(selectorName, item);
                      setOpen(false);
                      setIsSearch(false);
                    }}
                  >
                    {item.flag && <Flag source={item.flag} />}
                    {item.label}
                  </Button>
                </li>
              ))}
            </>
          )}
        </Scrollbar>
      </List>
    </Container>
  );
};

export default SelectField;
