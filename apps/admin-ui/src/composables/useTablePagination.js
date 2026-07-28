import { computed, ref, unref, watch } from 'vue';

export const TABLE_LIMIT_OPTIONS = [
  { title: '10', value: 10 },
  { title: '20', value: 20 },
  { title: '50', value: 50 },
];

export function useClientTablePagination(itemsSource, { initialLimit = 20 } = {}) {
  const page = ref(1);
  const limit = ref(initialLimit);

  const total = computed(() => (unref(itemsSource) || []).length);
  const pageCount = computed(() => Math.max(1, Math.ceil(total.value / limit.value) || 1));

  const pagedItems = computed(() => {
    const list = unref(itemsSource) || [];
    const start = (page.value - 1) * limit.value;
    return list.slice(start, start + limit.value);
  });

  watch(limit, () => {
    page.value = 1;
  });

  watch(total, (count) => {
    const maxPage = Math.max(1, Math.ceil(count / limit.value) || 1);
    if (page.value > maxPage) {
      page.value = maxPage;
    }
  });

  return {
    page,
    limit,
    total,
    pageCount,
    pagedItems,
    limitOptions: TABLE_LIMIT_OPTIONS,
  };
}

export function useServerTablePagination({ load, initialLimit = 20 } = {}) {
  const page = ref(1);
  const limit = ref(initialLimit);
  const total = ref(0);
  const pageCount = computed(() => Math.max(1, Math.ceil(total.value / limit.value) || 1));

  watch(limit, async () => {
    if (page.value !== 1) {
      page.value = 1;
      return;
    }
    if (load) {
      await load();
    }
  });

  watch(page, async () => {
    if (load) {
      await load();
    }
  });

  const syncTotal = (nextTotal) => {
    total.value = nextTotal || 0;
    if (page.value > pageCount.value) {
      page.value = pageCount.value;
    }
  };

  return {
    page,
    limit,
    total,
    pageCount,
    syncTotal,
    limitOptions: TABLE_LIMIT_OPTIONS,
  };
}
