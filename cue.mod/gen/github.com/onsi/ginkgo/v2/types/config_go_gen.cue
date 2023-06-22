// Code generated by cue get go. DO NOT EDIT.

//cue:generate cue get go github.com/onsi/ginkgo/v2/types --exclude=deprecated

package types

import "pkg.go.dev/time"

// Configuration controlling how an individual test suite is run
#SuiteConfig: {
	RandomSeed:        int64
	RandomizeAllSpecs: bool
	FocusStrings: [...string] @go(,[]string)
	SkipStrings: [...string] @go(,[]string)
	FocusFiles: [...string] @go(,[]string)
	SkipFiles: [...string] @go(,[]string)
	LabelFilter:           string
	FailOnPending:         bool
	FailFast:              bool
	FlakeAttempts:         int
	DryRun:                bool
	PollProgressAfter:     time.#Duration
	PollProgressInterval:  time.#Duration
	Timeout:               time.#Duration
	EmitSpecProgress:      bool
	OutputInterceptorMode: string
	SourceRoots: [...string] @go(,[]string)
	GracePeriod:     time.#Duration
	ParallelProcess: int
	ParallelTotal:   int
	ParallelHost:    string
}

#VerbosityLevel: uint // #enumVerbosityLevel

#enumVerbosityLevel:
	#VerbosityLevelSuccinct |
	#VerbosityLevelNormal |
	#VerbosityLevelVerbose |
	#VerbosityLevelVeryVerbose

#values_VerbosityLevel: {
	VerbosityLevelSuccinct:    #VerbosityLevelSuccinct
	VerbosityLevelNormal:      #VerbosityLevelNormal
	VerbosityLevelVerbose:     #VerbosityLevelVerbose
	VerbosityLevelVeryVerbose: #VerbosityLevelVeryVerbose
}

#VerbosityLevelSuccinct:    #VerbosityLevel & 0
#VerbosityLevelNormal:      #VerbosityLevel & 1
#VerbosityLevelVerbose:     #VerbosityLevel & 2
#VerbosityLevelVeryVerbose: #VerbosityLevel & 3

// Configuration for Ginkgo's reporter
#ReporterConfig: {
	NoColor:        bool
	Succinct:       bool
	Verbose:        bool
	VeryVerbose:    bool
	FullTrace:      bool
	ShowNodeEvents: bool
	JSONReport:     string
	JUnitReport:    string
	TeamcityReport: string
}

// Configuration for the Ginkgo CLI
#CLIConfig: {
	//for build, run, and watch
	Recurse:      bool
	SkipPackage:  string
	RequireSuite: bool
	NumCompilers: int

	//for run and watch only
	Procs:                     int
	Parallel:                  bool
	AfterRunHook:              string
	OutputDir:                 string
	KeepSeparateCoverprofiles: bool
	KeepSeparateReports:       bool

	//for run only
	KeepGoing:       bool
	UntilItFails:    bool
	Repeat:          int
	RandomizeSuites: bool

	//for watch only
	Depth:       int
	WatchRegExp: string
}

// Configuration for the Ginkgo CLI capturing available go flags
// A subset of Go flags are exposed by Ginkgo.  Some are available at compile time (e.g. ginkgo build) and others only at run time (e.g. ginkgo run - which has both build and run time flags).
// More details can be found at:
// https://docs.google.com/spreadsheets/d/1zkp-DS4hU4sAJl5eHh1UmgwxCPQhf3s5a8fbiOI8tJU/
#GoFlagsConfig: {
	//build-time flags for code-and-performance analysis
	Race:      bool
	Cover:     bool
	CoverMode: string
	CoverPkg:  string
	Vet:       string

	//run-time flags for code-and-performance analysis
	BlockProfile:         string
	BlockProfileRate:     int
	CoverProfile:         string
	CPUProfile:           string
	MemProfile:           string
	MemProfileRate:       int
	MutexProfile:         string
	MutexProfileFraction: int
	Trace:                string

	//build-time flags for building
	A:             bool
	ASMFlags:      string
	BuildMode:     string
	Compiler:      string
	GCCGoFlags:    string
	GCFlags:       string
	InstallSuffix: string
	LDFlags:       string
	LinkShared:    bool
	Mod:           string
	N:             bool
	ModFile:       string
	ModCacheRW:    bool
	MSan:          bool
	PkgDir:        string
	Tags:          string
	TrimPath:      bool
	ToolExec:      string
	Work:          bool
	X:             bool
}