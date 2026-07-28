import { createVuetify } from 'vuetify';
import { md3 } from 'vuetify/blueprints';
import * as components from 'vuetify/components';
import * as directives from 'vuetify/directives';

import '@mdi/font/css/materialdesignicons.css';
import 'vuetify/styles';

const savedLanguage = localStorage.getItem('language') || 'th';
document.documentElement.setAttribute('lang', savedLanguage);

export const vuetify = createVuetify({
  blueprint: md3,
  components,
  directives,
  icons: {
    defaultSet: 'mdi',
  },
  theme: {
    defaultTheme: 'auzy',
    themes: {
      auzy: {
        dark: false,
        colors: {
          background: '#eef3f0',
          surface: '#f7faf8',
          primary: '#1f7a66',
          secondary: '#3d524a',
          error: '#b42318',
          info: '#2563eb',
          success: '#1f7a66',
          warning: '#b45309',
          'on-background': '#12201c',
          'on-surface': '#12201c',
          'on-primary': '#ffffff',
        },
      },
    },
  },
  defaults: {
    VBtn: {
      rounded: 'lg',
      style: 'letter-spacing: 0; text-transform: none; font-weight: 600;',
    },
    VTextField: {
      variant: 'outlined',
      density: 'comfortable',
      color: 'primary',
      hideDetails: 'auto',
    },
    VSelect: {
      variant: 'outlined',
      density: 'comfortable',
      color: 'primary',
      hideDetails: 'auto',
    },
    VChip: {
      rounded: 'lg',
    },
    VDataTable: {
      hover: true,
    },
  },
});
