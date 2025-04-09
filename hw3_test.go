package main

import (
	"testing"
)

// Test PrepareData method of TextDocument
func TestTextDocument_PrepareData(t *testing.T) {
	td := &TextDocument{}
	expected := "This is the raw text data."
	actual := td.PrepareData()
	if actual != expected {
		t.Errorf("TestTextDocument_PrepareData failed: expected '%s', got '%s'", expected, actual)
	} else {
		t.Logf("TestTextDocument_PrepareData passed: expected '%s', got '%s'", expected, actual)
	}
}

// Test FormatContent method of TextDocument
func TestTextDocument_FormatContent(t *testing.T) {
	td := &TextDocument{}
	data := "test data"
	expected := "Formatted Text: test data"
	actual := td.FormatContent(data)
	if actual != expected {
		t.Errorf("TestTextDocument_FormatContent failed: expected '%s', got '%s'", expected, actual)
	} else {
		t.Logf("TestTextDocument_FormatContent passed: expected '%s', got '%s'", expected, actual)
	}
}

// Test Save method of TextDocument
func TestTextDocument_Save(t *testing.T) {
	td := &TextDocument{}
	content := "Formatted Text: test data"
	expected := "Saving text document: Formatted Text: test data"
	actual := td.Save(content)
	if actual != expected {
		t.Errorf("TestTextDocument_Save failed: expected '%s', got '%s'", expected, actual)
	} else {
		t.Logf("TestTextDocument_Save passed: expected '%s', got '%s'", expected, actual)
	}
}

// Test PrepareData method of HTMLDocument
func TestHTMLDocument_PrepareData(t *testing.T) {
	hd := &HTMLDocument{}
	expected := "<html><body>This is raw HTML data.</body></html>"
	actual := hd.PrepareData()
	if actual != expected {
		t.Errorf("TestHTMLDocument_PrepareData failed: expected '%s', got '%s'", expected, actual)
	} else {
		t.Logf("TestHTMLDocument_PrepareData passed: expected '%s', got '%s'", expected, actual)
	}
}

// Test FormatContent method of HTMLDocument
func TestHTMLDocument_FormatContent(t *testing.T) {
	hd := &HTMLDocument{}
	data := "html test data"
	expected := "<div>html test data</div>"
	actual := hd.FormatContent(data)
	if actual != expected {
		t.Errorf("TestHTMLDocument_FormatContent failed: expected '%s', got '%s'", expected, actual)
	} else {
		t.Logf("TestHTMLDocument_FormatContent passed: expected '%s', got '%s'", expected, actual)
	}
}

// Test Save method of HTMLDocument
func TestHTMLDocument_Save(t *testing.T) {
	hd := &HTMLDocument{}
	content := "<div>formatted html</div>"
	expected := "Saving HTML document: <div>formatted html</div>"
	actual := hd.Save(content)
	if actual != expected {
		t.Errorf("TestHTMLDocument_Save failed: expected '%s', got '%s'", expected, actual)
	} else {
		t.Logf("TestHTMLDocument_Save passed: expected '%s', got '%s'", expected, actual)
	}
}

// Test BaseGenerator Generate method with TextDocument
func TestBaseGenerator_Generate_TextDocument(t *testing.T) {
	textDoc := &TextDocument{}
	baseGenerator := &BaseGenerator{generator: textDoc}
	expectedOutput := "Saving text document: Formatted Text: This is the raw text data."
	output := baseGenerator.Generate()
	if output != expectedOutput {
		t.Errorf("TestBaseGenerator_Generate_TextDocument failed: expected '%s', got '%s'", expectedOutput, output)
	} else {
		t.Logf("TestBaseGenerator_Generate_TextDocument passed: expected '%s', got '%s'", expectedOutput, output)
	}
}

// Test BaseGenerator Generate method with HTMLDocument
func TestBaseGenerator_Generate_HTMLDocument(t *testing.T) {
	htmlDoc := &HTMLDocument{}
	baseGenerator := &BaseGenerator{generator: htmlDoc}
	expectedOutput := "Saving HTML document: <div><html><body>This is raw HTML data.</body></html></div>"
	output := baseGenerator.Generate()
	if output != expectedOutput {
		t.Errorf("TestBaseGenerator_Generate_HTMLDocument failed: expected '%s', got '%s'", expectedOutput, output)
	} else {
		t.Logf("TestBaseGenerator_Generate_HTMLDocument passed: expected '%s', got '%s'", expectedOutput, output)
	}
}
