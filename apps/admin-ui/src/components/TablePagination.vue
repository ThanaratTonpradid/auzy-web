<script setup>
import { computed } from 'vue';
import { useDisplay } from 'vuetify';
import { useI18n } from 'vue-i18n';
import { TABLE_LIMIT_OPTIONS } from '../composables/useTablePagination';

const props = defineProps({
  page: {
    type: Number,
    required: true,
  },
  limit: {
    type: Number,
    required: true,
  },
  total: {
    type: Number,
    required: true,
  },
  totalText: {
    type: String,
    default: '',
  },
});

const emit = defineEmits(['update:page', 'update:limit']);

const { t } = useI18n();
const { mdAndUp } = useDisplay();

const pageCount = computed(() => Math.max(1, Math.ceil(props.total / props.limit) || 1));

const resolvedTotalText = computed(
  () => props.totalText || t('common.totalCount', { count: props.total })
);

const pageModel = computed({
  get: () => props.page,
  set: (value) => emit('update:page', value),
});

const limitModel = computed({
  get: () => props.limit,
  set: (value) => emit('update:limit', value),
});
</script>

<template>
  <div class="table-pagination">
    <div class="table-pagination__meta">
      <div class="table-pagination__total">{{ resolvedTotalText }}</div>
      <div class="table-pagination__page">
        {{ $t('common.pageOf', { page, total: pageCount }) }}
      </div>
    </div>

    <div class="table-pagination__controls">
      <v-select
        v-model="limitModel"
        :items="TABLE_LIMIT_OPTIONS"
        :label="$t('common.itemsPerPage')"
        density="compact"
        variant="outlined"
        hide-details
        class="table-pagination__limit"
      />
      <v-pagination
        v-model="pageModel"
        :length="pageCount"
        :total-visible="mdAndUp ? 7 : 3"
        density="comfortable"
        active-color="primary"
        rounded="lg"
      />
    </div>
  </div>
</template>

<style scoped>
.table-pagination {
  display: flex;
  align-items: center;
  justify-content: space-between;
  gap: 1rem;
  margin-top: 1rem;
  flex-wrap: wrap;
}

.table-pagination__meta {
  display: grid;
  gap: 0.2rem;
}

.table-pagination__total,
.table-pagination__page {
  color: var(--ink-soft);
  font-size: 0.9rem;
}

.table-pagination__controls {
  display: flex;
  align-items: center;
  gap: 0.75rem;
  flex-wrap: wrap;
  justify-content: flex-end;
}

.table-pagination__limit {
  width: 7.5rem;
  flex-shrink: 0;
}

@media (max-width: 720px) {
  .table-pagination {
    flex-direction: column;
    align-items: stretch;
  }

  .table-pagination__controls {
    justify-content: stretch;
  }

  .table-pagination__limit {
    width: 100%;
  }
}
</style>
