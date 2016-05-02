package action

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/miclle/lisa/msg"
)

// Watcher func
func Watcher(name, command string) {
	if command == "" {
		msg.Info("lisa watching the path %s", name)
	} else {
		msg.Info("lisa watching the path %s then execute command: %s", name, command)
	}

	if watcher, err := NewRecursiveWatcher(name, command); err != nil {
		msg.Err(err.Error())
	} else {
		defer watcher.Close()
		done := make(chan bool)
		watcher.Run()
		<-done
	}
}

// RecursiveWatcher struct
type RecursiveWatcher struct {
	*fsnotify.Watcher
	*Walker
	Command string
}

// NewRecursiveWatcher return a recursive watcher
func NewRecursiveWatcher(name, command string) (*RecursiveWatcher, error) {
	rw := &RecursiveWatcher{
		Command: command,
		Walker: &Walker{
			IgnorePrefix: ".",
		},
	}

	folders := []string{}

	if fi, err := os.Stat(name); err != nil {
		msg.Err("error: %s", err.Error())
	} else if fi.IsDir() {
		folders = rw.Subfolders(name)
	} else {
		folders = append(folders, name)
	}

	if len(folders) == 0 {
		return nil, errors.New("No folders or file to watch.")
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	rw.Watcher = watcher

	for _, folder := range folders {
		rw.AddFolder(folder)
	}
	return rw, nil
}

// AddFolder add folder to recursive watcher
func (watcher *RecursiveWatcher) AddFolder(folder string) {
	err := watcher.Add(folder)
	if err != nil {
		msg.Err("Error watching: %s, %s", folder, err.Error())
	} else {
		msg.Info("Lisa watching: %s", folder)
	}
}

// execCommand execute the command
func (watcher *RecursiveWatcher) execCommand() {
	if watcher.Command == "" {
		return
	}

	cmd := exec.Command(watcher.Command)

	msg.Info(strings.Join(cmd.Args, " "))

	out, err := cmd.CombinedOutput()
	if err != nil {
		msg.Err(err.Error())
	}
	msg.Info(string(out))

	if cmd.ProcessState != nil && cmd.ProcessState.Success() {
		msg.Info("execute the command `%s` was PASS", watcher.Command)
	} else {
		msg.Info("execute the command `%s` was FAIL", watcher.Command)
	}

	if cmd.ProcessState != nil {
		msg.Info("execute the command latency (%.2f seconds)\n", cmd.ProcessState.UserTime().Seconds())
	}
}

// Run execute the recursive watcher
func (watcher *RecursiveWatcher) Run() {
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Create == fsnotify.Create {
					if fi, err := os.Stat(event.Name); err != nil {
						msg.Err("error: %s", err.Error())
					} else if fi.IsDir() {
						msg.Info("directory created: %s", event.Name)
						if !watcher.IgnoreFile(filepath.Base(event.Name)) {
							watcher.AddFolder(event.Name)
						}
					} else {
						msg.Info("file created: %s", event.Name)
					}
				}

				if event.Op&fsnotify.Remove == fsnotify.Remove {
					msg.Info("file remove: %s", event.Name)
				}

				if event.Op&fsnotify.Write == fsnotify.Write {
					msg.Info("file modified: %s", event.Name)
				}

				if event.Op&fsnotify.Rename == fsnotify.Rename {
					msg.Info("file rename: %s", event.Name)
				}
				if event.Op&fsnotify.Chmod == fsnotify.Chmod {
					msg.Info("file chmod: %s", event.Name)
				}

				if !watcher.IgnoreFile(event.Name) {
					watcher.execCommand()
				}

			case err := <-watcher.Errors:
				msg.Err("error: %s", err.Error())
			}
		}
	}()
}

// Walker a file path walker
type Walker struct {
	IgnorePrefix string
}

// Subfolders returns a slice of subfolders (recursive), including the folder provided.
func (walker *Walker) Subfolders(path string) (paths []string) {
	filepath.Walk(path, func(newPath string, info os.FileInfo, err error) error {
		if err != nil {
			msg.Err("error: %s", err.Error())
			return err
		}

		if info.IsDir() {
			name := info.Name()
			// skip folders that begin with a dot
			if walker.IgnoreFile(name) && name != "." && name != ".." {
				return filepath.SkipDir
			}
			paths = append(paths, newPath)
		}
		return nil
	})
	return paths
}

// IgnoreFile determines if a file should be ignored.
func (walker *Walker) IgnoreFile(name string) bool {
	return strings.HasPrefix(name, walker.IgnorePrefix)
}
