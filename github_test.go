package main

import (
	"reflect"
	"testing"

	"github.com/google/go-github/github"
)

func TestNewGitHubClient(t *testing.T) {
	tests := []struct {
		name string
		want *github.Client
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGitHubClient(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGitHubClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEvents(t *testing.T) {
	type args struct {
		client *github.Client
		org    *string
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			GetEvents(tt.args.client, tt.args.org)
		})
	}
}

func TestSieveOutEvents(t *testing.T) {
	type args struct {
		events []*github.Event
		org    *string
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SieveOutEvents(tt.args.events, tt.args.org)
		})
	}
}
