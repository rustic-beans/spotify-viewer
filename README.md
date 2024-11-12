<h1 align="center">Spotify Viewer</h1>
<h3 align="center">A website for viewing what is currently playing on Spotify</h3>

<p align="center">
    <img src="docs/imgs/spotify_viewer_logo.png" alt="Logo" height="300" width="300" />
</p>

<p align="center">
    <a href="https://github.com/rustic-beans/spotify-viewer/releases/latest"><img src="https://img.shields.io/github/v/release/rustic-beans/spotify-viewer?logo=github" alt="Github Latest Release"></a>
    <a href="https://github.com/rustic-beans/spotify-viewer/actions/workflows/build_test.yml"> <img src="https://img.shields.io/github/actions/workflow/status/rustic-beans/spotify-viewer/build_test" alt="Github Actions Build" /></a>
    <a href="https://github.com/rustic-beans/spotify-viewer/actions/workflows/test_on_push.yml"> <img src="https://img.shields.io/github/actions/workflow/status/rustic-beans/spotify-viewer/test_on_push?label=tests" alt="Github Actions Tests" /></a>
    <a href="https://pkg.go.dev/github.com/rustic-beans/spotify-viewer"> <img src="https://img.shields.io/badge/_-reference-blue?logo=go&label=%E2%80%8E%20" alt="Go Reference" /></a>
</p>

## Wnat is it?
This project is a basic server-client application that allows users to view what is currently playing on their Spotify account through a website. The backend server is written in Go and uses the Spotify API to get information about the user's spotify account. The server will log songs that the user listens to in a SQLite database, the logic for when the server logs a song is explained in the docs [here](docs/spotify-song-logging.md) WIP. The frontend is written in Vue 3. The frontend will display the currently playing song (e.g. album art, artist name and song name), how far into the song the user is, as well as a small history of the last songs played and the future queue of songs to be played.
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
```
cd frontend
npm i
npm run generate
npm run dev
```

## Host
TODO

