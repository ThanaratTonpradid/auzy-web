<script setup>
defineProps({
  headers: {
    type: Array,
    required: true,
  },
  items: {
    type: Array,
    default: () => [],
  },
  itemValue: {
    type: String,
    default: 'id',
  },
  itemsPerPage: {
    type: [Number, String],
    default: -1,
  },
});
</script>

<template>
  <section class="surface-panel quiet-table overflow-hidden admin-data-table">
    <v-data-table
      :headers="headers"
      :items="items"
      :item-value="itemValue"
      class="bg-transparent"
      :items-per-page="itemsPerPage"
      hide-default-footer
    >
      <template v-for="(_, name) in $slots" #[name]="slotData">
        <slot :name="name" v-bind="slotData || {}" />
      </template>
    </v-data-table>
  </section>
</template>
