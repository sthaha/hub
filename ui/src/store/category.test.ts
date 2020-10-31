import { when } from 'mobx';
import { getSnapshot } from 'mobx-state-tree';
import { FakeHub } from '../api/testutil';
import { CategoryStore, Category, Tag } from './category';

const TESTDATA_DIR = `${__dirname}/testdata`;
const api = new FakeHub(TESTDATA_DIR);

describe('Store Object', () => {
  it('can create a tag object', () => {
    const store = Tag.create({
      id: 1,
      name: 'cli'
    });
    expect(store.name).toBe('cli');
  });
  it('can create a category object', () => {
    const category = Category.create({
      id: 1,
      name: 'test',
      tags: ['1']
    });

    expect(category.name).toBe('test');
    expect(category.tags.length).toBe(1);
  });
});

describe('Store functions', () => {
  it('can create a category store', (done) => {
    const store = CategoryStore.create({}, { api });
    expect(store.count).toBe(0);
    expect(store.isLoading).toBe(true);
    when(
      () => !store.isLoading,
      () => {
        expect(store.count).toBe(5);
        expect(store.isLoading).toBe(false);

        expect(getSnapshot(store)).toMatchSnapshot();

        done();
      }
    );
  });

  it('can toggle the selected category', (done) => {
    const store = CategoryStore.create({}, { api });
    expect(store.count).toBe(0);
    expect(store.isLoading).toBe(true);

    when(
      () => !store.isLoading,
      () => {
        expect(store.count).toBe(5);
        expect(store.isLoading).toBe(false);

        store.categories.get('1')?.toggle();

        expect(store.categories.get('1')?.selected).toBe(true);

        done();
      }
    );
  });

  it('can clear all the categories', (done) => {
    const store = CategoryStore.create({}, { api });
    expect(store.count).toBe(0);

    when(
      () => !store.isLoading,
      () => {
        expect(store.count).toBe(5);
        expect(store.isLoading).toBe(false);

        store.categories.get('1')?.toggle();
        store.categories.get('2')?.toggle();

        store.clearSelected();
        expect(store.categories.get('1')?.selected).toBe(false);

        done();
      }
    );
  });

  it('can return the tags for the categories which are selected', (done) => {
    const store = CategoryStore.create({}, { api });
    expect(store.count).toBe(0);
    expect(store.isLoading).toBe(true);

    when(
      () => !store.isLoading,
      () => {
        expect(store.count).toBe(5);
        expect(store.isLoading).toBe(false);

        store.categories.get('1')?.toggle();
        store.categories.get('2')?.toggle();

        expect(store.selectedTags.size).toBe(2);

        done();
      }
    );
  });
});
