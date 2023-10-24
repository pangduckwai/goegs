package config

type Config struct {
	// variant     string   // Game variant ID
	// players     int      // Number of players
	DbHost []string // Main database host
	DbPort []string // Main database port
	DbName []string // Main database name
	DbUser []string // Main database role name
	Dbpass []string // Main database password
	// DbPool      int      // Main database connection poll size
	// Socket      string   // Port of the simulation process to bound to
	// Concurrent  int      // Number of concurrent threads to run simulations
	// Runs        int      // Number of simulation batches to run
	// Simulations int      // Number of simulations in a batch
	// Levels      int      // Number of levels after the current node to display
	// Sleep       int      // Total no. of milliseconds to sleep for each run
	// verbose     bool     // Enable more verbose run log if true
	// command     string   // Specify the task to run
	// message     string   // Message to sent to the running instance
	// file        string   // Path of the configuration file
	// path        []uint8  // Path of tree to navigate
	// nid         []int64  // Node id in database
}
