package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// DocumentGenerator interface defines the basic steps for document generation.
// Each document type must implement these methods.
type DocumentGenerator interface {
	PrepareData() string              // Prepare the raw content of the document
	FormatContent(data string) string // Format the content
	Save(content string) string       // Save the formatted content
}

// BaseGenerator class is a base document generator.
// It contains the Generate() method, which executes the three steps defined in the DocumentGenerator interface sequentially.
// The Generate() method itself does not implement specific logic; instead, it relies on the document type's implementation of the DocumentGenerator interface.
type BaseGenerator struct {
	generator DocumentGenerator
}

// Generate method is the template method. It defines the sequence of steps for document generation,
// and delegates the specific implementation of each step to the underlying DocumentGenerator instance.
func (bg *BaseGenerator) Generate() string {
	data := bg.generator.PrepareData()
	formattedContent := bg.generator.FormatContent(data)
	return bg.generator.Save(formattedContent)
}

// TextDocument class implements the logic for generating plain text documents.
type TextDocument struct{}

// PrepareData implements the logic for preparing plain text data.
func (td *TextDocument) PrepareData() string {
	return "This is the raw text data."
}

// FormatContent implements the logic for formatting plain text content.
func (td *TextDocument) FormatContent(data string) string {
	return fmt.Sprintf("Formatted Text: %s", data)
}

// Save implements the logic for saving plain text documents.
func (td *TextDocument) Save(content string) string {
	return fmt.Sprintf("Saving text document: %s", content)
}

// HTMLDocument class implements the logic for generating HTML documents.
type HTMLDocument struct{}

// PrepareData implements the logic for preparing HTML data.
func (hd *HTMLDocument) PrepareData() string {
	return "<html><body>This is raw HTML data.</body></html>"
}

// FormatContent implements the logic for formatting HTML content.
func (hd *HTMLDocument) FormatContent(data string) string {
	return fmt.Sprintf("<div>%s</div>", data)
}

// Save implements the logic for saving HTML documents.
func (hd *HTMLDocument) Save(content string) string {
	return fmt.Sprintf("Saving HTML document: %s", content)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Summit froms of text：")
	fmt.Println("0: Plain Text")
	fmt.Println("1: HTML Text")

	fmt.Print("Please enter 0 or 1: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var generator *BaseGenerator

	switch input {
	case "0":
		generator = &BaseGenerator{generator: &TextDocument{}}
	case "1":
		generator = &BaseGenerator{generator: &HTMLDocument{}}
	default:
		fmt.Println("Not available option.\n Goodbye.")
		return
	}

	if generator != nil {
		output := generator.Generate()
		fmt.Println("\nGenerating document...")
		fmt.Println(output)
	}
}
