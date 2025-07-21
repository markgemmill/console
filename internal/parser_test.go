package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ParseTestData struct {
	TaggedText     string
	ExpectedTagged string
	ExpectedRaw    string
}

func NewParseTestData(input, tagged, raw string) ParseTestData {
	return ParseTestData{
		TaggedText:     input,
		ExpectedTagged: tagged,
		ExpectedRaw:    raw,
	}
}

// "I'm going [yellow]HOME[/yellow] to mother!",

var testData = []ParseTestData{
	NewParseTestData(
		"There are no tags in this text.",
		"There are no tags in this text.",
		"There are no tags in this text.",
	),
	NewParseTestData(
		"I'm going [yellow]HOME[/yellow] to mother!",
		"I'm going <yellow>HOME</yellow> to mother!",
		"I'm going HOME to mother!",
	),
	NewParseTestData(
		"Hello [bold]world[/bold]!",
		"Hello <bold>world</bold>!",
		"Hello world!",
	),
	NewParseTestData(
		"[red]Start[/red] middle [blue]end[/blue]",
		"<red>Start</red> middle <blue>end</blue>",
		"Start middle end",
	),
	NewParseTestData(
		"No tags here",
		"No tags here",
		"No tags here",
	),
	NewParseTestData(
		"[green]Only tagged[/green]",
		"<green>Only tagged</green>",
		"Only tagged",
	),
	NewParseTestData(
		"Multiple [red]red[/red] and [blue]blue[/blue] tags [yellow]here[/yellow]!",
		"Multiple <red>red</red> and <blue>blue</blue> tags <yellow>here</yellow>!",
		"Multiple red and blue tags here!",
	),
	NewParseTestData(
		"I'm [yellow]going to the moon[/yellow], wanna come with? Text me...",
		"I'm <yellow>going to the moon</yellow>, wanna come with? Text me...",
		"I'm going to the moon, wanna come with? Text me...",
	),
	NewParseTestData(
		"[yellow]DO THIS[/yellow]: [green]the thing to do[/green]!!!",
		"<yellow>DO THIS</yellow>: <green>the thing to do</green>!!!",
		"DO THIS: the thing to do!!!",
	),
	NewParseTestData(
		"DO: [yellow]%s[/yellow]\n",
		"DO: <yellow>%s</yellow>\n",
		"DO: %s\n",
	),
	NewParseTestData(
		"",
		"",
		"",
	),
}

func TestParsing(t *testing.T) {
	for _, data := range testData {

		parser := NewMessageParser(data.TaggedText)
		err := parser.Parse()

		assert.Nil(t, err)

		assert.Equal(t, data.ExpectedTagged, parser.TaggedString())
		assert.Equal(t, data.ExpectedRaw, parser.RawString())
	}
}
