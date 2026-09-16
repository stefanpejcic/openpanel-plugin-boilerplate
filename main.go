// Command plugin is an OpenPanel plugin boilerplate: fork this repo, rename
// it, and edit servePage below to build your own plugin. OpenPanel execs
// this binary directly - see README.md for the full contract (subcommands,
// readme.txt fields, install flow, multi-arch build/install).
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: plugin <page>")
		os.Exit(1)
	}
	switch os.Args[1] {
	case "page":
		servePage()
	default:
		fmt.Fprintln(os.Stderr, "Unknown command:", os.Args[1])
		os.Exit(1)
	}
}

// servePage prints the HTML OpenPanel embeds into the page at this plugin's
// readme.txt "link" (see readme.txt: link=/new-plugin). OpenPanel wraps this
// in its own page chrome - sidebar, topbar, flashes - so print only the
// content for that one page, not a full <html> document.
func servePage() {
	fmt.Print(`<div class="rounded-lg border border-gray-200 bg-white p-6 shadow-sm dark:border-gray-900 dark:bg-[#090E1A]">
  <h1 class="text-xl font-bold text-gray-900 dark:text-gray-50">Hello from your plugin!</h1>
  <p class="mt-2 text-sm text-gray-500">This page is served by this plugin's own binary. Edit servePage() in main.go to replace it with your own content.</p>
</div>`)
}
