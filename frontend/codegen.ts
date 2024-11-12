import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
  schema: [
    '../graph/*.graphql',
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
    "./src/__generated__/types.ts": {
      plugins: ["typescript", "typescript-operations", "typescript-vue-apollo"],
    }
  }
}

export default config;
