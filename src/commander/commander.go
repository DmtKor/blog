package commander

import (
	"blog/src/managedb"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/go-chi/chi/v5"
)


type Commander struct {
	Database managedb.DB      // DB struct, all ops with data here
	Router *chi.Mux           // HTTP router
	HandlersWG sync.WaitGroup // Every handler except err_handler increases this and do .Done() in the end
	CritErrorCh chan string   // Critical errors from handlers (ones that require to stop server from working)
	UserInputCh chan string   // CLI input from user
	ShouldHandle bool         // Signals middleware if requests should be redirected to err_handler
	ShouldHandleMx sync.Mutex // Locked when accessing ShouldHandle (it will not cost much but will prevent some rare errors)
}

// Initialise server, starts routine commander.Wait
// dbInitStr - "user=your_username password=your_password dbname=name_of_database sslmode=disable"
// addr - address to listen (for HTTP server)
func (commander *Commander) InitServer(dbInitStr string, addr string) {
	// Initialize DB
	commander.Database = managedb.DB{}
	err := commander.Database.Init(dbInitStr)
	if err != nil {
		fmt.Println("Can not access DB:", err)
		os.Exit(1) // In this moment its safe to just close app
	}
	// Initialize goroutine tools (channels do not need buffers since they'll be used only once)
	// There's no need to somehow 'initialise' WaitGroup
	commander.CritErrorCh = make(chan string)
	commander.UserInputCh = make(chan string)
	// Initialize router and start server
	commander.Router = chi.NewRouter()
	// Start listening for events
	go commander.wait()
	// Start listening CLI
	go commander.listenCLI()

	/*
	TODO: add handlers and use middleware on them!! (err_handler should not be added)
	*/
	commander.ShouldHandle = true
	http.ListenAndServe(addr, commander.Router)
}

// Listen channels (select), listen channels, call commander.ShutdownServer when needed
func (commander *Commander) wait() {
	var s string
	for {
		select {
		// Handler sent 'critical error' signal to commander
		case s = <-commander.CritErrorCh:
			fmt.Println("[CRITICAL ERROR] " + s)
			commander.shutdownServer(1) // End of work
		// User sent '!quit' command
		case s = <-commander.UserInputCh:
			if s == "!quit" {
				commander.shutdownServer(0) // End of work
			} else {
				fmt.Println("Unknown command")
				continue
			}
		}
	}
}

// Set ShouldHandle to false, wait for everything in WaitGroup to be done, close DB, close app
// Also closes its channels (its not really needed)
// code - which value program should return
func (commander *Commander) shutdownServer(code int) {
	fmt.Println("[INFO] Shutting down...")
	commander.ShouldHandleMx.Lock()
	commander.ShouldHandle = false
	commander.ShouldHandleMx.Unlock()
	commander.HandlersWG.Wait()
	commander.Database.Close()
	close(commander.CritErrorCh)
	close(commander.UserInputCh)
	os.Exit(code)
}

// Listens os.Stdin for input (with newline delimiter) and send text to commander.wait routine
func (commander *Commander) listenCLI() {
	for {
		text := ""
		fmt.Fscanln(os.Stdout, &text) // I DON'T KNOW WHY ITS STDOUT. STDIN DO NOT WORK NO MATTER WHAT I DO
									  // MAYBE ITS BECAUSE "GO TEST" DOES SMTH
									  // I DON'T CARE
		commander.UserInputCh<- text
	}
}