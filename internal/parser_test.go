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
		"[green]OK[/green] [yellow]%s[/yellow]\n",
		"<green>OK</green> <yellow>%s</yellow>\n",
		"OK %s\n",
	),
	NewParseTestData(
		"",
		"",
		"",
	),
	NewParseTestData(
		"     [yellow]right justified[/yellow] [green]%s[/green]\n",
		"     <yellow>right justified</yellow> <green>%s</green>\n",
		"     right justified %s\n",
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

func TestParsingDetails(t *testing.T) {
	msg := "    [yellow]name[/yellow]  [green]of the[/green]  [magenta]rose[/magenta] !"
	parser := NewMessageParser(msg)
	err := parser.Parse()

	assert.Nil(t, err)

	assert.Equal(t, 7, len(parser.segments))
	assert.Equal(t, "    ", parser.segments[0].Text)
	assert.Equal(t, "name", parser.segments[1].Text)
	assert.Equal(t, "  ", parser.segments[2].Text)
	assert.Equal(t, "of the", parser.segments[3].Text)
	assert.Equal(t, "  ", parser.segments[4].Text)
	assert.Equal(t, "rose", parser.segments[5].Text)
	assert.Equal(t, " !", parser.segments[6].Text)
}
