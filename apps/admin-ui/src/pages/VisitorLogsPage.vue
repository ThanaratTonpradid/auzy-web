<script setup>
import { computed, onMounted, ref } from 'vue';
import { useDisplay } from 'vuetify';
import { useI18n } from 'vue-i18n';
import { listVisitorLogsService } from '../services/visitor';
import { useAppStore } from '../stores/app';
import { useServerTablePagination } from '../composables/useTablePagination';
import AdminDataTable from '../components/AdminDataTable.vue';
import TablePagination from '../components/TablePagination.vue';

const appStore = useAppStore();
const { mdAndUp } = useDisplay();
const { t } = useI18n();

const items = ref([]);
const detailOpen = ref(false);
const selected = ref(null);

const pagination = {};

const load = async () => {
  try {
    appStore.setLoading(true);
    const res = await listVisitorLogsService({
      page: pagination.page.value,
      limit: pagination.limit.value,
    });
    items.value = (res.items || []).map((item) => ({
      ...item,
      location: formatLocation(item),
      createdAtLabel: formatTime(item.createdAt),
    }));
    pagination.syncTotal(res.total || 0);
  } catch (error) {
    appStore.setError(error.message || t('error.unknownError'));
  } finally {
    appStore.setLoading(false);
  }
};

Object.assign(pagination, useServerTablePagination({ load }));
const { page, limit, total } = pagination;

const headers = computed(() => [
  { title: 'ID', key: 'id', width: 80 },
  { title: t('visitors.ip'), key: 'ip' },
  { title: t('visitors.location'), key: 'location', sortable: false },
  { title: t('visitors.path'), key: 'path' },
  { title: t('visitors.userAgent'), key: 'userAgent' },
  { title: t('visitors.createdAt'), key: 'createdAt' },
  {
    title: t('common.actions'),
    key: 'actions',
    sortable: false,
    align: 'end',
    width: 120,
  },
]);

const selectedMeta = computed(() => selected.value?.metadata || {});
const totalText = computed(() => t('visitors.total', { count: total.value }));

const formatLocation = (item) => {
  const meta = item.metadata || {};
  const parts = [meta.city, meta.region, meta.country].filter(Boolean);
  return parts.length ? parts.join(', ') : '-';
};

const formatTime = (unix) => {
  if (!unix) return '-';
  return new Date(unix * 1000).toLocaleString();
};

const formatCoord = (value) => {
  if (value === null || value === undefined || Number.isNaN(Number(value))) {
    return '-';
  }
  return Number(value).toFixed(5);
};

const openDetails = (item) => {
  selected.value = item;
  detailOpen.value = true;
};

const closeDetails = () => {
  detailOpen.value = false;
  selected.value = null;
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

    <AdminDataTable
      v-if="mdAndUp"
      :headers="headers"
      :items="items"
      :items-per-page="limit"
    >
      <template #[`item.createdAt`]="{ item }">
        {{ item.createdAtLabel }}
      </template>
      <template #[`item.userAgent`]="{ item }">
        <span class="ua-cell">
          {{ item.userAgent || '-' }}
        </span>
      </template>
      <template #[`item.actions`]="{ item }">
        <v-btn
          size="small"
          variant="tonal"
          color="primary"
          rounded="lg"
          prepend-icon="mdi-eye-outline"
          @click="openDetails(item)"
        >
          {{ $t('common.details') }}
        </v-btn>
      </template>
    </AdminDataTable>

    <section v-else class="mobile-list">
      <article
        v-for="item in items"
        :key="item.id"
        class="surface-panel surface-panel--pad mobile-card"
      >
        <div class="mobile-card__head">
          <div>
            <h2 class="mobile-card__title mobile-card__title--break">{{ item.ip || '-' }}</h2>
            <p class="mobile-card__meta">#{{ item.id }} · {{ item.createdAtLabel }}</p>
          </div>
          <span class="meta-chip">{{ item.path || '/' }}</span>
        </div>
        <dl class="mobile-card__details">
          <div>
            <dt>{{ $t('visitors.location') }}</dt>
            <dd>{{ item.location }}</dd>
          </div>
          <div>
            <dt>{{ $t('visitors.userAgent') }}</dt>
            <dd class="mobile-card__ua">{{ item.userAgent || '-' }}</dd>
          </div>
        </dl>
        <v-btn
          class="mt-3"
          block
          size="small"
          variant="tonal"
          color="primary"
          rounded="lg"
          prepend-icon="mdi-eye-outline"
          @click="openDetails(item)"
        >
          {{ $t('common.details') }}
        </v-btn>
      </article>
      <p v-if="!items.length" class="mobile-list__empty">{{ $t('common.noData') }}</p>
    </section>

    <TablePagination
      v-model:page="page"
      v-model:limit="limit"
      :total="total"
      :total-text="totalText"
    />

    <v-dialog v-model="detailOpen" :fullscreen="!mdAndUp" max-width="560" @after-leave="closeDetails">
      <v-card v-if="selected" rounded="xl">
        <v-card-title class="text-h6 d-flex align-center justify-space-between">
          <span>{{ $t('visitors.detailTitle') }}</span>
          <v-btn icon="mdi-close" variant="text" @click="detailOpen = false" />
        </v-card-title>
        <v-card-text>
          <dl class="detail-grid">
            <div>
              <dt>ID</dt>
              <dd>{{ selected.id }}</dd>
            </div>
            <div>
              <dt>{{ $t('visitors.ip') }}</dt>
              <dd>{{ selected.ip || '-' }}</dd>
            </div>
            <div>
              <dt>{{ $t('visitors.path') }}</dt>
              <dd>{{ selected.path || '-' }}</dd>
            </div>
            <div>
              <dt>{{ $t('visitors.referer') }}</dt>
              <dd class="detail-grid__break">{{ selected.referer || '-' }}</dd>
            </div>
            <div>
              <dt>{{ $t('visitors.createdAt') }}</dt>
              <dd>{{ selected.createdAtLabel }}</dd>
            </div>
            <div>
              <dt>{{ $t('visitors.country') }}</dt>
              <dd>{{ selectedMeta.country || '-' }}</dd>
            </div>
            <div>
              <dt>{{ $t('visitors.region') }}</dt>
              <dd>{{ selectedMeta.region || '-' }}</dd>
            </div>
            <div>
              <dt>{{ $t('visitors.city') }}</dt>
              <dd>{{ selectedMeta.city || '-' }}</dd>
            </div>
            <div>
              <dt>{{ $t('visitors.latitude') }}</dt>
              <dd>{{ formatCoord(selectedMeta.latitude) }}</dd>
            </div>
            <div>
              <dt>{{ $t('visitors.longitude') }}</dt>
              <dd>{{ formatCoord(selectedMeta.longitude) }}</dd>
            </div>
            <div>
              <dt>{{ $t('visitors.source') }}</dt>
              <dd>{{ selectedMeta.source || '-' }}</dd>
            </div>
            <div class="detail-grid__full">
              <dt>{{ $t('visitors.userAgent') }}</dt>
              <dd class="detail-grid__break">{{ selected.userAgent || '-' }}</dd>
            </div>
          </dl>
        </v-card-text>
        <v-card-actions class="pa-4">
          <v-spacer />
          <v-btn color="primary" rounded="lg" @click="detailOpen = false">
            {{ $t('common.close') }}
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
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

.mobile-card__title--break {
  word-break: break-all;
}

.mobile-card__details {
  display: grid;
  gap: 0.75rem;
  margin: 0.95rem 0 0;
}

.mobile-card__details dt {
  font-size: 0.72rem;
  letter-spacing: 0.08em;
  text-transform: uppercase;
  color: var(--ink-soft);
}

.mobile-card__details dd {
  margin: 0.2rem 0 0;
  font-weight: 550;
  line-height: 1.45;
}

.mobile-card__ua {
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
  font-weight: 500;
  color: var(--ink-soft);
}

.detail-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 0.9rem 1rem;
  margin: 0;
}

.detail-grid dt {
  font-size: 0.72rem;
  letter-spacing: 0.08em;
  text-transform: uppercase;
  color: var(--ink-soft);
}

.detail-grid dd {
  margin: 0.25rem 0 0;
  font-weight: 600;
  line-height: 1.45;
}

.detail-grid__full {
  grid-column: 1 / -1;
}

.detail-grid__break {
  word-break: break-word;
  font-weight: 500;
}

@media (max-width: 720px) {
  .detail-grid {
    grid-template-columns: 1fr;
  }
}
</style>
