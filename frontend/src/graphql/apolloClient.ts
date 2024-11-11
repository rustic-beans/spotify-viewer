import { ApolloClient, InMemoryCache, createHttpLink } from '@apollo/client/core';
import { provideApolloClient } from '@vue/apollo-composable';

// HTTP connection to the API
const link = createHttpLink({
  // TODO: Change this to the real URL
  uri: 'http://127.0.0.1:8080/query',
});

// Cache implementation
const cache = new InMemoryCache();

// Create the apollo client
const apolloClient = new ApolloClient({
  link: link,
  cache,
});

// Provide the client to the app
provideApolloClient(apolloClient);

export default apolloClient;

