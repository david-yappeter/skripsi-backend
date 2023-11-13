package util

import (
	"testing"
)

type lessTextTest struct {
	texts  []string
	length int
	output string
}

var lessTextTests = []lessTextTest{
	{
		texts:  []string{},
		length: 0,
		output: "",
	},
	{
		texts:  []string{},
		length: 100,
		output: "",
	},
	{
		texts:  []string{"Q1", "Q2", "Q3", "Q4", "Q5", "Q6", "Q7"},
		length: -1,
		output: "Q1, Q2, Q3, Q4, Q5, Q6 and Q7",
	},
	{
		texts:  []string{"Q1"},
		length: 1,
		output: "Q1",
	},
	{
		texts:  []string{"Q1"},
		length: 2,
		output: "Q1",
	},
	{
		texts:  []string{"Q1"},
		length: 3,
		output: "Q1",
	},
	{
		texts:  []string{"Q1", "Q2"},
		length: 1,
		output: "Q1 and 1 more",
	},
	{
		texts:  []string{"Q1", "Q2"},
		length: 2,
		output: "Q1 and Q2",
	},
	{
		texts:  []string{"Q1", "Q2", "Q3", "Q4", "Q5", "Q6", "Q7"},
		length: 2,
		output: "Q1, Q2 and 5 more",
	},
}

func TestLessText(t *testing.T) {
	for i, ltt := range lessTextTests {
		if output := LessText(ltt.texts, ltt.length); output != ltt.output {
			t.Errorf("test %d: have \"%s\" want \"%s\"\n", i, output, ltt.output)
		}
	}
}
