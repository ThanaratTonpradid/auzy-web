import { createVuetify } from 'vuetify';
import { md3 } from 'vuetify/blueprints';
import * as components from 'vuetify/components';
import * as directives from 'vuetify/directives';

import '@mdi/font/css/materialdesignicons.css';
import 'vuetify/styles';

const icons = {
  defaultSet: 'mdi',
};

// Get saved language to set initial font
const savedLanguage = localStorage.getItem('language') || 'th';
const fontFamily = savedLanguage === 'th' 
  ? "'Sukhumvit Set', -apple-system, BlinkMacSystemFont, sans-serif"
  : "'Prompt', -apple-system, BlinkMacSystemFont, sans-serif";

export const vuetify = createVuetify({
  blueprint: md3, // Material Design 3
  components,
  directives,
  icons,
  theme: {
    defaultTheme: 'light',
  },
  defaults: {
    VBtn: {
      style: 'letter-spacing: 0;',
    },
    global: {
      font: {
        family: fontFamily,
      },
    },
  },
});
