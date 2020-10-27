import React from 'react';
import { useObserver } from 'mobx-react';

import { Button, Checkbox, Text, TextVariants, Grid, GridItem } from '@patternfly/react-core';
import { IconSize, TimesIcon } from '@patternfly/react-icons';

import Icon from '../Icon/Icon';
import { IconName } from '../../icons';
import { titleCase } from '../../utils/titlecase';

import './Filter.css';

interface Filterable {
  id: number;
  name: string;
  selected: boolean;
  toggle(): void;
  icon(): IconName;
}

interface Store {
  list: Filterable[];
  clear(): void;
}

type iconMapper = (name: string) => IconName;

interface FilterList {
  store: Store;
  header: string;
  iconForFilter?: iconMapper;
}

const labelWithIcon = (label: string, iconFn?: iconMapper) => (
  <Grid>
    {iconFn && (
      <GridItem span={2}>
        <Icon id={iconFn(label)} size={IconSize.sm} label={label} />
      </GridItem>
    )}
    <GridItem span={iconFn ? 10 : 12}>{titleCase(label)}</GridItem>
  </Grid>
);

const checkboxes = (items: Filterable[], iconForFilter?: iconMapper) =>
  items.map((c: Filterable) => (
    <Checkbox
      key={c.id}
      label={labelWithIcon(c.name, iconForFilter)}
      isChecked={c.selected}
      onChange={() => c.toggle()}
      aria-label="controlled checkbox"
      id={String(c.id)}
      name={c.name}
    />
  ));

const Filter: React.FC<FilterList> = ({ store, header, iconForFilter }) => {
  return useObserver(() => (
    <div className="Filter">
      <Grid sm={6} md={4} lg={3} xl2={1}>
        <GridItem className="hub-filter-header" span={1} rowSpan={1}>
          <Text component={TextVariants.h1} style={{ fontWeight: 'bold' }}>
            {header}
          </Text>
        </GridItem>

        <GridItem rowSpan={1}>
          <Button variant="plain" aria-label="Clear" onClick={store.clear}>
            <TimesIcon />
          </Button>
        </GridItem>
      </Grid>

      <Grid>
        <GridItem className="hub-filter-checkboxes">
          {checkboxes(store.list, iconForFilter)}
        </GridItem>
      </Grid>
    </div>
  ));
};

export default Filter;
