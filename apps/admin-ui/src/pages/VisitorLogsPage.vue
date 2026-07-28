<script setup>
import { computed, onMounted, ref } from 'vue';
import { listVisitorLogsService } from '../services/visitor';
import { useAppStore } from '../stores/app';
import { i18n } from '../plugins/i18n';

const appStore = useAppStore();
const items = ref([]);
const total = ref(0);
const page = ref(1);
const limit = ref(20);

const headers = computed(() => [
  { title: 'ID', key: 'id', width: 80 },
  { title: i18n.global.t('visitors.ip'), key: 'ip' },
  { title: i18n.global.t('visitors.location'), key: 'location', sortable: false },
  { title: i18n.global.t('visitors.path'), key: 'path' },
  { title: i18n.global.t('visitors.userAgent'), key: 'userAgent' },
  { title: i18n.global.t('visitors.createdAt'), key: 'createdAt' },
]);

const formatLocation = (item) => {
  const meta = item.metadata || {};
  const parts = [meta.city, meta.region, meta.country].filter(Boolean);
  return parts.length ? parts.join(', ') : '-';
};

const formatTime = (unix) => {
  if (!unix) return '-';
  return new Date(unix * 1000).toLocaleString();
};

const load = async () => {
  try {
    appStore.setLoading(true);
    const res = await listVisitorLogsService({ page: page.value, limit: limit.value });
    items.value = (res.items || []).map((item) => ({
      ...item,
      location: formatLocation(item),
      createdAtLabel: formatTime(item.createdAt),
    }));
    total.value = res.total || 0;
  } catch (error) {
    appStore.setError(error.message || i18n.global.t('error.unknownError'));
  } finally {
    appStore.setLoading(false);
  }
};

const onPageChange = async (nextPage) => {
  page.value = nextPage;
  await load();
};

onMounted(load);
</script>

<template>
  <div class="page-shell">
    <header class="page-header">
      <p class="page-header__eyebrow">{{ $t('common.appName') }}</p>
      <h1 class="page-header__title">{{ $t('visitors.title') }}</h1>
      <p class="page-header__subtitle">{{ $t('visitors.subtitle') }}</p>
    </header>

    <section class="surface-panel quiet-table overflow-hidden">
      <v-data-table
        :headers="headers"
        :items="items"
        item-value="id"
        class="bg-transparent"
        :items-per-page="limit"
        hide-default-footer
      >
        <template #[`item.createdAt`]="{ item }">
          {{ item.createdAtLabel }}
        </template>
        <template #[`item.userAgent`]="{ item }">
          <span class="ua-cell">
            {{ item.userAgent || '-' }}
          </span>
        </template>
      </v-data-table>
    </section>

    <div class="visitors-footer">
      <div class="visitors-footer__total">
        {{ $t('visitors.total', { count: total }) }}
      </div>
      <div class="visitors-footer__pager">
        <v-btn
          variant="tonal"
          color="primary"
          rounded="lg"
          :disabled="page <= 1"
          @click="onPageChange(page - 1)"
        >
          {{ $t('common.previous') }}
        </v-btn>
        <v-btn
          variant="tonal"
          color="primary"
          rounded="lg"
          :disabled="page * limit >= total"
          @click="onPageChange(page + 1)"
        >
          {{ $t('common.next') }}
        </v-btn>
      </div>
    </div>
  </div>
</template>

<style scoped>
.ua-cell {
  display: inline-block;
  max-width: 280px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
  vertical-align: bottom;
}

.visitors-footer {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  margin-top: 1rem;
}

.visitors-footer__total {
  color: var(--ink-soft);
  font-size: 0.9rem;
}

.visitors-footer__pager {
  display: flex;
  gap: 0.5rem;
}

@media (max-width: 720px) {
  .visitors-footer {
    flex-direction: column;
    align-items: stretch;
  }
}
</style>
