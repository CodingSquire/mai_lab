package table

import (
	"fmt"
	"net/http"
	"strings"
	"unicode/utf8"
)

const COLOR_DEFAULT = "\033[39m"
const COLOR_RED = "\033[91m"
const COLOR_GREEN = "\033[92m"
const COLOR_BLUE = "\033[94m"
const COLOR_WHITE = "\033[97m"

const TABLE_TOP = "━"
const TABLE_BOTTOM = "━"
const TABLE_VERTICAL = "┃"
const TABLE_LEFT = "┃"
const TABLE_RIGHT = "┃"
const TABLE_CROSS_LEFT = "┣"
const TABLE_CROSS_RIGHT = "┫"
const TABLE_TOP_LEFT = "┏"
const TABLE_TOP_RIGHT = "┓"
const TABLE_BOTTOM_LEFT = "┗"
const TABLE_BOTTOM_RIGHT = "┛"

const LEFT_PADDING = "\t"

func GetColorByMethod(method string) string {
	switch method {
	case http.MethodGet:
		return string(COLOR_GREEN)
	case http.MethodDelete:
		return string(COLOR_RED)
	case http.MethodPost:
		return string(COLOR_BLUE)
	default:
		return string(COLOR_WHITE)
	}
}

type TableBuilder struct {
	maxLength int
	stringBuilder strings.Builder
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func (b *TableBuilder) AppendRoute(method ,path string) {
	color := GetColorByMethod(method)
	var row strings.Builder

	row.WriteString(LEFT_PADDING)
	row.WriteString(COLOR_DEFAULT)
	row.WriteString(fmt.Sprintf("%-4s", TABLE_LEFT))
	row.WriteString(color)
	row.WriteString(fmt.Sprintf("%-20s", method))
	row.WriteString(COLOR_DEFAULT)
	row.WriteString(fmt.Sprintf("%-4s", TABLE_VERTICAL))
	row.WriteString(COLOR_WHITE)
	row.WriteString(fmt.Sprintf("%-30s", path))
	row.WriteString(COLOR_DEFAULT)
	row.WriteString(TABLE_RIGHT)
	row.WriteString("\n")

	formattedString := row.String()
	stringLength := utf8.RuneCountInString(formattedString) - utf8.RuneCountInString(COLOR_DEFAULT) * 5 - 2

	b.maxLength = Max(stringLength, b.maxLength)
	b.stringBuilder.WriteString(formattedString)
}

func (b *TableBuilder) getTop() string {
	var row strings.Builder

	row.WriteString(LEFT_PADDING)
	row.WriteString(TABLE_TOP_LEFT)	
	row.WriteString(strings.Repeat(TABLE_TOP, b.maxLength - 2))	
	row.WriteString(TABLE_TOP_RIGHT)	
	row.WriteString("\n")

	return row.String()
}

func (b *TableBuilder) getBottom() string {
	var row strings.Builder

	row.WriteString(LEFT_PADDING)
	row.WriteString(TABLE_BOTTOM_LEFT)	
	row.WriteString(strings.Repeat(TABLE_BOTTOM, b.maxLength - 2))	
	row.WriteString(TABLE_BOTTOM_RIGHT)	
	row.WriteString("\n")

	return row.String()
}

func (b * TableBuilder) makeLine(symb string) {
	var row strings.Builder
	
	row.WriteString(LEFT_PADDING)
	row.WriteString(TABLE_LEFT)
	row.WriteString(strings.Repeat(symb, b.maxLength - 2))
	row.WriteString(TABLE_RIGHT)
	row.WriteString("\n")
	
	b.stringBuilder.WriteString(row.String())
}

func (b *TableBuilder) makeEmpty() {
	b.makeLine(" ")
}

func (b *TableBuilder) makeCross() {
	b.makeLine(TABLE_TOP)
}

func (b *TableBuilder) makeMsgLine(message *string) {
	messageFormatted := fmt.Sprintf("  %s", *message)
	messageLength := utf8.RuneCountInString(messageFormatted)

	b.stringBuilder.WriteString(LEFT_PADDING)
	b.stringBuilder.WriteString(TABLE_LEFT)
	b.stringBuilder.WriteString(messageFormatted)
	b.stringBuilder.WriteString(strings.Repeat(" ", b.maxLength - messageLength - 2))
	b.stringBuilder.WriteString(TABLE_RIGHT)
	b.stringBuilder.WriteString("\n")
}

func (b *TableBuilder) AppendLine(message string) {
	b.makeCross()
	//b.makeEmpty()
	b.makeMsgLine(&message)
	//b.makeEmpty()
}

func (b *TableBuilder) PrependLine(message string) {
	oldBuilder := b.stringBuilder
	b.stringBuilder = strings.Builder{}

	//b.makeEmpty()
	b.makeMsgLine(&message)
	//b.makeEmpty()
	b.makeCross()

	b.stringBuilder.WriteString(oldBuilder.String())
}

func (b *TableBuilder) String() string {
	var table strings.Builder

	table.WriteString(b.getTop())
	table.WriteString(b.stringBuilder.String())
	table.WriteString(b.getBottom())

	return table.String()
}
