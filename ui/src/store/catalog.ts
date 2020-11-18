import { Instance, types } from 'mobx-state-tree';
import { IconName } from '../icons';

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
    },
    icon() {
      if (self.type === 'official') return IconName.cat;
      else if (self.type === 'verified') return IconName.certificate;
      else if (self.type === 'community') return IconName.user;
    }
  }));

export type ICatalog = Instance<typeof Catalog>;
export type ICatalogStore = Instance<typeof CatalogStore>;

export const CatalogStore = types
  .model({
    items: types.map(Catalog)
  })

  .actions((self) => ({
    add(item: ICatalog) {
      self.items.put({ id: item.id, name: item.name, type: item.type });
    },
    clearSelected() {
      self.items.forEach((c) => {
        c.selected = false;
      });
    }
  }))

  .views((self) => ({
    get values() {
      return Array.from(self.items.values());
    },

    get selected() {
      const list = new Set();
      self.items.forEach((c: ICatalog) => {
        if (c.selected) {
          list.add(c.id);
        }
      });

      return list;
    }
  }));
