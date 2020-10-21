import React from 'react';
import { useObserver } from 'mobx-react';

import { Button, Checkbox, Text, TextVariants, Grid, GridItem } from '@patternfly/react-core';
import { IconSize, TimesIcon } from '@patternfly/react-icons';

import Icon from '../Icon/Icon';
import { titleCase } from '../../utils/titlecase';

import './Filter.css';
import { IconName } from '../../icons';

interface Filterable {
  id: number;
  name: string;
  selected: boolean;
  toggle(): void;
  icon(): IconName;
}

interface Store {
  values: Filterable[];
  clearSelected(): void;
}

export type iconMapper = (name: string) => IconName;

interface Props {
  store: Store;
  header: string;
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

const checkboxes = (values: Filterable[]) => {
  return values.map((c: Filterable) => (
    <Checkbox
      key={String(c.id)}
      label={labelWithIcon(c.name, c.icon)}
      isChecked={c.selected}
      onChange={() => c.toggle()}
      aria-label="controlled checkbox"
      id={String(c.id)}
      name={c.name}
    />
  ));
};

const Filter: React.FC<Props> = ({ store, header }) => {
  return useObserver(() => (
    <div className="hub-filter">
      <Grid sm={6} md={4} lg={3} xl2={1}>
        <GridItem className="hub-filter-header" span={1} rowSpan={1}>
          <Text component={TextVariants.h1} style={{ fontWeight: 'bold' }}>
            {header}
          </Text>
        </GridItem>

        <GridItem rowSpan={2}>
          <Button variant="plain" aria-label="Clear" onClick={store.clearSelected}>
            <TimesIcon />
          </Button>
        </GridItem>
      </Grid>

      <Grid>
        <GridItem className="hub-filter-checkboxes">{checkboxes(store.values)}</GridItem>
      </Grid>
    </div>
  ));
};

export default Filter;
