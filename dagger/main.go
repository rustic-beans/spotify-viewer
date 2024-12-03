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
		WithExec([]string{"go", "generate", "./ent"})
}

func (m *SpotifyViewer) graphqlGenerate(source *dagger.Directory) *dagger.Container {
	return m.BuildEnv(source).
		WithExec([]string{"go", "generate", "./ent"})
}

// Runs the `spotify-viewer` module tests
func (m *SpotifyViewer) Test(ctx context.Context, source *dagger.Directory) (string, error) {
	return m.BuildEnv(source). // call the test runner
					WithExec([]string{"go", "test", "./..."}).
		// capture and return the command output
		Stdout(ctx)
}

// Build the application container
// To Export please call it with ... export --path=./dist from the source folder
func (m *SpotifyViewer) Build(source *dagger.Directory) *dagger.Directory {
	return m.BuildEnv(source).
		WithExec([]string{"go", "build", "-C", "cmd", "-o", "../dist/backend"}).
		Directory("./dist")
}
