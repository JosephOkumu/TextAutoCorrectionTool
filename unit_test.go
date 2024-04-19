package autocorrectiontool
import (
	"testing"
)
func TestHexDecimal(t *testing.T) {
	checks := []struct {
		input    string
		accepted string
	}{
		{"1E (hex) files were added", "30 files were added"},
		{"2FF (hex) children got admitted into the university", "767 children got admitted into the university"},
	}
	for _, check := range checks {
		output := HexToDecimal(check.input)
		if output != check.accepted {
			t.Errorf("HexDecimal(%q) = %q; want %q", check.input, output, check.accepted)
		}
	}
}
func TestBinToDecimal(t *testing.T) {
	checks := []struct {
		input    string
		accepted string
	}{
		{"It has been 10 (bin) years", "It has been 2 years"},
		{"The chicken laid 100 (bin) eggs", "The chicken laid 4 eggs"},
	}
	for _, check := range checks {
		output := BinToDecimal(check.input)
		if output != check.accepted {
			t.Errorf("BinDecimal(%q) = %q; want %q", check.input, output, check.accepted)
		}
	}
}
func TestReplaceWithAn(t *testing.T) {
	checks := []struct {
		input    string
		accepted string
	}{
		{"There it was. A amazing rock!", "There it was. An amazing rock!"},
		{"There it was a amazing rock!", "There it was an amazing rock!"},
		{"That is a egg", "That is an egg"},
	}
	for _, check := range checks {
		output := ReplaceWithAn(check.input)
		if output != check.accepted {
			t.Errorf("ReplaceWithAn(%q) = %q; want %q", check.input, output, check.accepted)
		}
	}
}
func TestUpperCase(t *testing.T) {
	checks := []struct {
		input    string
		accepted string
	}{
		{"Ready, set, go (up)", "Ready, set, GO"},
		{"I am coming (up) to the wedding (up)", "I am COMING to the WEDDING"},
	}
	for _, check := range checks {
		output := UpperCase(check.input)
		if output != check.accepted {
			t.Errorf("UpperCase(%q) = %q; want %q", check.input, output, check.accepted)
		}
	}
}
func TestLowCase(t *testing.T) {
	checks := []struct {
		input    string
		accepted string
	}{
		{"I should stop SHOUTING (low)", "I should stop shouting"},
		{"The MANGO (low) is in the PLATE (low)", "The mango is in the plate"},
	}
	for _, check := range checks {
		output := LowerCase(check.input)
		if output != check.accepted {
			t.Errorf("LowCase(%q) = %q; want %q", check.input, output, check.accepted)
		}
	}
}
func TestCapitalize(t *testing.T) {
	checks := []struct {
		input    string
		accepted string
	}{
		{"Welcome to the Brooklyn bridge (cap)", "Welcome to the Brooklyn Bridge"},
		{"I am going (cap) to the library (cap)", "I am Going to the Library"},
	}
	for _, check := range checks {
		output := Capitalize(check.input)
		if output != check.accepted {
			t.Errorf("Capitalize (%q) = %q; want %q", check.input, output, check.accepted)
		}
	}
}
func TestModifyText(t *testing.T) {
	checks := []struct {
		input    string
		accepted string
	}{
		{"This is so exciting (up, 2)", "This is SO EXCITING"},
		{"The weather is (up, 3) very very horrible (cap, 2)", "THE WEATHER IS very Very Horrible"},
	}
	for _, check := range checks {
		output := ModifyText(check.input)
		if output != check.accepted {
			t.Errorf("ModifyText (%q) = %q; want %q", check.input, output, check.accepted)
		}
	}
}
func TestSetPunctuation(t *testing.T) {
	checks := []struct {
		input    string
		accepted string
	}{
		{"I am exactly how they describe me: ' awesome '", "I am exactly how they describe me: 'awesome'"},
		{"As Elton John said: ' I am the most well-known homosexual in the world '", "As Elton John said: 'I am the most well-known homosexual in the world'"},
	}
	for _, check := range checks {
		output := SetPunctuation(check.input)
		if output != check.accepted {
			t.Errorf("SetPunctuation (%q) = %q; want %q", check.input, output, check.accepted)
		}
	}
}
