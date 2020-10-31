import { types, Instance } from 'mobx-state-tree';

export const Kind = types
  .model({
    name: types.identifier,
    selected: false
  })
  .actions((self) => ({
    toggle() {
      self.selected = !self.selected;
    }
  }));

export type IKind = Instance<typeof Kind>;
export type IKindStore = Instance<typeof KindStore>;

export const KindStore = types
  .model({
    kinds: types.map(Kind)
  })

  .actions((self) => ({
    add(item: string) {
      self.kinds.put({ name: item, selected: false });
    },

    clearSelected() {
      self.kinds.forEach((k) => {
        k.selected = false;
      });
    }
  }))

  .views((self) => ({
    get items() {
      return Array.from(self.kinds.values());
    },

    get selected() {
      const list = new Set();
      self.kinds.forEach((c: IKind) => {
        if (c.selected) {
          list.add(c.name);
        }
      });

      return list;
    }
  }));
