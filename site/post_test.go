package site

import (
	"reflect"
	"testing"
	"time"

	"brage/files"
)

func TestMakePost(t *testing.T) {
	file := files.File{
		files.MarkdownFile,
		"/tmp/test.md",
		[]byte(`---
title: Testing!
published_date: 2020-10-01
---

This is just a test.`),
	}
	expectedTime, _ := time.Parse("2006-01-02", "2020-10-01")
	expected := Post{
		"/blog/test",
		"Testing!",
		expectedTime,
		"<p>This is just a test.</p>\n",
	}

	result := MakePost(file, "/blog")

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Received:\n%+v\nExpected:\n%+v", result, expected)
	}
}

func TestMakePostDefaultMetadata(t *testing.T) {
	file := files.File{
		files.MarkdownFile,
		"/tmp/some-test.md",
		[]byte("This is a test"),
	}
	expected := Post{
		"/blog/some-test",
		"Some Test",
		time.Now(),
		"<p>This is a test</p>\n",
	}

	result := MakePost(file, "/blog")

	result.PublishedDate = expected.PublishedDate // we'll be milliseconds off, so no point in checking

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Received:\n%+v\nExpected:\n%+v", result, expected)
	}
}

func TestMakePostHtmlFile(t *testing.T) {
	file := files.File{
		files.HtmlFile,
		"/tmp/another-test.html",
		[]byte("This is a test"),
	}
	expected := Post{
		"/another-test",
		"Another Test",
		time.Now(),
		"This is a test",
	}

	result := MakePost(file, "/")

	result.PublishedDate = expected.PublishedDate // we'll be milliseconds off, so no point in checking

	if !reflect.DeepEqual(result, expected) {
		t.Fatalf("Received:\n%+v\nExpected:\n%+v", result, expected)
	}
}
