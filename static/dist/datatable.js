"use strict";
var DataTable = (() => {
  var __defProp = Object.defineProperty;
  var __getOwnPropDesc = Object.getOwnPropertyDescriptor;
  var __getOwnPropNames = Object.getOwnPropertyNames;
  var __hasOwnProp = Object.prototype.hasOwnProperty;
  var __export = (target, all) => {
    for (var name in all)
      __defProp(target, name, { get: all[name], enumerable: true });
  };
  var __copyProps = (to, from, except, desc) => {
    if (from && typeof from === "object" || typeof from === "function") {
      for (let key of __getOwnPropNames(from))
        if (!__hasOwnProp.call(to, key) && key !== except)
          __defProp(to, key, { get: () => from[key], enumerable: !(desc = __getOwnPropDesc(from, key)) || desc.enumerable });
    }
    return to;
  };
  var __toCommonJS = (mod) => __copyProps(__defProp({}, "__esModule", { value: true }), mod);

  // js/datatable/src/index.ts
  var src_exports = {};
  __export(src_exports, {
    DataTable: () => DataTable,
    EventEmitter: () => EventEmitter,
    ExportPlugin: () => ExportPlugin,
    FilterEngine: () => FilterEngine,
    SortEngine: () => SortEngine,
    StateManager: () => StateManager,
    VERSION: () => VERSION,
    VirtualScrollPlugin: () => VirtualScrollPlugin,
    calculateOptimalOverscan: () => calculateOptimalOverscan,
    createDataTable: () => createDataTable,
    createExportPlugin: () => createExportPlugin,
    createVirtualScrollPlugin: () => createVirtualScrollPlugin,
    debounce: () => debounce,
    default: () => src_default,
    measureRowHeight: () => measureRowHeight
  });

  // js/datatable/src/utils/event-emitter.ts
  var EventEmitter = class {
    constructor() {
      this.events = /* @__PURE__ */ new Map();
    }
    on(event, callback) {
      if (!this.events.has(event)) {
        this.events.set(event, /* @__PURE__ */ new Set());
      }
      this.events.get(event).add(callback);
      return () => this.off(event, callback);
    }
    off(event, callback) {
      const callbacks = this.events.get(event);
      if (callbacks) {
        callbacks.delete(callback);
      }
    }
    emit(event, data) {
      const callbacks = this.events.get(event);
      if (callbacks) {
        callbacks.forEach((callback) => {
          try {
            callback(data);
          } catch (error) {
            console.error(`EventEmitter: Error in ${String(event)} handler:`, error);
          }
        });
      }
    }
    clear() {
      this.events.clear();
    }
    removeAllListeners(event) {
      if (event) {
        this.events.delete(event);
      } else {
        this.events.clear();
      }
    }
  };

  // js/datatable/src/utils/state-manager.ts
  var StateManager = class {
    constructor() {
      this.storage = null;
      try {
        if (typeof window !== "undefined" && window.localStorage) {
          this.storage = window.localStorage;
        }
      } catch (e) {
        console.warn("StateManager: localStorage is not available");
      }
    }
    save(key, state) {
      if (!this.storage)
        return;
      try {
        const serialized = JSON.stringify(state, this.replacer);
        this.storage.setItem(key, serialized);
      } catch (error) {
        console.error("StateManager: Error saving state:", error);
      }
    }
    load(key) {
      if (!this.storage)
        return null;
      try {
        const serialized = this.storage.getItem(key);
        if (!serialized)
          return null;
        return JSON.parse(serialized, this.reviver);
      } catch (error) {
        console.error("StateManager: Error loading state:", error);
        return null;
      }
    }
    remove(key) {
      if (!this.storage)
        return;
      try {
        this.storage.removeItem(key);
      } catch (error) {
        console.error("StateManager: Error removing state:", error);
      }
    }
    clear() {
      if (!this.storage)
        return;
      try {
        this.storage.clear();
      } catch (error) {
        console.error("StateManager: Error clearing state:", error);
      }
    }
    // Custom replacer to handle Sets, Maps, etc.
    replacer(_key, value) {
      if (value instanceof Set) {
        return {
          __type: "Set",
          __value: Array.from(value)
        };
      }
      if (value instanceof Map) {
        return {
          __type: "Map",
          __value: Array.from(value.entries())
        };
      }
      if (value instanceof Date) {
        return {
          __type: "Date",
          __value: value.toISOString()
        };
      }
      return value;
    }
    // Custom reviver to restore Sets, Maps, etc.
    reviver(_key, value) {
      if (value && typeof value === "object" && value.__type) {
        switch (value.__type) {
          case "Set":
            return new Set(value.__value);
          case "Map":
            return new Map(value.__value);
          case "Date":
            return new Date(value.__value);
          default:
            return value;
        }
      }
      return value;
    }
  };

  // js/datatable/src/utils/filter-engine.ts
  var FilterEngine = class {
    constructor(options) {
      this.options = options;
    }
    /**
     * Apply all filters to rows
     */
    filter(rows, filterState, columns) {
      let filteredRows = rows;
      if (filterState.filters.length > 0) {
        filteredRows = this.applyColumnFilters(filteredRows, filterState.filters, columns);
      }
      if (filterState.globalSearch && filterState.globalSearch.trim()) {
        filteredRows = this.applyGlobalSearch(filteredRows, filterState.globalSearch, columns);
      }
      return filteredRows;
    }
    /**
     * Apply individual column filters
     */
    applyColumnFilters(rows, filters, columns) {
      return rows.filter((row) => {
        return filters.every((filter) => this.evaluateFilter(row, filter, columns));
      });
    }
    /**
     * Apply global search across all searchable columns
     */
    applyGlobalSearch(rows, query, columns) {
      const searchableColumns = columns.filter((col) => col.filterable !== false);
      if (searchableColumns.length === 0) {
        return rows;
      }
      const normalizedQuery = this.options.filtering.caseSensitive ? query : query.toLowerCase();
      return rows.filter((row) => {
        return searchableColumns.some((column) => {
          const value = this.getCellValue(row, column);
          const normalizedValue = this.normalizeValue(value);
          return normalizedValue.includes(normalizedQuery);
        });
      });
    }
    /**
     * Evaluate a single filter against a row
     */
    evaluateFilter(row, filter, columns) {
      if (this.options.filtering.customFilters?.[filter.columnId]) {
        return this.options.filtering.customFilters[filter.columnId](row, filter);
      }
      const column = columns.find((col) => col.id === filter.columnId);
      if (!column)
        return true;
      const cellValue = this.getCellValue(row, column);
      return this.applyOperator(cellValue, filter.operator, filter.value);
    }
    /**
     * Apply filter operator
     */
    applyOperator(cellValue, operator, filterValue) {
      const normalizedCellValue = this.normalizeValue(cellValue);
      const normalizedFilterValue = Array.isArray(filterValue) ? filterValue.map((v) => this.normalizeValue(v)) : this.normalizeValue(filterValue);
      switch (operator) {
        case "equals":
          return normalizedCellValue === normalizedFilterValue;
        case "notEquals":
          return normalizedCellValue !== normalizedFilterValue;
        case "contains":
          return normalizedCellValue.includes(normalizedFilterValue);
        case "notContains":
          return !normalizedCellValue.includes(normalizedFilterValue);
        case "startsWith":
          return normalizedCellValue.startsWith(normalizedFilterValue);
        case "endsWith":
          return normalizedCellValue.endsWith(normalizedFilterValue);
        case "isEmpty":
          return cellValue === null || cellValue === void 0 || cellValue === "";
        case "isNotEmpty":
          return cellValue !== null && cellValue !== void 0 && cellValue !== "";
        case "greaterThan":
          return this.compareNumeric(cellValue, filterValue) > 0;
        case "greaterThanOrEqual":
          return this.compareNumeric(cellValue, filterValue) >= 0;
        case "lessThan":
          return this.compareNumeric(cellValue, filterValue) < 0;
        case "lessThanOrEqual":
          return this.compareNumeric(cellValue, filterValue) <= 0;
        case "between":
          if (!Array.isArray(filterValue) || filterValue.length !== 2)
            return false;
          const num = this.toNumber(cellValue);
          const min = this.toNumber(filterValue[0]);
          const max = this.toNumber(filterValue[1]);
          return num >= min && num <= max;
        case "in":
          if (!Array.isArray(filterValue))
            return false;
          return normalizedFilterValue.includes(normalizedCellValue);
        case "notIn":
          if (!Array.isArray(filterValue))
            return false;
          return !normalizedFilterValue.includes(normalizedCellValue);
        default:
          console.warn(`FilterEngine: Unknown operator ${operator}`);
          return true;
      }
    }
    /**
     * Get cell value from row using column definition
     */
    getCellValue(row, column) {
      if (column.accessor) {
        return column.accessor(row.data);
      }
      return row.data[column.field];
    }
    /**
     * Normalize value for comparison
     */
    normalizeValue(value) {
      if (value === null || value === void 0) {
        return "";
      }
      const stringValue = String(value);
      return this.options.filtering.caseSensitive ? stringValue : stringValue.toLowerCase();
    }
    /**
     * Convert value to number for numeric comparisons
     */
    toNumber(value) {
      if (typeof value === "number")
        return value;
      if (typeof value === "string") {
        const num = parseFloat(value);
        return isNaN(num) ? 0 : num;
      }
      return 0;
    }
    /**
     * Compare numeric values
     */
    compareNumeric(a, b) {
      return this.toNumber(a) - this.toNumber(b);
    }
  };
  function debounce(func, wait) {
    let timeout = null;
    return function(...args) {
      if (timeout) {
        clearTimeout(timeout);
      }
      timeout = setTimeout(() => {
        func(...args);
      }, wait);
    };
  }

  // js/datatable/src/utils/sort-engine.ts
  var SortEngine = class {
    constructor(options) {
      /**
       * Default comparison logic
       */
      this.defaultComparator = (a, b) => {
        if (a === null || a === void 0)
          return b === null || b === void 0 ? 0 : -1;
        if (b === null || b === void 0)
          return 1;
        if (typeof a === "number" && typeof b === "number") {
          return a - b;
        }
        if (typeof a === "boolean" && typeof b === "boolean") {
          return a === b ? 0 : a ? 1 : -1;
        }
        if (this.isDate(a) && this.isDate(b)) {
          return a.getTime() - b.getTime();
        }
        const aNum = this.tryParseNumber(a);
        const bNum = this.tryParseNumber(b);
        if (aNum !== null && bNum !== null) {
          return aNum - bNum;
        }
        const aDate = this.tryParseDate(a);
        const bDate = this.tryParseDate(b);
        if (this.isDate(aDate) && this.isDate(bDate)) {
          return aDate.getTime() - bDate.getTime();
        }
        return String(a).localeCompare(String(b), void 0, {
          numeric: true,
          sensitivity: "base"
        });
      };
      this.options = options;
    }
    /**
     * Sort rows based on sort state
     */
    sort(rows, sortState, columns) {
      const sortedRows = [...rows];
      if ("sorts" in sortState) {
        return this.multiSort(sortedRows, sortState.sorts, columns);
      } else {
        return this.singleSort(sortedRows, sortState, columns);
      }
    }
    /**
     * Single column sort
     */
    singleSort(rows, sortState, columns) {
      const column = columns.find((col) => col.id === sortState.columnId);
      if (!column || !sortState.direction)
        return rows;
      const comparator = this.getComparator(column);
      const multiplier = sortState.direction === "asc" ? 1 : -1;
      return rows.sort((a, b) => {
        const aValue = this.getCellValue(a, column);
        const bValue = this.getCellValue(b, column);
        return comparator(aValue, bValue) * multiplier;
      });
    }
    /**
     * Multi-column sort
     */
    multiSort(rows, sorts, columns) {
      return rows.sort((a, b) => {
        for (const sort of sorts) {
          if (!sort.direction)
            continue;
          const column = columns.find((col) => col.id === sort.columnId);
          if (!column)
            continue;
          const comparator = this.getComparator(column);
          const multiplier = sort.direction === "asc" ? 1 : -1;
          const aValue = this.getCellValue(a, column);
          const bValue = this.getCellValue(b, column);
          const result = comparator(aValue, bValue) * multiplier;
          if (result !== 0)
            return result;
        }
        return 0;
      });
    }
    /**
     * Get comparator function for column
     */
    getComparator(column) {
      if (column.comparator) {
        return column.comparator;
      }
      if (this.options.customComparators[column.id]) {
        return this.options.customComparators[column.id];
      }
      return this.defaultComparator;
    }
    /**
     * Get cell value from row using column definition
     */
    getCellValue(row, column) {
      if (column.accessor) {
        return column.accessor(row.data);
      }
      return row.data[column.field];
    }
    /**
     * Try to parse value as number
     */
    tryParseNumber(value) {
      if (typeof value === "number")
        return value;
      if (typeof value === "string") {
        if (/^\d{4}-\d{2}-\d{2}/.test(value) || /^\d{1,2}\/\d{1,2}\/\d{4}/.test(value) || /^\d{1,2}-\d{1,2}-\d{4}/.test(value)) {
          return null;
        }
        const cleaned = value.replace(/[,$]/g, "");
        const num = parseFloat(cleaned);
        if (isNaN(num) || cleaned !== String(num)) {
          return null;
        }
        return num;
      }
      return null;
    }
    /**
     * Try to parse value as date
     */
    tryParseDate(value) {
      if (this.isDate(value))
        return value;
      if (typeof value === "string" || typeof value === "number") {
        if (typeof value === "string") {
          if (/^\d{4}-\d{2}-\d{2}/.test(value)) {
            const date2 = new Date(value);
            return isNaN(date2.getTime()) ? null : date2;
          }
          if (/^\d{1,2}\/\d{1,2}\/\d{4}/.test(value) || /^\d{1,2}-\d{1,2}-\d{4}/.test(value)) {
            const date2 = new Date(value);
            return isNaN(date2.getTime()) ? null : date2;
          }
        }
        const date = new Date(value);
        return isNaN(date.getTime()) ? null : date;
      }
      return null;
    }
    /**
     * Check if value is a Date
     */
    isDate(value) {
      return value instanceof Date;
    }
  };

  // js/datatable/src/core.ts
  var DataTable = class {
    constructor(options) {
      this.plugins = /* @__PURE__ */ new Map();
      this.initialized = false;
      this.destroyed = false;
      this.options = this.normalizeOptions(options);
      this.eventEmitter = new EventEmitter();
      this.stateManager = new StateManager();
      this.filterEngine = new FilterEngine(this.options);
      this.sortEngine = new SortEngine(this.options);
      this.state = this.createInitialState();
      this.initialize();
    }
    // ============================================================================
    // Initialization
    // ============================================================================
    normalizeOptions(options) {
      return {
        data: options.data || [],
        columns: options.columns,
        dataMode: options.dataMode || "client",
        serverSide: {
          url: "",
          method: "GET",
          headers: {},
          params: {},
          dataPath: "data",
          totalPath: "total",
          transformer: (response) => ({
            data: response.data || response,
            total: response.total || response.data?.length || 0
          }),
          ...options.serverSide
        },
        sorting: {
          enabled: true,
          multiSort: false,
          defaultSort: void 0,
          ...options.sorting
        },
        filtering: {
          enabled: true,
          globalSearch: true,
          debounceMs: 300,
          caseSensitive: false,
          customFilters: {},
          ...options.filtering
        },
        pagination: {
          enabled: true,
          pageSize: 10,
          pageSizeOptions: [5, 10, 25, 50, 100],
          showFirstLast: true,
          ...options.pagination
        },
        selection: {
          enabled: false,
          multiple: true,
          mode: "checkbox",
          selectOnRowClick: false,
          ...options.selection
        },
        virtualScroll: {
          enabled: false,
          rowHeight: 48,
          overscan: 5,
          ...options.virtualScroll
        },
        rowId: options.rowId || "id",
        enableRowExpansion: options.enableRowExpansion || false,
        preserveState: options.preserveState || false,
        stateKey: options.stateKey || "datatable-state",
        customComparators: options.customComparators || {},
        customFormatters: options.customFormatters || {},
        onInit: options.onInit || (() => {
        }),
        onDataChange: options.onDataChange || (() => {
        }),
        onError: options.onError || (() => {
        })
      };
    }
    createInitialState() {
      const columns = this.options.columns.map((col) => ({
        visible: true,
        sortable: true,
        filterable: true,
        ...col
      }));
      return {
        rows: [],
        columns,
        sortState: this.options.sorting.defaultSort || void 0,
        filterState: {
          filters: [],
          globalSearch: ""
        },
        paginationState: {
          pageIndex: 0,
          pageSize: this.options.pagination.pageSize,
          totalRows: 0,
          totalPages: 0
        },
        selectionState: {
          selectedRowIds: /* @__PURE__ */ new Set(),
          isAllSelected: false,
          isPartiallySelected: false
        },
        loadState: {
          loading: false,
          error: null
        },
        visibleRows: [],
        filteredRows: [],
        sortedRows: [],
        paginatedRows: [],
        visibleColumns: columns.filter((col) => col.visible !== false)
      };
    }
    initialize() {
      if (this.initialized)
        return;
      if (this.options.preserveState) {
        this.restoreState();
      }
      if (this.options.dataMode === "client" && this.options.data.length > 0) {
        this.setData(this.options.data);
      } else if (this.options.dataMode === "server") {
        this.reload();
      }
      this.initialized = true;
      this.emit("init");
      this.options.onInit();
    }
    // ============================================================================
    // Data Management
    // ============================================================================
    setData(data) {
      this.ensureNotDestroyed();
      const rows = data.map((item) => ({
        id: this.getRowId(item),
        data: item,
        selected: false,
        expanded: false,
        disabled: false
      }));
      this.state.rows = rows;
      this.state.paginationState.totalRows = rows.length;
      this.recompute();
      this.emit("data:change", rows);
      this.emit("data:load", rows);
      this.options.onDataChange(rows);
    }
    getData() {
      return [...this.state.rows];
    }
    addRow(data) {
      this.ensureNotDestroyed();
      const row = {
        id: this.getRowId(data),
        data,
        selected: false,
        expanded: false,
        disabled: false
      };
      this.state.rows.push(row);
      this.state.paginationState.totalRows = this.state.rows.length;
      this.recompute();
      this.emit("data:change", this.state.rows);
    }
    updateRow(id, data) {
      this.ensureNotDestroyed();
      const row = this.state.rows.find((r) => r.id === id);
      if (!row) {
        console.warn(`DataTable: Row with id ${id} not found`);
        return;
      }
      row.data = { ...row.data, ...data };
      this.recompute();
      this.emit("data:change", this.state.rows);
    }
    deleteRow(id) {
      this.ensureNotDestroyed();
      const index = this.state.rows.findIndex((r) => r.id === id);
      if (index === -1) {
        console.warn(`DataTable: Row with id ${id} not found`);
        return;
      }
      this.state.rows.splice(index, 1);
      this.state.paginationState.totalRows = this.state.rows.length;
      this.state.selectionState.selectedRowIds.delete(id);
      this.recompute();
      this.emit("data:change", this.state.rows);
    }
    clearData() {
      this.ensureNotDestroyed();
      this.state.rows = [];
      this.state.paginationState.totalRows = 0;
      this.state.selectionState.selectedRowIds.clear();
      this.recompute();
      this.emit("data:change", []);
    }
    // ============================================================================
    // Columns
    // ============================================================================
    getColumns() {
      return [...this.state.columns];
    }
    setColumns(columns) {
      this.ensureNotDestroyed();
      this.state.columns = columns.map((col) => ({
        visible: true,
        sortable: true,
        filterable: true,
        ...col
      }));
      this.state.visibleColumns = this.state.columns.filter((col) => col.visible !== false);
      this.emit("state:change", this.state);
    }
    showColumn(columnId) {
      this.toggleColumnVisibility(columnId, true);
    }
    hideColumn(columnId) {
      this.toggleColumnVisibility(columnId, false);
    }
    toggleColumn(columnId) {
      const column = this.state.columns.find((col) => col.id === columnId);
      if (column) {
        this.toggleColumnVisibility(columnId, !column.visible);
      }
    }
    toggleColumnVisibility(columnId, visible) {
      this.ensureNotDestroyed();
      const column = this.state.columns.find((col) => col.id === columnId);
      if (!column) {
        console.warn(`DataTable: Column ${columnId} not found`);
        return;
      }
      column.visible = visible;
      this.state.visibleColumns = this.state.columns.filter((col) => col.visible !== false);
      this.emit("column:visibility", { columnId, visible });
      this.emit("state:change", this.state);
    }
    // ============================================================================
    // Sorting
    // ============================================================================
    sort(columnId, direction) {
      this.ensureNotDestroyed();
      if (!this.options.sorting.enabled) {
        console.warn("DataTable: Sorting is disabled");
        return;
      }
      const column = this.state.columns.find((col) => col.id === columnId);
      if (!column || column.sortable === false) {
        console.warn(`DataTable: Column ${columnId} is not sortable`);
        return;
      }
      let newDirection = direction ?? "asc";
      if (!direction) {
        const currentSort = this.state.sortState;
        if (currentSort?.columnId === columnId) {
          if (currentSort.direction === "asc") {
            newDirection = "desc";
          } else if (currentSort.direction === "desc") {
            newDirection = null;
          }
        }
      }
      this.state.sortState = newDirection ? { columnId, direction: newDirection } : void 0;
      if (this.options.dataMode === "client") {
        this.recompute();
      } else {
        this.reload();
      }
      this.emit("sort:change", this.state.sortState);
      this.emit("state:change", this.state);
    }
    multiSort(sorts) {
      this.ensureNotDestroyed();
      if (!this.options.sorting.enabled || !this.options.sorting.multiSort) {
        console.warn("DataTable: Multi-sort is not enabled");
        return;
      }
      this.state.sortState = { sorts };
      if (this.options.dataMode === "client") {
        this.recompute();
      } else {
        this.reload();
      }
      this.emit("sort:change", this.state.sortState);
      this.emit("state:change", this.state);
    }
    clearSort() {
      this.state.sortState = void 0;
      if (this.options.dataMode === "client") {
        this.recompute();
      } else {
        this.reload();
      }
      this.emit("sort:change", this.state.sortState);
      this.emit("state:change", this.state);
    }
    getSortState() {
      return this.state.sortState;
    }
    // ============================================================================
    // Filtering
    // ============================================================================
    filter(filters) {
      this.ensureNotDestroyed();
      if (!this.options.filtering.enabled) {
        console.warn("DataTable: Filtering is disabled");
        return;
      }
      this.state.filterState.filters = filters;
      if (this.options.dataMode === "client") {
        this.recompute();
      } else {
        this.reload();
      }
      this.emit("filter:change", this.state.filterState);
      this.emit("state:change", this.state);
    }
    addFilter(filter) {
      this.ensureNotDestroyed();
      this.state.filterState.filters = this.state.filterState.filters.filter(
        (f) => f.columnId !== filter.columnId
      );
      this.state.filterState.filters.push(filter);
      if (this.options.dataMode === "client") {
        this.recompute();
      } else {
        this.reload();
      }
      this.emit("filter:change", this.state.filterState);
      this.emit("state:change", this.state);
    }
    removeFilter(columnId) {
      this.ensureNotDestroyed();
      this.state.filterState.filters = this.state.filterState.filters.filter(
        (f) => f.columnId !== columnId
      );
      if (this.options.dataMode === "client") {
        this.recompute();
      } else {
        this.reload();
      }
      this.emit("filter:change", this.state.filterState);
      this.emit("state:change", this.state);
    }
    clearFilters() {
      this.state.filterState.filters = [];
      this.state.filterState.globalSearch = "";
      if (this.options.dataMode === "client") {
        this.recompute();
      } else {
        this.reload();
      }
      this.emit("filter:change", this.state.filterState);
      this.emit("state:change", this.state);
    }
    search(query) {
      this.ensureNotDestroyed();
      if (!this.options.filtering.enabled || !this.options.filtering.globalSearch) {
        console.warn("DataTable: Global search is not enabled");
        return;
      }
      this.state.filterState.globalSearch = query;
      if (this.options.dataMode === "client") {
        this.recompute();
      } else {
        this.reload();
      }
      this.emit("search:change", query);
      this.emit("filter:change", this.state.filterState);
      this.emit("state:change", this.state);
    }
    getFilterState() {
      return { ...this.state.filterState };
    }
    // ============================================================================
    // Pagination
    // ============================================================================
    goToPage(page) {
      this.ensureNotDestroyed();
      if (!this.options.pagination.enabled) {
        console.warn("DataTable: Pagination is disabled");
        return;
      }
      const totalPages = this.state.paginationState.totalPages;
      const newPage = Math.max(0, Math.min(page, totalPages - 1));
      if (newPage === this.state.paginationState.pageIndex)
        return;
      this.state.paginationState.pageIndex = newPage;
      if (this.options.dataMode === "client") {
        this.recompute();
      } else {
        this.reload();
      }
      this.emit("page:change", this.state.paginationState);
      this.emit("state:change", this.state);
    }
    nextPage() {
      this.goToPage(this.state.paginationState.pageIndex + 1);
    }
    previousPage() {
      this.goToPage(this.state.paginationState.pageIndex - 1);
    }
    firstPage() {
      this.goToPage(0);
    }
    lastPage() {
      this.goToPage(this.state.paginationState.totalPages - 1);
    }
    setPageSize(size) {
      this.ensureNotDestroyed();
      if (!this.options.pagination.enabled) {
        console.warn("DataTable: Pagination is disabled");
        return;
      }
      this.state.paginationState.pageSize = size;
      this.state.paginationState.pageIndex = 0;
      if (this.options.dataMode === "client") {
        this.recompute();
      } else {
        this.reload();
      }
      this.emit("page:change", this.state.paginationState);
      this.emit("state:change", this.state);
    }
    getPaginationState() {
      return { ...this.state.paginationState };
    }
    // ============================================================================
    // Selection
    // ============================================================================
    selectRow(id) {
      this.ensureNotDestroyed();
      if (!this.options.selection.enabled) {
        console.warn("DataTable: Selection is disabled");
        return;
      }
      if (!this.options.selection.multiple) {
        this.state.selectionState.selectedRowIds.clear();
      }
      const row = this.state.rows.find((r) => r.id === id);
      if (!row || row.disabled)
        return;
      row.selected = true;
      this.state.selectionState.selectedRowIds.add(id);
      this.updateSelectionState();
      this.emit("row:select", { row, selected: true });
      this.emit("selection:change", this.state.selectionState);
      this.emit("state:change", this.state);
    }
    deselectRow(id) {
      this.ensureNotDestroyed();
      const row = this.state.rows.find((r) => r.id === id);
      if (!row)
        return;
      row.selected = false;
      this.state.selectionState.selectedRowIds.delete(id);
      this.updateSelectionState();
      this.emit("row:select", { row, selected: false });
      this.emit("selection:change", this.state.selectionState);
      this.emit("state:change", this.state);
    }
    toggleRowSelection(id) {
      if (this.state.selectionState.selectedRowIds.has(id)) {
        this.deselectRow(id);
      } else {
        this.selectRow(id);
      }
    }
    selectAll() {
      this.ensureNotDestroyed();
      if (!this.options.selection.enabled || !this.options.selection.multiple) {
        console.warn("DataTable: Multi-selection is not enabled");
        return;
      }
      this.state.visibleRows.forEach((row) => {
        if (!row.disabled) {
          row.selected = true;
          this.state.selectionState.selectedRowIds.add(row.id);
        }
      });
      this.updateSelectionState();
      this.emit("selection:change", this.state.selectionState);
      this.emit("state:change", this.state);
    }
    deselectAll() {
      this.ensureNotDestroyed();
      this.state.rows.forEach((row) => {
        row.selected = false;
      });
      this.state.selectionState.selectedRowIds.clear();
      this.updateSelectionState();
      this.emit("selection:change", this.state.selectionState);
      this.emit("state:change", this.state);
    }
    getSelectedRows() {
      return this.state.rows.filter(
        (row) => this.state.selectionState.selectedRowIds.has(row.id)
      );
    }
    getSelectionState() {
      return {
        selectedRowIds: new Set(this.state.selectionState.selectedRowIds),
        isAllSelected: this.state.selectionState.isAllSelected,
        isPartiallySelected: this.state.selectionState.isPartiallySelected
      };
    }
    updateSelectionState() {
      const selectableRows = this.state.visibleRows.filter((row) => !row.disabled);
      const selectedCount = this.state.selectionState.selectedRowIds.size;
      this.state.selectionState.isAllSelected = selectableRows.length > 0 && selectedCount === selectableRows.length;
      this.state.selectionState.isPartiallySelected = selectedCount > 0 && selectedCount < selectableRows.length;
    }
    // ============================================================================
    // State Management
    // ============================================================================
    getState() {
      return { ...this.state };
    }
    setState(newState) {
      this.ensureNotDestroyed();
      this.state = { ...this.state, ...newState };
      this.emit("state:change", this.state);
    }
    resetState() {
      this.state = this.createInitialState();
      this.clearFilters();
      this.clearSort();
      this.deselectAll();
      this.firstPage();
      this.emit("state:change", this.state);
    }
    saveState() {
      if (!this.options.preserveState)
        return;
      this.stateManager.save(this.options.stateKey, this.state);
    }
    restoreState() {
      const savedState = this.stateManager.load(this.options.stateKey);
      if (savedState) {
        this.state = { ...this.state, ...savedState };
      }
    }
    // ============================================================================
    // Server-side Operations
    // ============================================================================
    async reload() {
      if (this.options.dataMode !== "server") {
        console.warn("DataTable: Server-side mode is not enabled");
        return;
      }
      if (!this.options.serverSide.url) {
        console.error("DataTable: Server URL is not configured");
        return;
      }
      this.state.loadState.loading = true;
      this.state.loadState.error = null;
      this.emit("state:change", this.state);
      try {
        const params = this.buildServerParams();
        const url = new URL(this.options.serverSide.url);
        if (this.options.serverSide.method === "GET") {
          Object.entries(params).forEach(([key, value]) => {
            url.searchParams.append(key, String(value));
          });
        }
        const response = await fetch(url.toString(), {
          method: this.options.serverSide.method,
          headers: {
            "Content-Type": "application/json",
            ...this.options.serverSide.headers
          },
          body: this.options.serverSide.method === "POST" ? JSON.stringify(params) : void 0
        });
        if (!response.ok) {
          throw new Error(`HTTP ${response.status}: ${response.statusText}`);
        }
        const json = await response.json();
        const { data, total } = this.options.serverSide.transformer(json);
        this.setData(data);
        this.state.paginationState.totalRows = total;
        this.state.paginationState.totalPages = Math.ceil(total / this.state.paginationState.pageSize);
        this.state.loadState.lastFetch = /* @__PURE__ */ new Date();
        this.emit("data:load", this.state.rows);
      } catch (error) {
        this.state.loadState.error = error;
        this.emit("data:error", error);
        this.options.onError(error);
      } finally {
        this.state.loadState.loading = false;
        this.emit("state:change", this.state);
      }
    }
    buildServerParams() {
      const params = {
        ...this.options.serverSide.params
      };
      if (this.options.pagination.enabled) {
        params.page = this.state.paginationState.pageIndex;
        params.pageSize = this.state.paginationState.pageSize;
      }
      if (this.state.sortState) {
        if ("sorts" in this.state.sortState) {
          params.sort = this.state.sortState.sorts;
        } else {
          params.sortBy = this.state.sortState.columnId;
          params.sortDir = this.state.sortState.direction;
        }
      }
      if (this.state.filterState.filters.length > 0) {
        params.filters = this.state.filterState.filters;
      }
      if (this.state.filterState.globalSearch) {
        params.search = this.state.filterState.globalSearch;
      }
      return params;
    }
    // ============================================================================
    // Computation Pipeline
    // ============================================================================
    recompute() {
      let rows = [...this.state.rows];
      if (this.options.filtering.enabled) {
        rows = this.filterEngine.filter(rows, this.state.filterState, this.state.columns);
      }
      this.state.filteredRows = rows;
      if (this.options.sorting.enabled && this.state.sortState) {
        rows = this.sortEngine.sort(rows, this.state.sortState, this.state.columns);
      }
      this.state.sortedRows = rows;
      this.state.paginationState.totalRows = rows.length;
      this.state.paginationState.totalPages = Math.ceil(
        rows.length / this.state.paginationState.pageSize
      );
      if (this.options.pagination.enabled) {
        const start = this.state.paginationState.pageIndex * this.state.paginationState.pageSize;
        const end = start + this.state.paginationState.pageSize;
        rows = rows.slice(start, end);
      }
      this.state.paginatedRows = rows;
      this.state.visibleRows = rows;
      this.updateSelectionState();
      if (this.options.preserveState) {
        this.saveState();
      }
      this.emit("state:change", this.state);
    }
    // ============================================================================
    // Event System
    // ============================================================================
    on(event, callback) {
      return this.eventEmitter.on(event, callback);
    }
    off(event, callback) {
      this.eventEmitter.off(event, callback);
    }
    emit(event, data) {
      this.eventEmitter.emit(event, data);
    }
    // ============================================================================
    // Plugin System
    // ============================================================================
    use(plugin) {
      if (this.plugins.has(plugin.name)) {
        console.warn(`DataTable: Plugin ${plugin.name} is already installed`);
        return;
      }
      plugin.install(this);
      this.plugins.set(plugin.name, plugin);
    }
    // ============================================================================
    // Utilities
    // ============================================================================
    getRowId(row) {
      if (typeof this.options.rowId === "function") {
        return this.options.rowId(row);
      }
      return row[this.options.rowId];
    }
    ensureNotDestroyed() {
      if (this.destroyed) {
        throw new Error("DataTable: Instance has been destroyed");
      }
    }
    // ============================================================================
    // Lifecycle
    // ============================================================================
    destroy() {
      if (this.destroyed)
        return;
      this.emit("destroy");
      this.eventEmitter.clear();
      this.plugins.forEach((plugin) => {
        if (plugin.uninstall) {
          plugin.uninstall(this);
        }
      });
      this.plugins.clear();
      this.destroyed = true;
    }
  };
  function createDataTable(options) {
    return new DataTable(options);
  }

  // js/datatable/src/plugins/export.ts
  var ExportPlugin = class {
    constructor() {
      this.name = "export";
      this.version = "1.0.0";
      this.table = null;
    }
    install(table) {
      this.table = table;
    }
    /**
     * Export table data
     */
    export(options) {
      if (!this.table) {
        console.error("ExportPlugin: Table instance not available");
        return;
      }
      const state = this.table.getState();
      let rows = state.rows;
      if (options.selectedOnly) {
        rows = this.table.getSelectedRows();
      } else if (options.visibleOnly) {
        rows = state.visibleRows;
      }
      const columns = options.columns ? state.columns.filter((col) => options.columns.includes(col.id)) : state.visibleColumns;
      let content;
      let mimeType;
      let extension;
      switch (options.format) {
        case "csv":
          content = this.toCSV(rows, columns, options.includeHeaders ?? true);
          mimeType = "text/csv";
          extension = "csv";
          break;
        case "json":
          content = this.toJSON(rows, columns);
          mimeType = "application/json";
          extension = "json";
          break;
        case "txt":
          content = this.toTXT(rows, columns, options.includeHeaders ?? true);
          mimeType = "text/plain";
          extension = "txt";
          break;
        case "xlsx":
          console.warn("ExportPlugin: XLSX export requires additional library");
          return;
        default:
          console.error(`ExportPlugin: Unsupported format ${options.format}`);
          return;
      }
      this.download(
        content,
        options.filename || `export-${Date.now()}`,
        mimeType,
        extension
      );
    }
    /**
     * Convert to CSV format
     */
    toCSV(rows, columns, includeHeaders) {
      const lines = [];
      if (includeHeaders) {
        const headers = columns.map((col) => this.escapeCSV(col.label || col.id));
        lines.push(headers.join(","));
      }
      rows.forEach((row) => {
        const values = columns.map((col) => {
          const value = this.getCellValue(row, col);
          return this.escapeCSV(value);
        });
        lines.push(values.join(","));
      });
      return lines.join("\n");
    }
    /**
     * Convert to JSON format
     */
    toJSON(rows, columns) {
      const data = rows.map((row) => {
        const obj = {};
        columns.forEach((col) => {
          obj[col.label || col.id] = this.getCellValue(row, col);
        });
        return obj;
      });
      return JSON.stringify(data, null, 2);
    }
    /**
     * Convert to plain text format
     */
    toTXT(rows, columns, includeHeaders) {
      const lines = [];
      const widths = columns.map((col) => {
        const headerWidth = (col.label || col.id).length;
        const maxDataWidth = rows.length > 0 ? Math.max(...rows.map((row) => String(this.getCellValue(row, col)).length), 0) : 0;
        return Math.max(headerWidth, maxDataWidth);
      });
      if (includeHeaders) {
        const headers = columns.map((col, i) => (col.label || col.id).padEnd(widths[i] || 0)).join(" | ");
        lines.push(headers);
        lines.push(widths.map((w) => "-".repeat(w || 0)).join("-+-"));
      }
      rows.forEach((row) => {
        const values = columns.map((col, i) => String(this.getCellValue(row, col)).padEnd(widths[i] || 0)).join(" | ");
        lines.push(values);
      });
      return lines.join("\n");
    }
    /**
     * Get cell value
     */
    getCellValue(row, column) {
      if (column.accessor) {
        return column.accessor(row.data);
      }
      if (column.formatter) {
        return column.formatter(row.data[column.field], row.data);
      }
      return row.data[column.field];
    }
    /**
     * Escape CSV value
     */
    escapeCSV(value) {
      if (value === null || value === void 0) {
        return "";
      }
      const str = String(value);
      if (str.includes(",") || str.includes('"') || str.includes("\n")) {
        return `"${str.replace(/"/g, '""')}"`;
      }
      return str;
    }
    /**
     * Download content as file
     */
    download(content, filename, mimeType, extension) {
      const blob = new Blob([content], { type: mimeType });
      const url = URL.createObjectURL(blob);
      const link = document.createElement("a");
      link.href = url;
      link.download = `${filename}.${extension}`;
      link.click();
      setTimeout(() => URL.revokeObjectURL(url), 100);
    }
  };
  function createExportPlugin() {
    return new ExportPlugin();
  }

  // js/datatable/src/plugins/virtual-scroll.ts
  var VirtualScrollPlugin = class {
    constructor(options) {
      this.name = "virtualScroll";
      this.version = "1.0.0";
      this.table = null;
      this.scrollContainer = null;
      this.rafId = null;
      this.options = {
        overscan: 5,
        ...options
      };
      this.state = {
        scrollTop: 0,
        visibleStartIndex: 0,
        visibleEndIndex: 0,
        virtualRows: [],
        totalHeight: 0,
        offsetY: 0
      };
    }
    install(table) {
      this.table = table;
      table.on("data:change", this.handleDataChange.bind(this));
      table.on("state:change", this.handleStateChange.bind(this));
    }
    uninstall(_table) {
      this.cleanup();
    }
    /**
     * Attach to scroll container
     */
    attachToContainer(container) {
      this.scrollContainer = container;
      this.scrollContainer.addEventListener("scroll", this.handleScroll.bind(this));
      this.calculateVisibleRows();
    }
    /**
     * Detach from scroll container
     */
    detachFromContainer() {
      if (this.scrollContainer) {
        this.scrollContainer.removeEventListener("scroll", this.handleScroll.bind(this));
        this.scrollContainer = null;
      }
    }
    /**
     * Get virtual scroll state
     */
    getState() {
      return { ...this.state };
    }
    /**
     * Handle scroll event
     */
    handleScroll(_event) {
      if (!this.scrollContainer)
        return;
      if (this.rafId) {
        cancelAnimationFrame(this.rafId);
      }
      this.rafId = requestAnimationFrame(() => {
        this.state.scrollTop = this.scrollContainer.scrollTop;
        this.calculateVisibleRows();
        if (this.options.onScroll) {
          this.options.onScroll(this.state.scrollTop);
        }
      });
    }
    /**
     * Calculate visible rows based on scroll position
     */
    calculateVisibleRows() {
      if (!this.table)
        return;
      const tableState = this.table.getState();
      let rows;
      if (tableState.paginationState?.enabled && tableState.paginatedRows.length > 0) {
        rows = tableState.paginatedRows;
      } else {
        rows = tableState.filteredRows.length > 0 ? tableState.filteredRows : tableState.rows;
      }
      const totalRows = rows.length;
      const { rowHeight, containerHeight, overscan = 5 } = this.options;
      this.state.totalHeight = totalRows * rowHeight;
      if (containerHeight <= 0) {
        this.state.visibleStartIndex = 0;
        this.state.visibleEndIndex = 0;
        this.state.offsetY = 0;
        this.state.virtualRows = [];
        return;
      }
      const scrollTop = this.state.scrollTop;
      const visibleStart = Math.floor(scrollTop / rowHeight);
      const visibleEnd = Math.ceil((scrollTop + containerHeight) / rowHeight);
      const start = Math.max(0, visibleStart - overscan);
      const end = Math.min(totalRows, visibleEnd + overscan);
      this.state.visibleStartIndex = start;
      this.state.visibleEndIndex = end;
      this.state.offsetY = start * rowHeight;
      this.state.virtualRows = rows.slice(start, end);
    }
    /**
     * Handle data change
     */
    handleDataChange(_rows) {
      this.calculateVisibleRows();
    }
    /**
     * Handle state change
     */
    handleStateChange() {
      this.calculateVisibleRows();
    }
    /**
     * Scroll to row
     */
    scrollToRow(index) {
      if (!this.scrollContainer)
        return;
      const scrollTop = index * this.options.rowHeight;
      this.scrollContainer.scrollTop = scrollTop;
    }
    /**
     * Scroll to top
     */
    scrollToTop() {
      this.scrollToRow(0);
    }
    /**
     * Scroll to bottom
     */
    scrollToBottom() {
      if (!this.table)
        return;
      const tableState = this.table.getState();
      let rows;
      if (tableState.paginationState?.enabled && tableState.paginatedRows.length > 0) {
        rows = tableState.paginatedRows;
      } else {
        rows = tableState.filteredRows.length > 0 ? tableState.filteredRows : tableState.rows;
      }
      this.scrollToRow(rows.length - 1);
    }
    /**
     * Cleanup
     */
    cleanup() {
      this.detachFromContainer();
      if (this.rafId) {
        cancelAnimationFrame(this.rafId);
        this.rafId = null;
      }
    }
  };
  function createVirtualScrollPlugin(options) {
    return new VirtualScrollPlugin(options);
  }
  function measureRowHeight(element) {
    const computed = window.getComputedStyle(element);
    const marginTop = parseFloat(computed.marginTop) || 0;
    const marginBottom = parseFloat(computed.marginBottom) || 0;
    return element.offsetHeight + marginTop + marginBottom;
  }
  function calculateOptimalOverscan(rowHeight, containerHeight) {
    if (rowHeight <= 0 || containerHeight <= 0) {
      return 3;
    }
    const visibleRows = Math.ceil(containerHeight / rowHeight);
    return Math.max(3, Math.floor(visibleRows * 0.5));
  }

  // js/datatable/src/index.ts
  var VERSION = "2.0.0";
  var src_default = {
    DataTable,
    createDataTable,
    VERSION
  };
  return __toCommonJS(src_exports);
})();
/**
 * Awo DataTable
 * Headless DataTable for any design system
 * 
 * @package @awo/datatable
 * @author Awo ERP
 * @version 2.0.0
 * @license MIT
 */
//# sourceMappingURL=datatable.js.map
