// Module for testing the go backend code of spotify-viewer

package main

import (
	"context"
	"dagger/spotify-viewer/internal/dagger"
)

type SpotifyViewer struct{}

// Build a ready-to-use development environment
func (m *SpotifyViewer) BuildEnv(source *dagger.Directory) *dagger.Container {
	// create a Dagger cache volume for dependencies
	goCache := dag.CacheVolume("go-modules")
	return dag.Container().
		// start from a base Node.js container
		From("golang").
		// add the source code at /src
		WithDirectory("/src", source).
		// mount the cache volume at /root/.npm
		WithMountedCache("/go-cache", goCache).
		// change the working directory to /src
		WithWorkdir("/src").
		// run npm install to install dependencies
		WithExec([]string{"go", "mod", "tidy"})
}

// Runs the `spotify-viewer` module tests
func (m *SpotifyViewer) Test(ctx context.Context, source *dagger.Directory) (string, error) {
	return m.BuildEnv(source).
		// call the test runner
		WithExec([]string{"go", "test", "./..."}).
		// capture and return the command output
		Stdout(ctx)
}
