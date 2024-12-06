// Module for testing the go backend code of spotify-viewer

package main

import (
	"context"
	"dagger/spotify-viewer/internal/dagger"
	"fmt"
	"math"
	"math/rand/v2"
)

type SpotifyViewer struct{}

// Build a ready-to-use development environment
func (m *SpotifyViewer) BuildEnv(source *dagger.Directory) *dagger.Container {
	// create a Dagger cache volume for dependencies
	goCache := dag.CacheVolume("go-modules")
	return dag.Go().
		WithSource(source).
		WithModuleCache(goCache).
		Container().
		WithExec([]string{"go", "generate", "./ent"})
}

// Runs the `spotify-viewer` module tests
func (m *SpotifyViewer) Test(ctx context.Context, source *dagger.Directory) (string, error) {
	output, err := m.BuildEnv(source). // call the test runner
						WithExec([]string{"go", "test", "./lib/..."}).Stdout(ctx)
	if err != nil {
		return "", err
	}
	println(output)

	return "", nil
}

// Build the application container
// To Export please call it with ... export --path=./dist from the source folder
func (m *SpotifyViewer) Build_Bin(source *dagger.Directory) *dagger.Directory {
	return m.BuildEnv(source).
		WithExec([]string{"go", "build", "-C", "cmd", "-o", "../dist/backend"}).
		Directory("./dist")
}

func (m *SpotifyViewer) Publish(ctx context.Context, source *dagger.Directory) (string, error) {
	// Publish the application to the cloud
	// test the application
	_, err := m.Test(ctx, source)
	if err != nil {
		return "", err
	}

	bin := m.Build_Bin(source)

	return dag.Container().From("debian:bookworm-slim").
		WithFile("/bin/backend", bin.File("backend")).
		WithEntrypoint([]string{"/bin/backend"}).
		WithExposedPort(8080).
		Publish(ctx, fmt.Sprintf("ttl.sh/spotify-viewer-%.0f", math.Floor(rand.Float64()*1000000000)))

	// build the application

}
