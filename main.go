package main

func main() {
	conf := &config{supportedCommands: getCommands()}
	startRepl(conf)
}
