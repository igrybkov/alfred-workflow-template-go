package main

// Package is called aw
import (
	aw "github.com/deanishe/awgo"
)

// Workflow is the main API
var wf *aw.Workflow

func init() {
	// Create a new Workflow using default settings.
	// Critical settings are provided by Alfred via environment variables,
	// so this *will* die in flames if not run in an Alfred-like environment.
	wf = aw.New()

}

// Your workflow starts here
func run() {
	query := ""

	// Use wf.Args so magic actions are handled
	args := wf.Args()
	if len(args) > 0 {
		query = args[0]
	}

	// Disable UIDs so Alfred respects our sort order. Without this,
	// it may bump read/unpublished books to the top of results, but
	// we want to force them to always be below unread books.
	wf.Configure(aw.SuppressUIDs(true))

	if query == "" {
		wf.WarnEmpty("No matching items", "Try a different query?")
	} else {
		wf.NewItem("Your query is " + query).Valid(true)
	}
	// Add a "Script Filter" result
	// Send results to Alfred
	wf.SendFeedback()
}

func main() {
	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	wf.Run(run)
}
