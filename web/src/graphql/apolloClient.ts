import { ApolloClient, InMemoryCache, createHttpLink, split } from '@apollo/client/core';
import { getMainDefinition } from '@apollo/client/utilities';
import { GraphQLWsLink } from '@apollo/client/link/subscriptions';
import { createClient } from 'graphql-ws';
import { provideApolloClient } from '@vue/apollo-composable';

// Cache implementation
const cache = new InMemoryCache();

const host = (import.meta.env.MODE === 'production') ? window.location.host : '127.0.0.1:8080'

// HTTP connection to the API
const httpLink = createHttpLink({
  // TODO: Change this to the real URL
  uri: `http://${host}/query`,
});

// HTTP connection to the API
const wsLink = new GraphQLWsLink(createClient({
  // TODO: Change this to the real URL
  url: `ws://${host}/query`,
}));

const splitLink = split(
  ({ query }) => {
    const definition = getMainDefinition(query);
    return (
      definition.kind === 'OperationDefinition' &&
      definition.operation === 'subscription'
    );
  },
  wsLink,
  httpLink,
);

// Create the apollo client
const apolloClient = new ApolloClient({
  link: splitLink,
  cache,
});

// Provide the client to the app
provideApolloClient(apolloClient);

export default apolloClient;

