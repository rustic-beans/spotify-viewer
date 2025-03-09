const WebSocket = require('ws');
const { createClient } = require('graphql-ws');
const { print } = require('graphql');
const gql = require('graphql-tag');
const dotenv = require('dotenv');

// Load the native module
const nativeCrash = require('./build/Release/native_crash');

// Load environment variables from .env file
dotenv.config();

// Configuration
const GRAPHQL_ENDPOINT = process.env.GRAPHQL_ENDPOINT || 'ws://127.0.0.1:8080/query';
const SUBSCRIPTION_QUERY = gql`
  subscription {
    playerState {
      progressMs
      track {
        name
        album {
          name
          images {
            url
          }
        }
        artists {
          name
        }
      }
    }
  }
`;

// Function to generate a random crash time between 5 and 60 seconds
function getRandomCrashTime() {
  return Math.floor(Math.random() * 55000) + 5000; // Between 5 and 60 seconds
}

// Create a GraphQL WebSocket client
const client = createClient({
  url: GRAPHQL_ENDPOINT,
  webSocketImpl: WebSocket,
  connectionParams: {
    // Add any authentication headers if needed
    // authToken: 'your-auth-token',
  },
  on: {
    connected: () => console.log('Successfully connected to GraphQL WebSocket\nWaiting for subscription messages...'),
    error: (error) => console.error('WebSocket connection error:', error),
    connecting: () => console.log('Connecting to WebSocket...'),
    closed: () => console.log('WebSocket connection closed')
  }
});

console.log('Connecting to GraphQL WebSocket endpoint...');

// Set up the subscription
const unsubscribe = client.subscribe(
  {
    query: print(SUBSCRIPTION_QUERY),
    variables: {}, // Add any variables your subscription needs
  },
  {
    next: (data) => {
      console.log('Received message:');
      console.log(JSON.stringify(data, null));
    },
    error: (error) => {
      console.error('Subscription error:', error);
    },
    complete: () => {
      console.log('Subscription completed');
    },
  }
);

// Schedule the random crash
const crashTime = getRandomCrashTime();
console.log(`Application will crash in ${crashTime / 1000} seconds`);

setTimeout(() => {
  console.error('SIMULATING APPLICATION CRASH! USING THE NUCLEAR OPTION');

  nativeCrash.crash();
}, crashTime);

// Handle SIGINT (Ctrl+C) to gracefully close the connection
// We intentionally don't handle SIGKILL to create a "dirty" disconnection
process.on('SIGINT', () => {
  console.log('Received SIGINT. Closing WebSocket connection...');
  unsubscribe();
  client.dispose();
  process.exit(0);
});
