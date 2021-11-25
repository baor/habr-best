package habr

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStripTags_PositiveCases(t *testing.T) {
	tests := []struct {
		input          string
		expectedOutput string
	}{
		{input: `notags`,
			expectedOutput: `notags`},
		{input: `<br>text<br>`,
			expectedOutput: ` text `},
		{input: `<br/>text<br/>`,
			expectedOutput: ` text `},
		{input: `<b>text</b>`,
			expectedOutput: ` text`},
		{input: `<b>text1 <a arg="aaa">text2</a> text3</b>`,
			expectedOutput: ` text1 <a arg="aaa">text2</a> text3`},
		{input: `<b>text1</b> text2 <b>text3</b>`,
			expectedOutput: ` text1 text2  text3`},
		{input: `<img>text`,
			expectedOutput: ` text`},
		{input: `<img src="" size="">text`,
			expectedOutput: ` text`},
		{input: `<a>text1</a> text2 <z>text3</z>`,
			expectedOutput: `<a>text1</a> text2  text3`},
		{input: `<p>ого <a href="https://audio-v-text.silero.ai/a%3BRedis%3BRocksDB"> text </a> Читать дальше &rarr; ого —&nbsp;ого</p>`,
			expectedOutput: " ого <a href=\"https://audio-v-text.silero.ai/a%3BRedis%3BRocksDB\"> text </a> Читать дальше → ого —\u00a0ого"},
		{input: `<img src="https://habrastorage.org/webt/xa/j5/jm/xaj5jmbdokbljvrnijpts6xlkoi.jpeg">WWW!`,
			expectedOutput: ` WWW!`},
	}

	for _, tc := range tests {
		outputActual, err := stripTags(tc.input)
		assert.Nil(t, err)
		assert.Equal(t, tc.expectedOutput, outputActual)
	}
}

func TestGetPostID(t *testing.T) {
	input := `https://habr.com/post/413925/?utm_source=habrahabr&utm_medium=rss&utm_campaign=413925`
	outputExpected := "413925"
	outputActual := getPostID(input)
	assert.Equal(t, outputExpected, outputActual)
}

// func TestManual(t *testing.T) {
// 	c := NewHabrReader()
// 	feeds := c.GetBestFeed()
// 	assert.Greater(t, len(feeds), 1)
// }
