package builder

import (
	"archive/zip"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"image/color"
	"image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/fogleman/gg"
)

// Builder is the main watchface builder
type Builder struct{}

// NewBuilder creates a new builder instance
func NewBuilder() *Builder {
	return &Builder{}
}

// BuildOptions contains options for building a watchface
type BuildOptions struct {
	Name            string   // Watchface name (required)
	Version         string   // Version number
	Author          string   // Author name
	Description     string   // Description
	Tags            []string // Tags
	Template        string   // Template type: simple, analog, digital, custom
	CustomHTML      string   // Custom HTML content (for custom template)
	CustomCSS       string   // Custom CSS content (for custom template)
	CustomJS        string   // Custom JS content (for custom template)
	OutputPath      string   // Output directory
	GeneratePreview bool     // Whether to generate preview image
}

// BuildResult contains the result of a build
type BuildResult struct {
	Success   bool              // Build success
	ZipPath   string            // Path to generated ZIP file
	FileHash  string            // SHA256 hash of ZIP file
	Size      int64             // ZIP file size in bytes
	FileCount int               // Number of files in package
	Files     []string          // List of files in package
	Manifest  string            // manifest.json content
	Error     string            // Error message if failed
	Metadata  map[string]interface{} // Additional metadata
}

// ManifestData represents the manifest.json structure
type ManifestData struct {
	Name        string    `json:"name"`
	Version     string    `json:"version"`
	Author      string    `json:"author"`
	Description string    `json:"description,omitempty"`
	Entrypoint  string    `json:"entrypoint"`
	Tags        []string  `json:"tags,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
}

// Build builds a watchface package
func (b *Builder) Build(options BuildOptions) (*BuildResult, error) {
	// Validate options
	if err := b.validateOptions(options); err != nil {
		return &BuildResult{
			Success: false,
			Error:   err.Error(),
		}, err
	}

	// Create temporary directory
	tempDir, err := os.MkdirTemp("", "watchface-*")
	if err != nil {
		return nil, fmt.Errorf("failed to create temp directory: %w", err)
	}
	defer os.RemoveAll(tempDir)

	// Generate files based on template
	files, err := b.generateTemplateFiles(options)
	if err != nil {
		return nil, fmt.Errorf("failed to generate template files: %w", err)
	}

	// Write files to temp directory
	fileList := []string{}
	for fileName, content := range files {
		filePath := filepath.Join(tempDir, fileName)
		if err := os.WriteFile(filePath, []byte(content), 0644); err != nil {
			return nil, fmt.Errorf("failed to write file %s: %w", fileName, err)
		}
		fileList = append(fileList, fileName)
	}

	// Generate preview image if requested
	if options.GeneratePreview {
		previewPath := filepath.Join(tempDir, "preview.png")
		if err := b.generatePreviewImage(previewPath, options.Name, options.Template); err == nil {
			fileList = append(fileList, "preview.png")
		}
	}

	// Generate manifest.json
	manifest := b.generateManifest(options)
	manifestJSON, _ := json.MarshalIndent(manifest, "", "  ")
	manifestPath := filepath.Join(tempDir, "manifest.json")
	if err := os.WriteFile(manifestPath, manifestJSON, 0644); err != nil {
		return nil, fmt.Errorf("failed to write manifest.json: %w", err)
	}
	fileList = append(fileList, "manifest.json")

	// Create output directory if it doesn't exist
	if err := os.MkdirAll(options.OutputPath, 0755); err != nil {
		return nil, fmt.Errorf("failed to create output directory: %w", err)
	}

	// Create ZIP file
	timestamp := time.Now().Format("20060102_150405")
	zipFileName := fmt.Sprintf("%s_v%s_%s.zip",
		sanitizeFileName(options.Name),
		options.Version,
		timestamp)
	zipPath := filepath.Join(options.OutputPath, zipFileName)

	if err := b.createZip(tempDir, zipPath, fileList); err != nil {
		return nil, fmt.Errorf("failed to create ZIP: %w", err)
	}

	// Calculate file hash
	fileHash, err := calculateFileHash(zipPath)
	if err != nil {
		return nil, fmt.Errorf("failed to calculate file hash: %w", err)
	}

	// Get file size
	fileInfo, _ := os.Stat(zipPath)
	fileSize := fileInfo.Size()

	return &BuildResult{
		Success:   true,
		ZipPath:   zipPath,
		FileHash:  fileHash,
		Size:      fileSize,
		FileCount: len(fileList),
		Files:     fileList,
		Manifest:  string(manifestJSON),
	}, nil
}

// validateOptions validates build options
func (b *Builder) validateOptions(options BuildOptions) error {
	if options.Name == "" {
		return fmt.Errorf("name is required")
	}
	if options.Version == "" {
		options.Version = "1.0.0"
	}
	if options.Template == "" {
		options.Template = "simple"
	}
	validTemplates := map[string]bool{
		"simple": true, "analog": true, "digital": true, "custom": true,
	}
	if !validTemplates[options.Template] {
		return fmt.Errorf("invalid template: %s", options.Template)
	}
	if options.Template == "custom" && options.CustomHTML == "" {
		return fmt.Errorf("custom template requires customHTML")
	}
	return nil
}

// generateTemplateFiles generates template files based on template type
func (b *Builder) generateTemplateFiles(options BuildOptions) (map[string]string, error) {
	switch options.Template {
	case "simple":
		return b.generateSimpleTemplate(options), nil
	case "analog":
		return b.generateAnalogTemplate(options), nil
	case "digital":
		return b.generateDigitalTemplate(options), nil
	case "custom":
		return b.generateCustomTemplate(options), nil
	default:
		return nil, fmt.Errorf("unknown template: %s", options.Template)
	}
}

// generateManifest generates manifest.json
func (b *Builder) generateManifest(options BuildOptions) ManifestData {
	return ManifestData{
		Name:        options.Name,
		Version:     options.Version,
		Author:      options.Author,
		Description: options.Description,
		Entrypoint:  "index.html",
		Tags:        options.Tags,
		CreatedAt:   time.Now(),
	}
}

// generatePreviewImage generates a preview image
func (b *Builder) generatePreviewImage(outputPath, name, template string) error {
	const width = 512
	const height = 512

	dc := gg.NewContext(width, height)

	// Background gradient based on template
	var bgColor1, bgColor2 color.Color
	switch template {
	case "analog":
		bgColor1 = color.RGBA{245, 245, 245, 255}
		bgColor2 = color.RGBA{220, 220, 220, 255}
	case "digital":
		bgColor1 = color.RGBA{10, 10, 30, 255}
		bgColor2 = color.RGBA{30, 30, 60, 255}
	default: // simple
		bgColor1 = color.RGBA{74, 144, 226, 255}
		bgColor2 = color.RGBA{142, 84, 233, 255}
	}

	// Draw gradient background
	for y := 0; y < height; y++ {
		r1, g1, b1, a1 := bgColor1.RGBA()
		r2, g2, b2, a2 := bgColor2.RGBA()
		ratio := float64(y) / float64(height)
		r := uint8((float64(r1)*(1-ratio) + float64(r2)*ratio) / 256)
		g := uint8((float64(g1)*(1-ratio) + float64(g2)*ratio) / 256)
		b := uint8((float64(b1)*(1-ratio) + float64(b2)*ratio) / 256)
		a := uint8((float64(a1)*(1-ratio) + float64(a2)*ratio) / 256)
		dc.SetRGB255(int(r), int(g), int(b))
		dc.SetRGBA255(int(r), int(g), int(b), int(a))
		dc.DrawLine(0, float64(y), float64(width), float64(y))
		dc.Stroke()
	}

	// Draw text
	dc.SetRGB(1, 1, 1)
	dc.DrawStringAnchored(name, float64(width)/2, float64(height)/2, 0.5, 0.5)

	// Save image
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	img := dc.Image()
	return png.Encode(file, img)
}

// createZip creates a ZIP file from a directory
func (b *Builder) createZip(sourceDir, zipPath string, files []string) error {
	zipFile, err := os.Create(zipPath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, fileName := range files {
		filePath := filepath.Join(sourceDir, fileName)
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", fileName, err)
		}

		zipFileWriter, err := zipWriter.Create(fileName)
		if err != nil {
			return fmt.Errorf("failed to create zip entry %s: %w", fileName, err)
		}

		if _, err := zipFileWriter.Write(fileContent); err != nil {
			return fmt.Errorf("failed to write zip entry %s: %w", fileName, err)
		}
	}

	return nil
}

// sanitizeFileName removes invalid characters from filename
func sanitizeFileName(name string) string {
	// Replace spaces with underscores
	name = strings.ReplaceAll(name, " ", "_")
	// Remove invalid characters
	invalid := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"}
	for _, char := range invalid {
		name = strings.ReplaceAll(name, char, "")
	}
	return name
}

// calculateFileHash calculates SHA256 hash of a file
func calculateFileHash(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
