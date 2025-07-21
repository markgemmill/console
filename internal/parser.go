package internal

import (
	"fmt"
	"regexp"
	"strings"
)

var tagRegex = regexp.MustCompile(`\[([a-z]+)\](.*?)\[/[a-z]+\]`)

// TextSegment represents a segment of text with optional tag information
type TextSegment struct {
	Text    string
	TagName string // empty for untagged text
	IsTag   bool   // true if this segment is inside a tag
}

// TagParser handles parsing of bracketed tags
type MessageParser struct {
	segments []TextSegment
	src      string
}

// NewMessageParser creates a new tag parser
func NewMessageParser(message string) *MessageParser {
	return &MessageParser{
		src:      message,
		segments: []TextSegment{},
	}
}

func (p *MessageParser) addTextSegment(segment TextSegment) {
	p.segments = append(p.segments, segment)
}

// Parse parses the input text and returns a slice of TextSegments
func (p *MessageParser) Parse() error {
	lastEnd := 0

	input := p.src

	// Find all tag matches
	matches := tagRegex.FindAllStringSubmatchIndex(input, -1)

	for _, match := range matches {
		start := match[0]        // Start of entire match
		end := match[1]          // End of entire match
		tagStart := match[2]     // Start of tag name
		tagEnd := match[3]       // End of tag name
		contentStart := match[4] // Start of content
		contentEnd := match[5]   // End of content

		// Add text before the tag (if any)
		if start > lastEnd {
			beforeText := input[lastEnd:start]
			if strings.TrimSpace(beforeText) != "" {
				p.addTextSegment(TextSegment{
					Text:    beforeText,
					TagName: "",
					IsTag:   false,
				})
			}
		}

		// Add the tagged content
		tagName := input[tagStart:tagEnd]
		content := input[contentStart:contentEnd]
		p.addTextSegment(TextSegment{
			Text:    content,
			TagName: tagName,
			IsTag:   true,
		})

		lastEnd = end
	}

	// Add any remaining text after the last tag
	if lastEnd < len(input) {
		afterText := input[lastEnd:]
		p.addTextSegment(TextSegment{
			Text:    afterText,
			TagName: "",
			IsTag:   false,
		})
	}

	// Handle case where no tags were found
	if len(p.segments) == 0 {
		p.addTextSegment(TextSegment{
			Text:    input,
			TagName: "",
			IsTag:   false,
		})
	}

	return nil
}

// SrcString returns the original message source.
func (p *MessageParser) SrcString() string {
	return p.src
}

// RawString extracts just the text content, removing all tags
func (p *MessageParser) RawString() string {
	var result strings.Builder
	for _, segment := range p.segments {
		result.WriteString(segment.Text)
	}
	return result.String()
}

// TaggedString returns text with tag names preserved as angle brackets.
func (p *MessageParser) TaggedString() string {
	var result strings.Builder
	for _, segment := range p.segments {
		if segment.IsTag {
			result.WriteString(fmt.Sprintf("<%s>%s</%s>",
				segment.TagName, segment.Text, segment.TagName))
		} else {
			result.WriteString(segment.Text)
		}
	}
	return result.String()
}

// TaggedText returns text with tag names preserved as angle brackets.
func (p *MessageParser) String() string {
	var result strings.Builder
	for _, segment := range p.segments {
		if segment.IsTag {
			colorize, okay := colors[segment.TagName]
			if okay {
				result.WriteString(colorize(segment.Text))
				continue
			}
		}
		result.WriteString(segment.Text)
	}
	return result.String()
}

var labelRx = regexp.MustCompile(`\[/?[a-z]+\]`)

func RequiresParsing(text string) bool {
	matches := labelRx.FindAllString(text, -1)
	if matches == nil {
		return false
	}
	if len(matches) <= 1 {
		return false
	}
	return true
}
