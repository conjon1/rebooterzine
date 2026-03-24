package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/fsnotify/fsnotify"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("Error creating watcher:", err)
	}
	defer watcher.Close()

	dirsToWatch := []string{"posts", "authors", "components"}
	for _, dir := range dirsToWatch {
		err = watcher.Add(dir)
		if err != nil {
			log.Printf("Warning: Could not watch directory %s: %v", dir, err)
		} else {
			log.Printf("Watching directory: %s", dir)
		}
	}

	log.Println("Watcher ready! Drop an HTML file into 'posts' or 'authors' to trigger a build.")

	var buildTimer *time.Timer

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}

			if event.Has(fsnotify.Chmod) || strings.HasSuffix(event.Name, "_templ.go") {
				continue
			}

			log.Printf("Event detected: %s (%s)", event.Name, event.Op)

			if buildTimer != nil {
				buildTimer.Stop()
			}

			buildTimer = time.AfterFunc(500*time.Millisecond, buildSite)

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Println("Watcher error:", err)
		}
	}
}

func buildSite() {
	log.Println("Rebuilding site...")

	templCmd := exec.Command("templ", "generate")
	templCmd.Stdout = os.Stdout
	templCmd.Stderr = os.Stderr
	if err := templCmd.Run(); err != nil {
		log.Printf("Templ generate failed: %v", err)
	}

	buildCmd := exec.Command("go", "run", "./cmd/build")
	buildCmd.Stdout = os.Stdout
	buildCmd.Stderr = os.Stderr
	if err := buildCmd.Run(); err != nil {
		log.Printf("Build failed: %v", err)
	} else {
		log.Println("Rebuild complete. Waiting for new files...")
	}
}
