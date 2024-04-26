# A website for viewing what is currently playing on Spotify
## Usage
### Run
To run the backend first configure the config variables by copying `config-example.yaml` to `config.yaml` and filling in the values.

To get the `clientId` and `clientSecret` first follow the instructions [here to create an app](https://developer.spotify.com/documentation/web-api/tutorials/getting-started#create-an-app), and then get the values from the [Spotfy App dashboard](https://developer.spotify.com/dashboard).

Then you have to generate the GraphQL schema and database types by running
```bash
make generate
```
Every time you change something in `ent/schema` you have to run this command again.

Then run the backend with the following command:
```bash
make start
```

### Test
TODO

### Frontend 
TODO

## Host
TODO

