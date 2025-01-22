import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
  schema: [
    '../api/*.graphql',
    './src/graphql/*.graphql',
  ],
  documents: [
    'src/**/*.vue',
    '!src/gql/**/*'
  ],
  ignoreNoDocuments: true,
  generates: {
    './src/__generated__/': {
      preset: 'client',
      config: {
        useTypeImports: true,
        withCompositionFunctions: true,
        vueCompositionApiImportFrom: 'vue'
      }
    },
  }
}

export default config;
