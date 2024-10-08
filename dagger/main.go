// A generated module for DummyTests functions
//
// im testing stuff....uh oh..........

package main

import (
	"context"
	"dummy-tests/dagger/internal/dagger"
	"fmt"
	"strings"
)

type DummyTests struct{} //this struct must match name of the repo root, or Dagger can't understand the below funcs

// Returns a container that echoes whatever string argument is provided
func (m *DummyTests) ContainerEcho(ctx context.Context,
	// +optional
	// +default="capitola"
	stringArg string) (string, error) {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", stringArg}).Stdout(ctx)
}

// Returns lines that match a pattern in the files of the provided Directory
func (m *DummyTests) GrepDir(ctx context.Context, directoryArg *dagger.Directory, pattern string) (string, error) {
	return dag.Container().
		From("alpine:latest").
		WithMountedDirectory("/mnt", directoryArg).
		WithWorkdir("/mnt").
		WithExec([]string{"grep", "-R", pattern, "."}).
		Stdout(ctx)
}

func (m *DummyTests) Egg(str string) *dagger.Container {
	return dag.Container().From("alpine:latest").WithExec([]string{"echo", str})
}

var defaultFigletContainer = dag.
	Container().
	From("alpine:latest").
	WithExec([]string{
		"apk", "add", "figlet",
	})

func (m *DummyTests) Hello(ctx context.Context,
	// +optional
	// +default="hello"
	greeting string,
	// Change the name
	// +optional
	// +default="world"
	name string,
	// Encode the message in giant multi-character letters
	// +optional
	// +default=true
	giant bool,
	// Make the message uppercase, and add more exclamation points
	// +optional
	// +default=true
	shout bool,
	// Custom container for running the figlet tool
	// +optional
	figletContainer *dagger.Container,
) (string, error) {
	message := fmt.Sprintf("%s %s!", greeting, name)
	if shout {
		message = strings.ToUpper(message) + "!!!!!"
	}
	if giant {
		// Run 'figlet' in a container to produce giant letters
		ctr := figletContainer
		if ctr == nil {
			ctr = defaultFigletContainer
		}
		return ctr.
			WithoutEntrypoint(). // clear the entrypoint to make sure 'figlet' is executed
			WithExec([]string{"figlet", message}).
			Stdout(ctx)
	}
	return message, nil
}
