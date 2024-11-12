import './assets/main.css'

// Components
import App from './App.vue';

// Composables
import { createApp, h, provide } from 'vue';
import { DefaultApolloClient } from '@vue/apollo-composable';

// Plugins
import router from './router'
import { createPinia } from 'pinia'
import apolloClient from './graphql/apolloClient';

const app = createApp({
  setup() {
    provide(DefaultApolloClient, apolloClient);
  },

  render: () => h(App),
});

app.use(router);
app.use(createPinia());

app.mount('#app');

