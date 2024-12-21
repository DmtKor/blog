package commander

// struct commander - stores DB struct, chi router, WaitGroup for handlers, channels for CLI input and 
// for critical errors, should_handle bool - if false, all requests -> err_handler

// commander.InitServer - initialise server, starts routine commander.Wait

// commander.Wait - listen channels (select), listen channels, call commander.ShutdownServer when needed

// commander.ShutdownServer - set should_handle to false, wait for everything in WaitGroup to be done, 
// close DB, close app