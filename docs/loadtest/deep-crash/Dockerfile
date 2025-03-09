FROM node:22

# Create app directory
WORKDIR /app

# Install build tools needed for native modules
RUN apt-get update && apt-get install -y \
    python3 \
    make \
    g++ \
    && rm -rf /var/lib/apt/lists/*

# Install app dependencies
# A wildcard is used to ensure both package.json AND package-lock.json are copied
COPY package*.json ./
COPY binding.gyp ./
COPY native-crash.c ./

RUN npm install node-gyp -g && \
    npm install && \
    node-gyp configure build

# Bundle app source
COPY . .

# Environment variables with defaults
ENV GRAPHQL_ENDPOINT=ws://localhost:8080/query

# Run the script
CMD ["node", "main.js"]
