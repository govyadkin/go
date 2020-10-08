package unique

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestUnique(t *testing.T) {
	input := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		" ",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}

	op := Options{
		Count:     false,
		Duplicate: false,
		Unique:    false,
		Fields:    false,
		NumFields: 0,
		Chars:     false,
		NumChars:  0,
		Register:  false,
	}
	output, err := FindUnique(input, op)
	if err != nil {
		t.Error(err)
	}
	answer := []string{
		" ",
		"I love music of Kartik.",
		"I love music.",
		"Thanks.",
	}
	require.Equal(t, answer, output, "TestUnique: The two slice should be the same.")
}

func TestUniqueCount(t *testing.T) {
	input := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		" ",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}

	op := Options{
		Count:     true,
		Duplicate: false,
		Unique:    false,
		Fields:    false,
		NumFields: 0,
		Chars:     false,
		NumChars:  0,
		Register:  false,
	}
	output, err := FindUnique(input, op)
	if err != nil {
		t.Error(err)
	}
	answer := []string{
		"1  ",
		"1 Thanks.",
		"2 I love music of Kartik.",
		"3 I love music.",
	}
	require.Equal(t, answer, output, "TestUniqueCount: The two slice should be the same.")
}

func TestUniqueD(t *testing.T) {
	input := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		" ",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}

	op := Options{
		Count:     false,
		Duplicate: true,
		Unique:    false,
		Fields:    false,
		NumFields: 0,
		Chars:     false,
		NumChars:  0,
		Register:  false,
	}
	output, err := FindUnique(input, op)
	if err != nil {
		t.Error(err)
	}
	answer := []string{
		"I love music of Kartik.",
		"I love music.",
	}
	require.Equal(t, answer, output, "TestUniqueD: The two slice should be the same.")
}

func TestUniqueU(t *testing.T) {
	input := []string{
		"I love music.",
		"I love music.",
		"I love music.",
		" ",
		"I love music of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}

	op := Options{
		Count:     false,
		Duplicate: false,
		Unique:    true,
		Fields:    false,
		NumFields: 0,
		Chars:     false,
		NumChars:  0,
		Register:  false,
	}
	output, err := FindUnique(input, op)
	if err != nil {
		t.Error(err)
	}
	answer := []string{
		" ",
		"Thanks.",
	}
	require.Equal(t, answer, output, "TestUniqueU: The two slice should be the same.")
}

func TestUniqueI(t *testing.T) {
	input := []string{
		"I LOVE MUSIC.",
		"I love music.",
		"I LoVe MuSiC.",
		" ",
		"I love MuSIc of Kartik.",
		"I love music of Kartik.",
		"Thanks.",
	}

	op := Options{
		Count:     false,
		Duplicate: false,
		Unique:    false,
		Fields:    false,
		NumFields: 0,
		Chars:     false,
		NumChars:  0,
		Register:  true,
	}
	output, err := FindUnique(input, op)
	if err != nil {
		t.Error(err)
	}
	answer := []string{
		" ",
		"I LOVE MUSIC.",
		"I love MuSIc of Kartik.",
		"Thanks.",
	}
	require.Equal(t, answer, output, "TestUniqueI: The two slice should be the same.")
}

func TestUniqueF(t *testing.T) {
	input := []string{
		"We love music.",
		"I love music.",
		"They love music.",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}

	op := Options{
		Count:     false,
		Duplicate: false,
		Unique:    false,
		Fields:    true,
		NumFields: 1,
		Chars:     false,
		NumChars:  0,
		Register:  false,
	}
	output, err := FindUnique(input, op)
	if err != nil {
		t.Error(err)
	}
	answer := []string{
		"I love music of Kartik.",
		"Thanks.",
		"We love music.",
	}
	require.Equal(t, answer, output, "TestUniqueF: The two slice should be the same.")
}

func TestUniqueS(t *testing.T) {
	input := []string{
		"I love music.",
		"A love music.",
		" ",
		"I love music of Kartik.",
		"We love music of Kartik.",
		"Thanks.",
	}

	op := Options{
		Count:     false,
		Duplicate: false,
		Unique:    false,
		Fields:    false,
		NumFields: 0,
		Chars:     true,
		NumChars:  1,
		Register:  false,
	}
	output, err := FindUnique(input, op)
	if err != nil {
		t.Error(err)
	}
	answer := []string{
		" ",
		"I love music of Kartik.",
		"I love music.",
		"Thanks.",
		"We love music of Kartik.",
	}
	require.Equal(t, answer, output, "TestUniqueS: The two slice should be the same.")
}

func TestUniqueCountI(t *testing.T) {
	input := []string{
		"I love music.",
		"I loVe music.",
		"I lOve musIc.",
		" ",
		"I love music of Kartik.",
		"I love music of KaRtik.",
		"Thanks.",
	}

	op := Options{
		Count:     true,
		Duplicate: false,
		Unique:    false,
		Fields:    false,
		NumFields: 0,
		Chars:     false,
		NumChars:  0,
		Register:  true,
	}
	output, err := FindUnique(input, op)
	if err != nil {
		t.Error(err)
	}
	answer := []string{
		"1  ",
		"1 Thanks.",
		"2 I love music of Kartik.",
		"3 I love music.",
	}
	require.Equal(t, answer, output, "TestUniqueCountI: The two slice should be the same.")
}

func TestUniqueFD(t *testing.T) {
	input := []string{
		"We love music.",
		"I lve music.",
		"They love music.",
		"I love music of Kartik.",
		"We loe music of Kartik.",
	}

	op := Options{
		Count:     false,
		Duplicate: true,
		Unique:    false,
		Fields:    true,
		NumFields: 2,
		Chars:     false,
		NumChars:  0,
		Register:  false,
	}
	output, err := FindUnique(input, op)
	if err != nil {
		t.Error(err)
	}
	answer := []string{
		"I love music of Kartik.",
		"We love music.",
	}
	require.Equal(t, answer, output, "TestUniqueFD: The two slice should be the same.")
}

func TestUniqueFError(t *testing.T) {
	input := []string{
		"We love music.",
		"I love music.",
		"They love music.",
		"I love music of Kartik.",
		"We love music of Kartik.",
	}

	op := Options{
		Count:     false,
		Duplicate: false,
		Unique:    false,
		Fields:    true,
		NumFields: 3,
		Chars:     false,
		NumChars:  0,
		Register:  false,
	}
	output, err := FindUnique(input, op)
	if err != nil {
		require.Equal(t, err, "incorrect parameters", "TestUniqueFError: Change Error.")
	}
	answer := []string{
		"I love music of Kartik.",
		"We love music.",
	}
	require.Equal(t, answer, output, "TestUniqueFError: The two slice should be the same.")
}
