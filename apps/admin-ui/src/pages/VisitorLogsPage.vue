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
  const parts = [item.city, item.region, item.country].filter(Boolean);
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
  <v-container>
    <div class="mb-4">
      <h1 class="text-h4">{{ $t('visitors.title') }}</h1>
      <p class="text-body-2 text-medium-emphasis">{{ $t('visitors.subtitle') }}</p>
    </div>

    <v-data-table
      :headers="headers"
      :items="items"
      item-value="id"
      class="elevation-0"
      border
      :items-per-page="limit"
      hide-default-footer
    >
      <template #[`item.createdAt`]="{ item }">
        {{ item.createdAtLabel }}
      </template>
      <template #[`item.userAgent`]="{ item }">
        <span class="text-truncate d-inline-block" style="max-width: 280px">
          {{ item.userAgent || '-' }}
        </span>
      </template>
    </v-data-table>

    <div class="d-flex align-center justify-space-between mt-4">
      <div class="text-body-2 text-medium-emphasis">
        {{ $t('visitors.total', { count: total }) }}
      </div>
      <div class="d-flex ga-2">
        <v-btn
          variant="tonal"
          :disabled="page <= 1"
          @click="onPageChange(page - 1)"
        >
          {{ $t('common.previous') }}
        </v-btn>
        <v-btn
          variant="tonal"
          :disabled="page * limit >= total"
          @click="onPageChange(page + 1)"
        >
          {{ $t('common.next') }}
        </v-btn>
      </div>
    </div>
  </v-container>
</template>
