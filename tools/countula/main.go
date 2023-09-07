package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	// Step 1: Define flags
	pathPtr := flag.String("path", ".", "path to the folder (default: current folder)")
	extensionsPtr := flag.String("extensions", "", "file extensions to include (comma separated)")
	excludesPtr := flag.String("excludes", "", "paths to exclude (comma separated)")
	gitignorePtr := flag.Bool("gitignore", true, "whether to use .gitignore (default: true)")
	dirsPtr := flag.Bool("dirs", false, "show by dir (default: false)")

	// x, y := "../../../TGD", "json"
	// extensionsPtr = &y
	// pathPtr = &x

	flag.Parse()

	extensions := strings.Split(*extensionsPtr, ",")
	excludes := strings.Split(*excludesPtr, ",")

	gitIgnorePatterns := []string{}
	if *gitignorePtr {
		gitIgnorePath := filepath.Join(*pathPtr, ".gitignore")
		// Step 2: Read .gitignore and store patterns in a slice
		file, err := os.Open(gitIgnorePath)
		if err == nil {
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				line := strings.TrimSpace(scanner.Text())
				if len(line) > 0 && !strings.HasPrefix(line, "#") {
					gitIgnorePatterns = append(gitIgnorePatterns, line)
				}
			}
			file.Close()
		}
	}

	countByExtension := make(map[string]int)
	countBySubfolder := make(map[string]int)

	err := filepath.Walk(*pathPtr, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Step 3: Skip .git folder and files matching .gitignore patterns
		if strings.HasSuffix(path, ".git") || matchesGitIgnore(path, gitIgnorePatterns) || matchesExcludes(path, excludes) {
			if info.IsDir() {
				return filepath.SkipDir
			}
			return nil
		}

		if !matchesExtensions(path, extensions) {
			return nil
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			regex := regexp.MustCompile(`\S`)
			for scanner.Scan() {
				line := scanner.Text()
				if len(regex.FindAllString(line, -1)) > 2 {
					extension := filepath.Ext(path)
					subfolder := filepath.Dir(path)
					countByExtension[strings.TrimPrefix(extension, ".")]++
					countBySubfolder[subfolder]++
				}
			}

			if err := scanner.Err(); err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var total int
	fmt.Println("Count by extension:")
	for _, ext := range extensions {
		count := countByExtension[strings.TrimPrefix(ext, ".")]
		fmt.Printf("  %s: %d\n", ext, count)
		total += count
	}
	fmt.Println("TOTAL:", total)

	if *dirsPtr {
		fmt.Println("\nCount by subfolder:")
		for folder, count := range countBySubfolder {
			fmt.Printf("%s: %d\n", folder, count)
		}
	}
}

func matchesGitIgnore(path string, patterns []string) bool {
	for _, pattern := range patterns {
		matched, err := filepath.Match(pattern, filepath.Base(path))
		if err == nil && matched {
			return true
		}
	}
	return false
}

func matchesExcludes(path string, excludes []string) bool {
	for _, exclude := range excludes {
		if len(exclude) != 0 && strings.Contains(path, exclude) {
			return true
		}
	}
	return false
}

func matchesExtensions(path string, extensions []string) bool {
	if len(extensions) == 0 || (len(extensions) == 1 && extensions[0] == "") {
		return true
	}
	ext := filepath.Ext(path)
	for _, extension := range extensions {
		if !strings.HasPrefix(extension, ".") {
			extension = "." + extension
		}
		if ext == extension {
			return true
		}
	}
	return false
}
