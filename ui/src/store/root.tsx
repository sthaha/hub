import React, { ReactChild, ReactChildren } from 'react';
import { types, getEnv, Instance } from 'mobx-state-tree';
import { CategoryStore } from './category';
import { ResourceStore } from './resource';
import { Hub, Api } from '../api';

export const Root = types.model('Root', {}).views((self) => ({
  get api(): Api {
    return getEnv(self).api;
  },
  get categories() {
    return getEnv(self).categories;
  },
  get resources() {
    return getEnv(self).resources;
  }
}));

export type IRoot = Instance<typeof Root>;

let RootContext: React.Context<IRoot>;
export const useMst = () => React.useContext(RootContext);

export const initRootStore = (api: Api) => {
  const categories = CategoryStore.create({}, { api });
  const resources = ResourceStore.create({ catalogs: {}, kinds: {} }, { api });
  return Root.create({}, { api, categories, resources });
};

interface Props {
  children: ReactChild | ReactChildren;
}

export const createProviderAndStore = (api?: Api) => {
  const root = initRootStore(api || new Hub());
  RootContext = React.createContext(root);

  const Provider = ({ children }: Props) => (
    <RootContext.Provider value={root}> {children} </RootContext.Provider>
  );
  return { Provider, root };
};

export const createProvider = (api?: Api) => {
  const { Provider } = createProviderAndStore(api);
  return Provider;
};
