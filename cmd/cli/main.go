package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/ziztechnology/WatchfaceBuilder/pkg/builder"
)

var (
	name           string
	version        string
	author         string
	description    string
	template       string
	tags           string
	output         string
	noPreview      bool
	interactive    bool
	listTemplates  bool
	customHTML     string
	customCSS      string
	customJS       string
	customHTMLFile string
	customCSSFile  string
	customJSFile   string
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "watchface-builder",
		Short: "A tool to quickly generate H5 watchface packages",
		Long: `Watchface Builder - 表盘包创建器

A lazy-friendly tool to quickly generate H5 watchface packages for WebView-based smartwatches.
为 WebView 套壳智能手表快速生成 H5 表盘包的懒人工具。

Examples:
  # Interactive mode (recommended for beginners)
  watchface-builder -i

  # Quick build
  watchface-builder -name "My Watchface" -template analog

  # With all options
  watchface-builder -name "My Watchface" -template digital \
    -version 1.0.0 -author "Your Name" -description "A cool watchface"

  # List available templates
  watchface-builder -list`,
		Run: runBuild,
	}

	rootCmd.Flags().StringVarP(&name, "name", "n", "", "Watchface name (required)")
	rootCmd.Flags().StringVarP(&version, "version", "v", "1.0.0", "Version number")
	rootCmd.Flags().StringVarP(&author, "author", "a", "Anonymous", "Author name")
	rootCmd.Flags().StringVarP(&description, "description", "d", "", "Watchface description")
	rootCmd.Flags().StringVarP(&template, "template", "t", "simple", "Template type: simple, analog, digital, custom")
	rootCmd.Flags().StringVar(&tags, "tags", "", "Tags, comma-separated")
	rootCmd.Flags().StringVarP(&output, "output", "o", ".", "Output directory")
	rootCmd.Flags().BoolVar(&noPreview, "no-preview", false, "Do not generate preview image")
	rootCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "Interactive mode")
	rootCmd.Flags().BoolVarP(&listTemplates, "list", "l", false, "List available templates")
	rootCmd.Flags().StringVar(&customHTML, "custom-html", "", "Custom HTML content")
	rootCmd.Flags().StringVar(&customCSS, "custom-css", "", "Custom CSS content")
	rootCmd.Flags().StringVar(&customJS, "custom-js", "", "Custom JS content")
	rootCmd.Flags().StringVar(&customHTMLFile, "custom-html-file", "", "Custom HTML file path")
	rootCmd.Flags().StringVar(&customCSSFile, "custom-css-file", "", "Custom CSS file path")
	rootCmd.Flags().StringVar(&customJSFile, "custom-js-file", "", "Custom JS file path")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func runBuild(cmd *cobra.Command, args []string) {
	printBanner()

	// List templates mode
	if listTemplates {
		printTemplateList()
		return
	}

	// Interactive mode
	if interactive {
		runInteractive()
		return
	}

	// Validate required parameters
	if name == "" {
		fmt.Println("❌ Error: Watchface name is required")
		fmt.Println()
		fmt.Println("Usage:")
		cmd.Help()
		fmt.Println()
		fmt.Println("Quick start:")
		fmt.Println("  watchface-builder -name \"My Watchface\"")
		fmt.Println("  watchface-builder -i  # Interactive mode")
		fmt.Println()
		os.Exit(1)
	}

	// Build watchface
	buildWatchface()
}

func buildWatchface() {
	// Parse tags
	var tagList []string
	if tags != "" {
		tagList = strings.Split(tags, ",")
		for i := range tagList {
			tagList[i] = strings.TrimSpace(tagList[i])
		}
	}

	// Read custom files if specified
	if customHTMLFile != "" {
		content, err := os.ReadFile(customHTMLFile)
		if err != nil {
			fmt.Printf("❌ Failed to read custom HTML file: %v\n", err)
			os.Exit(1)
		}
		customHTML = string(content)
	}
	if customCSSFile != "" {
		content, err := os.ReadFile(customCSSFile)
		if err != nil {
			fmt.Printf("❌ Failed to read custom CSS file: %v\n", err)
			os.Exit(1)
		}
		customCSS = string(content)
	}
	if customJSFile != "" {
		content, err := os.ReadFile(customJSFile)
		if err != nil {
			fmt.Printf("❌ Failed to read custom JS file: %v\n", err)
			os.Exit(1)
		}
		customJS = string(content)
	}

	// Create build options
	options := builder.BuildOptions{
		Name:            name,
		Version:         version,
		Author:          author,
		Description:     description,
		Template:        template,
		Tags:            tagList,
		OutputPath:      output,
		GeneratePreview: !noPreview,
		CustomHTML:      customHTML,
		CustomCSS:       customCSS,
		CustomJS:        customJS,
	}

	// Create builder
	b := builder.NewBuilder()

	fmt.Println("🔨 Building watchface package...")
	fmt.Println()

	// Execute build
	result, err := b.Build(options)
	if err != nil {
		fmt.Printf("❌ Build failed: %v\n", err)
		os.Exit(1)
	}

	// Print result
	printResult(result)
}

func runInteractive() {
	fmt.Println("🎯 Interactive Mode")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	// Name
	fmt.Print("Watchface name (required): ")
	scanner.Scan()
	name = scanner.Text()
	if name == "" {
		fmt.Println("❌ Name is required")
		os.Exit(1)
	}

	// Version
	fmt.Print("Version [1.0.0]: ")
	scanner.Scan()
	version = scanner.Text()
	if version == "" {
		version = "1.0.0"
	}

	// Author
	fmt.Print("Author [Anonymous]: ")
	scanner.Scan()
	author = scanner.Text()
	if author == "" {
		author = "Anonymous"
	}

	// Description
	fmt.Print("Description (optional): ")
	scanner.Scan()
	description = scanner.Text()

	// Template selection
	fmt.Println()
	fmt.Println("Select template:")
	fmt.Println("  1. Simple   - Minimalist digital clock")
	fmt.Println("  2. Analog   - Classic clock with hands")
	fmt.Println("  3. Digital  - Tech-style digital clock")
	fmt.Println("  4. Custom   - Custom HTML/CSS/JS")
	fmt.Print("Enter option [1]: ")
	scanner.Scan()
	choice := scanner.Text()
	switch choice {
	case "2":
		template = "analog"
	case "3":
		template = "digital"
	case "4":
		template = "custom"
		fmt.Println()
		fmt.Println("⚠️  Custom template requires HTML content")
		fmt.Print("HTML file path: ")
		scanner.Scan()
		customHTMLFile = scanner.Text()
		fmt.Print("CSS file path (optional): ")
		scanner.Scan()
		customCSSFile = scanner.Text()
		fmt.Print("JS file path (optional): ")
		scanner.Scan()
		customJSFile = scanner.Text()
	default:
		template = "simple"
	}

	// Tags
	fmt.Print("Tags (comma-separated, optional): ")
	scanner.Scan()
	tags = scanner.Text()

	// Preview
	fmt.Print("Generate preview image? [Y/n]: ")
	scanner.Scan()
	previewChoice := strings.ToLower(scanner.Text())
	noPreview = previewChoice == "n"

	fmt.Println()
	fmt.Println("🔨 Building...")
	fmt.Println()

	// Build
	buildWatchface()
}

func printBanner() {
	banner := `
╔═══════════════════════════════════════════════════════╗
║   Watchface Builder - 表盘包创建器                    ║
║   Quickly generate H5 watchface packages              ║
║   快速生成 H5 表盘包                                  ║
╚═══════════════════════════════════════════════════════╝
`
	fmt.Println(banner)
}

func printTemplateList() {
	fmt.Println("📋 Available Templates:")
	fmt.Println()
	fmt.Println("  1. simple")
	fmt.Println("     └─ Minimalist digital clock with gradient background")
	fmt.Println("     └─ 简洁的数字时钟，带渐变背景")
	fmt.Println()
	fmt.Println("  2. analog")
	fmt.Println("     └─ Classic analog clock with Canvas rendering")
	fmt.Println("     └─ 经典指针时钟，Canvas 绘制")
	fmt.Println()
	fmt.Println("  3. digital")
	fmt.Println("     └─ Tech-style digital clock with neon effects")
	fmt.Println("     └─ 科技感数字时钟，霓虹灯效果")
	fmt.Println()
	fmt.Println("  4. custom")
	fmt.Println("     └─ Fully customizable with your own HTML/CSS/JS")
	fmt.Println("     └─ 使用自定义 HTML/CSS/JS 完全自定义")
	fmt.Println()
	fmt.Println("Usage example:")
	fmt.Println("  watchface-builder -name \"My Watchface\" -template analog")
	fmt.Println()
}

func printResult(result *builder.BuildResult) {
	fmt.Println("✅ Build successful!")
	fmt.Println()
	fmt.Println("📦 Output:")
	fmt.Printf("  File path: %s\n", result.ZipPath)
	fmt.Printf("  File size: %.2f KB\n", float64(result.Size)/1024)
	fmt.Printf("  File hash: %s\n", result.FileHash)
	fmt.Printf("  File count: %d\n", result.FileCount)
	fmt.Println()
	fmt.Println("📋 Files included:")
	for _, file := range result.Files {
		fmt.Printf("  ✓ %s\n", file)
	}
	fmt.Println()
	fmt.Println("📄 Manifest:")
	var manifest map[string]interface{}
	if err := json.Unmarshal([]byte(result.Manifest), &manifest); err == nil {
		manifestJSON, _ := json.MarshalIndent(manifest, "  ", "  ")
		fmt.Println(string(manifestJSON))
	}
	fmt.Println()
	fmt.Println("🎉 Done! You can now upload and validate this watchface package.")
	fmt.Println()
	fmt.Println("Next steps:")
	fmt.Println("  1. Test locally: unzip", result.ZipPath)
	fmt.Println("  2. Validate: Upload to validation service")
	fmt.Println()
}
