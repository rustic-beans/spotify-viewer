package main

import (
	"context"
	"dagger/spotify-viewer/internal/dagger"
	"fmt"
	"math"
	"math/rand/v2"
)

type SpotifyViewer struct{}

func (m *SpotifyViewer) BuildNodeEnv(source *dagger.Directory) *dagger.Container {
	nodeCache := dag.CacheVolume("node")
	return dag.Container().
		From("node:22").
		WithDirectory("/source", source.Directory("./web")).
		WithMountedCache("/root/.npm", nodeCache).
		WithWorkdir("/source").
		WithExec([]string{"npm", "install"})
}

func (m *SpotifyViewer) BuildWeb(source *dagger.Directory) *dagger.Directory {
	return m.BuildNodeEnv(source).
		WithExec([]string{"npm", "run", "build"}).
		Directory("./dist")
}

// Build a ready-to-use development environment
func (m *SpotifyViewer) BuildGoEnv(source *dagger.Directory) *dagger.Container {
	// create a Dagger cache volume for dependencies
	goCache := dag.CacheVolume("go-modules")
	return dag.Go(dagger.GoOpts{Version: "1.23"}).
		WithSource(source).
		WithModuleCache(goCache).
		WithExec([]string{"make", "generate"}).
		Container()
}

// Runs the `spotify-viewer` module tests
func (m *SpotifyViewer) Test(ctx context.Context, source *dagger.Directory) (string, error) {
	output, err := m.BuildGoEnv(source). // call the test runner
						WithExec([]string{"go", "test", "./..."}).Stdout(ctx)
	if err != nil {
		return "", err
	}
	println(output)

	return "", nil
}

// Build the application container
// To Export please call it with ... export --path=./dist from the source folder
func (m *SpotifyViewer) Build_Bin(source *dagger.Directory) *dagger.Directory {
	return m.BuildGoEnv(source).
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
	web := m.BuildWeb(source)

	return dag.Container().From("debian:bookworm-slim").
		WithFile("/app/backend", bin.File("backend")).
		WithDirectory("/app/web", web).
		WithWorkdir("/app").
		WithEntrypoint([]string{"/app/backend"}).
		WithExposedPort(8080).
		Publish(ctx, fmt.Sprintf("ttl.sh/spotify-viewer-%.0f", math.Floor(rand.Float64()*1000000000)))

	// build the application

}
