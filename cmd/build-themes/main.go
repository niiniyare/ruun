package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/niiniyare/ruun"
)

func main() {
	var (
		themesDir = flag.String("themes", "views/style/themes", "Directory containing theme JSON files")
		outputDir = flag.String("output", "static/css/themes", "Output directory for compiled CSS")
		watch     = flag.Bool("watch", false, "Watch for changes and recompile")
		minify    = flag.Bool("minify", true, "Minify generated CSS")
		verbose   = flag.Bool("verbose", false, "Enable verbose output")
	)
	flag.Parse()

	if *verbose {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	// Initialize Theme API
	themeAPI := ruun.NewThemeAPI(*themesDir)

	// Load all themes from directory
	if err := themeAPI.LoadThemesFromDirectory(*themesDir); err != nil {
		log.Fatalf("Failed to load themes from %s: %v", *themesDir, err)
	}

	if *verbose {
		themes := themeAPI.ListThemes()
		fmt.Printf("Loaded %d themes:\n", len(themes))
		for _, theme := range themes {
			fmt.Printf("  - %s (%s)\n", theme.Name, theme.ID)
		}
	}

	// Ensure output directory exists
	if err := os.MkdirAll(*outputDir, 0755); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Compile all themes
	if err := compileThemes(themeAPI, *outputDir, *minify, *verbose); err != nil {
		log.Fatalf("Failed to compile themes: %v", err)
	}

	fmt.Printf("✅ Successfully compiled themes to %s\n", *outputDir)

	// Watch mode
	if *watch {
		fmt.Println("👀 Watching for changes... (Press Ctrl+C to stop)")
		watchThemes(themeAPI, *themesDir, *outputDir, *minify, *verbose)
	}
}

func compileThemes(api *ruun.ThemeAPI, outputDir string, minify, verbose bool) error {
	themes := api.ListThemes()

	for _, theme := range themes {
		if verbose {
			fmt.Printf("Compiling theme: %s\n", theme.Name)
		}

		css, err := api.CompileTheme(theme)
		if err != nil {
			return fmt.Errorf("failed to compile theme %s: %w", theme.ID, err)
		}

		// Write CSS to file
		outputPath := filepath.Join(outputDir, theme.ID+".css")
		if err := os.WriteFile(outputPath, []byte(css), 0644); err != nil {
			return fmt.Errorf("failed to write CSS file %s: %w", outputPath, err)
		}

		if verbose {
			fmt.Printf("  ✅ %s -> %s (%d bytes)\n", theme.ID, outputPath, len(css))
		}
	}

	return nil
}

func watchThemes(api *ruun.ThemeAPI, themesDir, outputDir string, minify, verbose bool) {
	// Simplified watch implementation
	// In a real implementation, use fsnotify or similar
	fmt.Println("Watch mode not fully implemented yet - use manual recompilation")
}