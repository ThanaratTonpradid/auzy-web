<script setup>
import { computed, onMounted, ref } from 'vue';
import { useDisplay } from 'vuetify';
import { listVisitorLogsService } from '../services/visitor';
import { useAppStore } from '../stores/app';
import { i18n } from '../plugins/i18n';

const appStore = useAppStore();
const { mdAndUp } = useDisplay();
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

    <section v-if="mdAndUp" class="surface-panel quiet-table overflow-hidden">
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

    <section v-else class="mobile-list">
      <article
        v-for="item in items"
        :key="item.id"
        class="surface-panel surface-panel--pad mobile-card"
      >
        <div class="mobile-card__head">
          <div>
            <h2 class="mobile-card__title">{{ item.ip || '-' }}</h2>
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
      </article>
      <p v-if="!items.length" class="mobile-list__empty">{{ $t('common.noData') }}</p>
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
          class="visitors-footer__btn"
          :disabled="page <= 1"
          @click="onPageChange(page - 1)"
        >
          {{ $t('common.previous') }}
        </v-btn>
        <v-btn
          variant="tonal"
          color="primary"
          rounded="lg"
          class="visitors-footer__btn"
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

.mobile-list {
  display: grid;
  gap: 0.75rem;
}

.mobile-card__head {
  display: flex;
  align-items: flex-start;
  justify-content: space-between;
  gap: 0.75rem;
}

.mobile-card__title {
  margin: 0;
  font-size: 1.05rem;
  letter-spacing: -0.02em;
  word-break: break-all;
}

.mobile-card__meta {
  margin: 0.25rem 0 0;
  font-size: 0.82rem;
  color: var(--ink-soft);
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

.mobile-list__empty {
  margin: 0;
  padding: 2rem 1rem;
  text-align: center;
  color: var(--ink-soft);
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

  .visitors-footer__pager {
    display: grid;
    grid-template-columns: 1fr 1fr;
  }

  .visitors-footer__btn {
    width: 100%;
  }
}
</style>
