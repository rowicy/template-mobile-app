# template-mobile-app

Templates for mobile app development.

## [mobile-app](./mobile-app)

This is a template for mobile app development using:

- [Expo](https://expo.dev/)
- [React Native](https://reactnative.dev/)
- [Node.js](https://nodejs.org/) v22.14.0

## [api](./api)

This is a template for building APIs using:

- [Go](https://go.dev/) // TODO: 選定して更新

## [openapi-specifications](./openapi-specifications)

This directory contains OpenAPI specifications swagger files.

- [OpenAPI](https://www.openapis.org/)
- [Swagger](https://swagger.io/)

## [ls-lint](./.ls-lint.yml)

This is a configuration file for [ls-lint](https://ls-lint.org/), a linter for directory structures.  
GitHub Actions are set up to run ls-lint on pull requests to ensure that the directory structure adheres to the defined rules.  
For [api](./api) and [mobile-app](./mobile-app) directories, add settings for each directory, as they are ignored except for markdown files.
