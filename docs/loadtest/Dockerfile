FROM node:22-alpine

# Create app directory
WORKDIR /app

# Install app dependencies
# A wildcard is used to ensure both package.json AND package-lock.json are copied
COPY package*.json ./

RUN npm install

# Bundle app source
COPY . .

# Environment variables with defaults
ENV GRAPHQL_ENDPOINT=ws://localhost:8080/query

# Run the script
CMD ["npm", "run", "start"]
