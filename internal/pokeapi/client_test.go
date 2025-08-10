package pokeapi

import (
	"testing"
)

func TestFetchLocationValidURL(t *testing.T) {
	// Test with the base URL
	result, err := FetchLocation(BaseURL)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
		return
	}

	if result == nil {
		t.Errorf("expected result to not be nil")
		return
	}

	// Check that we got some results
	if len(result.Results) == 0 {
		t.Errorf("expected some location results")
		return
	}

	// Check that the first result has a name
	if result.Results[0].Name == "" {
		t.Errorf("expected first result to have a name")
		return
	}
}

func TestFetchLocationInvalidURL(t *testing.T) {
	// Test with an invalid URL
	_, err := FetchLocation("https://invalid-url-that-does-not-exist.com")
	if err == nil {
		t.Errorf("expected an error for invalid URL")
		return
	}
}

func TestBaseURLConstant(t *testing.T) {
	expected := "https://pokeapi.co/api/v2/location-area/"
	if BaseURL != expected {
		t.Errorf("expected BaseURL to be %s, got %s", expected, BaseURL)
	}
}
