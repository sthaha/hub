import { ResourceStore, Resource } from './resource';
import { getSnapshot } from 'mobx-state-tree';
import { when } from 'mobx';
import { FakeHub } from '../api/testutil';
import { CategoryStore } from './category';

const TESTDATA_DIR = `${__dirname}/testdata`;
const api = new FakeHub(TESTDATA_DIR);

describe('Store Object', () => {
  it('can create a resource object', () => {
    const store = Resource.create({
      id: 5,
      name: 'buildah',
      catalog: '1',
      kind: 'Task',
      latestVersion: 1,
      tags: ['1'],
      rating: 5
    });

    expect(store.name).toBe('buildah');
  });
});

describe('Store functions', () => {
  it('creates a resource store', (done) => {
    const store = ResourceStore.create(
      {
        catalogs: {},
        kinds: {}
      },
      {
        api,
        categoryStore: CategoryStore.create({}, { api })
      }
    );

    expect(store.isLoading).toBe(true);

    when(
      () => !store.isLoading,
      () => {
        expect(store.isLoading).toBe(false);

        expect(store.resources.size).toBe(6);

        expect(getSnapshot(store.resources)).toMatchSnapshot();

        done();
      }
    );
  });

  it('creates a catalog store', (done) => {
    const store = ResourceStore.create(
      {
        catalogs: {},
        kinds: {}
      },
      {
        api,
        categoryStore: CategoryStore.create({}, { api })
      }
    );
    expect(store.isLoading).toBe(true);

    when(
      () => !store.isLoading,
      () => {
        expect(store.isLoading).toBe(false);

        expect(store.resources.size).toBe(6);

        expect(getSnapshot(store.catalogs)).toMatchSnapshot();

        done();
      }
    );
  });

  it('creates a kind store', (done) => {
    const store = ResourceStore.create(
      {
        catalogs: {},
        kinds: {}
      },
      {
        api,
        categoryStore: CategoryStore.create({}, { api })
      }
    );

    expect(store.isLoading).toBe(true);

    when(
      () => !store.isLoading,
      () => {
        expect(store.isLoading).toBe(false);
        expect(store.resources.size).toBe(6);

        expect(getSnapshot(store.kinds)).toMatchSnapshot();

        done();
      }
    );
  });

  it('filter resources based on selected catalog', (done) => {
    const store = ResourceStore.create(
      {
        catalogs: {},
        kinds: {}
      },
      {
        api,
        categoryStore: CategoryStore.create({}, { api })
      }
    );
    expect(store.isLoading).toBe(true);

    when(
      () => !store.isLoading,
      () => {
        expect(store.isLoading).toBe(false);
        expect(store.resources.size).toBe(6);

        store.catalogs.catalogs.get('2')?.toggle();

        expect(store.filteredResources.length).toBe(1);
        expect(store.filteredResources[0].name).toBe('hub');
        expect(store.filteredResources[0].catalog.name).toBe('tekton-hub');

        done();
      }
    );
  });

  it('filter resources based on selected kind', (done) => {
    const store = ResourceStore.create(
      {
        catalogs: {},
        kinds: {}
      },
      {
        api,
        categoryStore: CategoryStore.create({}, { api })
      }
    );
    expect(store.isLoading).toBe(true);

    when(
      () => !store.isLoading,
      () => {
        expect(store.isLoading).toBe(false);
        expect(store.resources.size).toBe(6);

        store.kinds.kinds.get('Pipeline')?.toggle();

        expect(store.filteredResources.length).toBe(1);
        expect(store.filteredResources[0].name).toBe('hub');
        expect(store.filteredResources[0].kind.name).toBe('Pipeline');

        done();
      }
    );
  });

  it('filter resources based on selected kind and catalog', (done) => {
    const store = ResourceStore.create(
      {
        catalogs: {},
        kinds: {}
      },
      {
        api,
        categoryStore: CategoryStore.create({}, { api })
      }
    );
    expect(store.isLoading).toBe(true);

    when(
      () => !store.isLoading,
      () => {
        expect(store.isLoading).toBe(false);
        expect(store.resources.size).toBe(6);

        store.catalogs.catalogs.get('1')?.toggle();
        store.kinds.kinds.get('Task')?.toggle();
        store.setSearch('golang');

        store.filteredResources;

        expect(store.filteredResources.length).toBe(1);

        done();
      }
    );
  });

  it('makes sure to not add duplicate resources', (done) => {
    const store = ResourceStore.create(
      {
        catalogs: {},
        kinds: {}
      },
      {
        api,
        categoryStore: CategoryStore.create({}, { api })
      }
    );

    expect(store.isLoading).toBe(true);

    when(
      () => !store.isLoading,
      () => {
        expect(store.isLoading).toBe(false);

        const item = Resource.create({
          id: 44,
          name: 'golang-build',
          catalog: 1,
          kind: 'Task',
          latestVersion: 47,
          tags: [1],
          rating: 5,
          versions: [47],
          displayName: 'golang build'
        });

        store.add(item);
        expect(store.resources.size).toBe(6);

        expect(getSnapshot(store.resources)).toMatchSnapshot();

        done();
      }
    );
  });

  it('it checks if the related date is a string', (done) => {
    const store = ResourceStore.create(
      {
        catalogs: {},
        kinds: {}
      },
      {
        api,
        categoryStore: CategoryStore.create({}, { api })
      }
    );
    expect(store.isLoading).toBe(true);
    when(
      () => !store.isLoading,
      () => {
        expect(store.isLoading).toBe(false);
        expect(store.resources.size).toBe(6);

        expect(typeof store.versions.get('1')?.updatedAt.fromNow()).toBe('string');

        done();
      }
    );
  });
});
