import { Instance, types } from 'mobx-state-tree';

export const Catalog = types
  .model({
    id: types.identifierNumber,
    name: types.optional(types.string, ''),
    type: types.optional(types.string, ''),
    selected: false
  })
  .actions((self) => ({
    toggle() {
      self.selected = !self.selected;
    }
  }));

export type ICatalog = Instance<typeof Catalog>;
export type ICatalogStore = Instance<typeof CatalogStore>;

export const CatalogStore = types
  .model({
    catalogs: types.map(Catalog)
  })

  .actions((self) => ({
    add(item: ICatalog) {
      self.catalogs.put({ id: item.id, name: item.name, type: item.type });
    },
    clearSelected() {
      self.catalogs.forEach((c) => {
        c.selected = false;
      });
    }
  }))

  .views((self) => ({
    get items() {
      return Array.from(self.catalogs.values());
    },

    get selected() {
      const list = new Set();
      self.catalogs.forEach((c: ICatalog) => {
        if (c.selected) {
          list.add(c.id);
        }
      });

      return list;
    }
  }));
