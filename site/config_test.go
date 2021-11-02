package site

import (
	"os"
	"path"
	"testing"
)

var exampleConfig = `
title: Young Fatigue
description: New Single ‘Dislocation‘ Out Now!
image: icon.png
rootUrl: https://youngfatigue.com/

data:
  links:
    - title: Instagram
      url: https://www.instagram.com/youngfatigue/
    - title: Facebook
      url: https://www.facebook.com/YoungFatigue
    - title: YouTube
      url: https://www.youtube.com/channel/UCfAGdI0HuS7C6J1m7_xwARQ
  quotes:
    - text: It’s actually really good... really good.
      author: Jeremy Vine, BBC Radio 2
    - text: "[Dislocation] is super fun!"
      author: Elise Cobain, Amazon Music
`

func createExampleSite() (string, error) {
	temporaryDirectory, err := os.MkdirTemp("", "examplesite")
	if err != nil {
		return "", err
	}

	filepath := path.Join(temporaryDirectory, "config.yaml")
	file, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	data := []byte(exampleConfig)
	_, err = file.Write(data)
	if err != nil {
		return "", err
	}

	return temporaryDirectory, nil
}

func TestLoad(t *testing.T) {
	dirPath, err := createExampleSite()
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	defer os.RemoveAll(dirPath)

	config, err := Load(dirPath)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if config.Title != "Young Fatigue" {
		t.Fatalf("Incorrect config loaded: %v", config)
	}
}

func TestLoadMissingSite(t *testing.T) {
	dirPath, err := createExampleSite()
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	os.RemoveAll(dirPath) // remove the directory so it doesn't exist

	_, err = Load(dirPath)
	if err == nil {
		t.Fatalf("Expected error but got nil")
	}
}

func TestLoadMissingConfig(t *testing.T) {
	dirPath, err := createExampleSite()
	if err != nil {
		t.Fatalf("Unable to create example site: %v", err)
	}
	os.Remove(path.Join(dirPath, "config.yaml")) // remove the config file so it doesn't exist

	_, err = Load(dirPath)
	if err == nil {
		t.Fatalf("Expected error but got nil")
	}
}
