package yr

import (
	"bufio"
	"os"
	"testing"
)

func TestCelsiusToFahrenheitString(t *testing.T) {
	type test struct {
		input string
		want  string
	}
	tests := []test{
		{input: "6", want: "42.8"},
		{input: "0", want: "32.0"},
	}

	for _, tc := range tests {
		got, _ := CelsiusToFahrenheitString(tc.input)
		if !(tc.want == got) {
			t.Errorf("expected %s, got: %s", tc.want, got)
		}
	}
}

func TestCelsiusToFahrenheitLine(t *testing.T) {
	type test struct {
		input string
		want  string
	}
	tests := []test{
		{input: "Kjevik;SN39040;18.03.2022 01:50;6", want: "Kjevik;SN39040;18.03.2022 01:50;42.8"},
		{input: "Kjevik;SN39040;07.03.2023 18:20;0", want: "Kjevik;SN39040;07.03.2023 18:20;32.0"},
		{input: "Kjevik;SN39040;08.03.2023 02:20;-11", want: "Kjevik;SN39040;08.03.2023 02:20;12.2"},
	}

	for _, tc := range tests {
		got, _ := CelsiusToFahrenheitLine(tc.input)
		if !(tc.want == got) {
			t.Errorf("expected %s, got: %s", tc.want, got)
		}
	}

}

func TestLineCount(t *testing.T) {

	//åpner filen som skal leses
	file, err := os.Open("kjevik-temp-fahr-20220318-20230318.csv")
	if err != nil {
		t.Fatalf("failed to open file %v", err)
	}
	defer file.Close()

	//teller antall linjer i filen
	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	//sjekker at linjeantallet samsvarer med det vi vil ha
	expectedCount := 16756
	if lineCount != expectedCount {
		t.Errorf("Unexpected number of lines. Got %v, expected %v", lineCount, expectedCount)
	}

}

func TestStudentname(t *testing.T) {
	type test struct {
		input string
		want  string
	}
	tests := []test{

		{input: "Data er gyldig per 18.03.2023 (CC BY 4.0), Meteorologisk institutt (MET);;;",
			want: "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Thea Nielsen"},
	}

	for _, tc := range tests {
		got := Studentname(tc.input)
		if !(tc.want == got) {
			t.Errorf("expected %s, got: %s", tc.want, got)
		}
	}

}
