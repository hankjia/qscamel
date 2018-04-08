package constants

import "runtime"

// DefaultConfigContent is the default config config.
const DefaultConfigContent = `concurrency: 0
log_file: ~/.qscamel/qscamel.log
log_level: error
pid_file: ~/.qscamel/qscamel.pid
database_file: ~/.qscamel/qscamel.db
`

var (
	// DefaultConcurrency is default num of objects being migrated concurrently.
	DefaultConcurrency = runtime.NumCPU() * 100
	// DefaultWorkerRatio is the default num of worker's ratio which is the golden ratio in math.
	// ref: https://en.wikipedia.org/wiki/Golden_ratio
	DefaultWorkerRatio = 0.618
)

// Path store all path related constants.
const (
	Path         = "~/.qscamel"
	ConfigPath   = Path + "/qscamel.yaml"
	DatabasePath = Path + "/qscamel.db"
	LogPath      = Path + "/qscamel.log"
	PIDPath      = Path + "/qscamel.pid"
)
