package main

import (
	"context"
	"dagger/spotify-viewer/internal/dagger"
	"fmt"
	"math"
	"math/rand/v2"
)

type SpotifyViewer struct{}

func (m *SpotifyViewer) BuildNodeEnv(
	// +ignore=["web/node_modules"]
	source *dagger.Directory,
) *dagger.Container {
	nodeCache := dag.CacheVolume("node")
	return dag.Container().
		From("node:22").
		WithDirectory("/source/web", source.Directory("./web")).
		WithDirectory("/source/api", source.Directory("./api")).
		WithMountedCache("/root/.npm", nodeCache).
		WithWorkdir("/source/web").
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
func (m *SpotifyViewer) BuildGoBin(source *dagger.Directory) *dagger.Directory {
	return m.BuildGoEnv(source).
		WithExec([]string{"go", "build", "-C", "cmd", "-o", "../dist/backend"}).
		Directory("./dist")
}

func (m *SpotifyViewer) Build(ctx context.Context, source *dagger.Directory) (*dagger.Container, error) {
	// Publish the application to the cloud
	// test the application
	_, err := m.Test(ctx, source)
	if err != nil {
		return nil, err
	}

	bin := m.BuildGoBin(source)
	web := m.BuildWeb(source)

	return dag.Container().From("debian:bookworm-slim").
		WithExec([]string{"apt-get", "update"}).
		WithExec([]string{"apt-get", "install", "-y", "ca-certificates"}).
		WithExec([]string{"update-ca-certificates"}).
		WithFile("/app/backend", bin.File("backend")).
		WithFile("/app/config.yaml", source.File("config-default.yaml")).
		WithDirectory("/app/web/dist", web).
		WithWorkdir("/app").
		WithEntrypoint([]string{"/app/backend"}).
		WithExposedPort(8080), nil
}

func (m *SpotifyViewer) PublishTTL(ctx context.Context, source *dagger.Directory) (string, error) {
	container, err := m.Build(ctx, source)
	if err != nil {
		return "", err
	}

	containerUrl := fmt.Sprintf("ttl.sh/spotify-viewer-%.0f", math.Floor(rand.Float64()*1000000000))

	return container.Publish(ctx, containerUrl)
}

func (m *SpotifyViewer) PublishGithub(
	ctx context.Context,
	source *dagger.Directory,
	registry,
	imageName,
	tag,
	username string,
	password *dagger.Secret,
) (string, error) {
	container, err := m.Build(ctx, source)
	if err != nil {
		return "", err
	}

	containerUrl := fmt.Sprintf("%s/%s", registry, imageName)
	containerTag := fmt.Sprintf("%s:%s", containerUrl, tag)
	containerLatest := fmt.Sprintf("%s:latest", containerUrl)

	authed := container.WithRegistryAuth(registry, username, password)

	_, err = authed.Publish(ctx, containerTag)
	if err != nil {
		return "", err
	}

	return authed.Publish(ctx, containerLatest)
}
