package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	prompt "github.com/c-bata/go-prompt"
	gosxnotifier "github.com/deckarep/gosx-notifier"
	banner "github.com/dimiro1/banner"
	asciigraph "github.com/guptarohit/asciigraph"
	table "github.com/jedib0t/go-pretty/table"
	opts "github.com/jpillora/opts"
)

//go-prompt
func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "users", Description: "Store the username and age"},
		{Text: "articles", Description: "Store the article text posted by user"},
		{Text: "comments", Description: "Store the text commented to articles"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	//banner start
	nyanFilePath := "assets/nyancat.txt"
	nyanBuf, err := ioutil.ReadFile(nyanFilePath)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}

	banner.Init(os.Stdout, true, true, bytes.NewBuffer(nyanBuf))
	//banner stop

	//go-prompt start
	fmt.Println("go-prompt demo,Please select item.")
	t := prompt.Input("> ", completer)
	fmt.Println("You selected " + t)
	//go-prompt end

	//At a minimum specifiy a message to display to end-user.
	note := gosxnotifier.NewNotification("tt") //Optionally, set a title
	note.Title = "tt override!!! 💰"

	//Optionally, set a subtitle
	note.Subtitle = "tt subtitle"
	//Optionally, set a sound from a predefined set.
	note.Sound = gosxnotifier.Basso

	//Optionally, specifiy a url or bundleid to open should the notification be
	//clicked.
	note.Link = "https://www.google.com" //or BundleID like: com.apple.Terminal

	//Optionally, an app icon (10.9+ ONLY)
	note.AppIcon = "gopher.png"

	//Optionally, a content image (10.9+ ONLY)
	note.ContentImage = "gopher.png"

	//If necessary, check error
	if err := note.Push(); err != nil {
		log.Println("Uh oh!")
	}

	typicalMap := make(map[string]string)
	typicalMap["-a"] = "All, append"
	typicalMap["-b"] = "Buffer,block size, batch"
	typicalMap["-c"] = "Command, check"
	typicalMap["-d"] = "Debug, delete, directory"
	typicalMap["-D"] = "Define"
	typicalMap["-e"] = "Execute, edit"
	typicalMap["-f"] = "File, force"
	typicalMap["-h"] = "Headers, help"
	typicalMap["-i"] = "Initialize"
	typicalMap["-I"] = "Include"
	typicalMap["-k"] = "Keep, kill"
	typicalMap["-l"] = "List, long, load"
	typicalMap["-m"] = "Message"
	typicalMap["-n"] = "Number, not"
	typicalMap["-o"] = "Output"
	typicalMap["-p"] = "Port, protocol"
	typicalMap["-q"] = "Quiet"
	typicalMap["-r"] = "Recurse, reverse"
	typicalMap["-s"] = "Silent, subject"
	typicalMap["-t"] = "Tag"
	typicalMap["-u"] = "User"
	typicalMap["-v"] = "Verbose"
	typicalMap["-V"] = "Version"
	typicalMap["-w"] = "Width, warning"
	typicalMap["-x"] = "Enable debugging, extract"
	typicalMap["-y"] = "Yes"
	typicalMap["-z"] = "Enable compression"

	//table start
	rt := table.NewWriter()
	rt.SetOutputMirror(os.Stdout)
	rt.AppendHeader(table.Row{"Option", "Typical meaning"})
	for op, tm := range typicalMap {
		rt.AppendRow([]interface{}{op, tm})
	}
	rt.AppendRow([]interface{}{})
	rt.Render()
	//table stop

	//opts start
	//https://github.com/jpillora/opts 相比之下这个更好一些
	//https://github.com/jessevdk/go-flags
	c := config{}
	opts.Parse(&c)
	log.Printf("%+v", c)
	//opts stop

	//cobra start most famous
	//https://github.com/spf13/cobra

	//cobra stop

	//promptui start
	//https://github.com/manifoldco/promptui

	//promotui stop

	//progressbar start
	//https://github.com/cheggaaa/pb
	//https://github.com/vbauerster/mpb
	//https://github.com/Gosuri/uilive
	//progressbar stop

	//asciigraph start
	data := []float64{3, 4, 9, 6, 2, 4, 5, 8, 5, 10, 2, 7, 2, 5, 6}
	graph := asciigraph.Plot(data)

	fmt.Println(graph)
	//asciigraph stop

	//color start
	//https://github.com/fatih/color
	//color stop

	//web start
	//https://github.com/gin-Gonic/gin
	//https://github.com/codemodus/parth
	//validate
	//https://github.com/Go-playground/validator
	//web stop

	//tool
	//https://github.com/thoas/Go-funk

	//rest client
	//https://github.com/h2non/gentleman
	//https://github.com/Go-resty/resty
	//deepcopy
	//https://github.com/ulule/deepcopier
	//deepcopy
}

/**
$ go build -o my-prog
$ ./my-prog --help

  Usage: my-prog [options]

  Options:
  --file, -f   file to load
  --lines, -l  number of lines to show
  --help, -h   display help

$ ./my-prog -f foo.txt -l 42
{File:foo.txt Lines:42}
*/
type config struct {
	File  string `opts:"help=file to load"`
	Lines int    `opts:"help=number of lines to show"`
}
