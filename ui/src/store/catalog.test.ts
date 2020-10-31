import { Catalog, CatalogStore } from './catalog';
import { getSnapshot } from 'mobx-state-tree';

describe('Store Object', () => {
  it('can create a catalog object', () => {
    const store = Catalog.create({
      id: 1,
      name: 'tekton',
      type: 'community'
    });

    expect(store.name).toBe('tekton');
  });

  it('creates a catalog store', (done) => {
    const store = CatalogStore.create({});

    const item = Catalog.create({
      id: 1,
      name: 'tekton',
      type: 'community'
    });

    store.add(item);

    expect(getSnapshot(store.catalogs)).toMatchSnapshot();

    done();
  });

  it('should toggle a selected catalog', (done) => {
    const store = CatalogStore.create({});

    const item = Catalog.create({
      id: 1,
      name: 'tekton',
      type: 'community'
    });

    store.add(item);
    store.catalogs.get('1')?.toggle();

    expect(store.selected.size).toBe(1);
    expect(store.catalogs.get('1')?.selected).toBe(true);

    done();
  });

  it('should clear all the selected catalog', (done) => {
    const store = CatalogStore.create({});

    const item = Catalog.create({
      id: 1,
      name: 'tekton',
      type: 'community'
    });

    store.add(item);
    store.catalogs.get('1')?.toggle();

    store.clearSelected();

    expect(store.catalogs.get('1')?.selected).toBe(false);

    done();
  });
});
