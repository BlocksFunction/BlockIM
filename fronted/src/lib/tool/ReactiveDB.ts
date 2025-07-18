import { cloneDeep, debounce, type DebouncedFunc, isEqual } from "lodash-es";
import { type Ref, ref, type ShallowRef, shallowRef, watch } from "vue";

interface IndexedDBConfig {
  dbName: string;
  version?: number;
  storeName: string;
  keyPath?: string;
  autoIncrement?: boolean;
  useShallowRef?: boolean;
  autoSaveDelay?: number;
}

class ReactiveDB<T extends Record<string, unknown>> {
  private db: IDBDatabase | null = null;
  private entries = new Map<string | number, Ref<T> | ShallowRef<T>>();
  private config: Required<IndexedDBConfig>;
  private autoSaveTimers = new Map<
    string | number,
    DebouncedFunc<() => void>
  >();

  constructor(config: IndexedDBConfig) {
    this.config = {
      version: 1,
      keyPath: "id",
      autoIncrement: true,
      useShallowRef: false,
      autoSaveDelay: 500,
      ...config,
    };
  }

  initialize(): Promise<void> {
    if (!globalThis.indexedDB) {
      throw new Error("该浏览器不支持IndexedDB, 请更换或升级浏览器");
    }

    return new Promise((resolve, reject) => {
      const request = indexedDB.open(
        this.config.dbName,
        this.config.version,
      );

      request.onupgradeneeded = (event) => {
        const db = (event.target as IDBOpenDBRequest).result;
        if (!db.objectStoreNames.contains(this.config.storeName)) {
          db.createObjectStore(this.config.storeName, {
            keyPath: this.config.keyPath,
            autoIncrement: this.config.autoIncrement,
          });
        }
      };

      request.onsuccess = (event) => {
        this.db = (event.target as IDBOpenDBRequest).result;
        resolve();
      };

      request.onerror = (event) => {
        reject((event.target as IDBOpenDBRequest).error);
      };
    });
  }

  private getObjectStore(mode: IDBTransactionMode = "readonly") {
    if (!this.db) throw new Error("数据库没有初始化");
    const transaction = this.db.transaction(this.config.storeName, mode);
    return transaction.objectStore(this.config.storeName);
  }

  private createReactiveEntry(data: T): Ref<T> | ShallowRef<T> {
    const key = data[this.config.keyPath] as string | number;
    const entry = this.config.useShallowRef
      ? shallowRef<T>(data)
      : ref<T>(data);

    if (this.config.autoSaveDelay > 0) {
      const saveHandler = debounce(() => {
        this.save(key);
      }, this.config.autoSaveDelay);

      this.autoSaveTimers.set(key, saveHandler);

      watch(
        entry,
        (newVal, oldVal) => {
          if (!isEqual(newVal, oldVal)) {
            saveHandler();
          }
        },
        { deep: !this.config.useShallowRef },
      );
    }

    return entry as Ref<T> | ShallowRef<T>;
  }

  add(item: T): Promise<T> {
    return new Promise((resolve, reject) => {
      const store = this.getObjectStore("readwrite");
      const request = store.add(item);

      request.onsuccess = (event) => {
        const key = (event.target as IDBRequest).result;
        const newItem = { ...item, [this.config.keyPath]: key } as T;
        const entry = this.createReactiveEntry(newItem);
        this.entries.set(key, entry);
        resolve(newItem);
      };

      request.onerror = (event) => {
        reject((event.target as IDBRequest).error);
      };
    });
  }

  get(id: string | number): Promise<Ref<T> | ShallowRef<T>> {
    if (this.entries.has(id)) {
      return Promise.resolve(this.entries.get(id)!);
    }

    return new Promise((resolve, reject) => {
      const store = this.getObjectStore();
      const request = store.get(id);

      request.onsuccess = (event) => {
        const data = (event.target as IDBRequest).result as
          | T
          | undefined;
        if (!data) {
          reject(new Error("未找到实体"));
          return;
        }

        const entry = this.createReactiveEntry(data);
        this.entries.set(id, entry);
        resolve(entry);
      };

      request.onerror = (event) => {
        reject((event.target as IDBRequest).error);
      };
    });
  }

  getAll(): Promise<Array<Ref<T> | ShallowRef<T>>> {
    return new Promise((resolve, reject) => {
      const store = this.getObjectStore();
      const request = store.getAll();

      request.onsuccess = (event) => {
        const data = (event.target as IDBRequest).result as
          | T[]
          | undefined;
        if (!data) {
          resolve([]);
          return;
        }

        const entries: Array<Ref<T> | ShallowRef<T>> = [];
        for (const item of data) {
          const key = item[this.config.keyPath] as string | number;
          let entry: Ref<T> | ShallowRef<T>;
          if (this.entries.has(key)) {
            entry = this.entries.get(key)!;
          } else {
            entry = this.createReactiveEntry(item);
            this.entries.set(key, entry);
          }
          entries.push(entry);
        }
        resolve(entries);
      };

      request.onerror = (event) => {
        reject((event.target as IDBRequest).error);
      };
    });
  }

  save(id: string | number): Promise<void> {
    if (!this.entries.has(id)) {
      throw new Error("未找到实体");
    }

    const entry = this.entries.get(id)!;
    const data = cloneDeep(entry.value);

    return new Promise((resolve, reject) => {
      const store = this.getObjectStore("readwrite");
      const request = store.put(data);

      request.onsuccess = () => resolve();
      request.onerror = (event) => {
        reject((event.target as IDBRequest).error);
      };
    });
  }

  delete(id: string | number): Promise<void> {
    if (this.entries.has(id)) {
      this.entries.delete(id);
      const timer = this.autoSaveTimers.get(id);
      if (timer) {
        timer.cancel?.();
        this.autoSaveTimers.delete(id);
      }
    }

    return new Promise((resolve, reject) => {
      const store = this.getObjectStore("readwrite");
      const request = store.delete(id);

      request.onsuccess = () => resolve();
      request.onerror = (event) => {
        reject((event.target as IDBRequest).error);
      };
    });
  }

  clearAll(): Promise<void> {
    this.entries.clear();
    this.autoSaveTimers.forEach((timer) => timer.cancel?.());
    this.autoSaveTimers.clear();

    return new Promise((resolve, reject) => {
      const store = this.getObjectStore("readwrite");
      const request = store.clear();

      request.onsuccess = () => resolve();
      request.onerror = (event) => {
        reject((event.target as IDBRequest).error);
      };
    });
  }
}

export default ReactiveDB;
