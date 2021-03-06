package hg

import (
	"os"
	"os/exec"
)

// Clone makes a child process and runs `git clone repoURL`.
// The cloned repo is stored at the path.
func Clone(repoURL, path string) error {
	// Save the cwd.
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	defer os.Chdir(cwd)

	// cd into the path.
	if err := os.Chdir(path); err != nil {
		return err
	}

	// Setup the command.
	cmd := exec.Command("hg", "clone", repoURL)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// And execute it.
	return cmd.Run()
}
