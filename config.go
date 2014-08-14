package main

const (
	HASH_TRUNC      = 8         // 64-bit
	SPACE_WF        = (1 << 32) // ~ 80G of RAM
	TIME_WF         = (1 << 34) // SPACE_WF * TIME_WF should be >= (1<<64)
	STORE_PROCS     = 16
	GENERATE_PROCS  = 6
	VERIFY_PROCS    = 6
	NUM_BUCKETS     = (SPACE_WF / STORE_PROCS * 5 / 4)
	SEED_BYTES      = 4
	REPORT_INTERVAL = (SPACE_WF / STORE_PROCS / 32) // Keep this a power of 2 if possible
)
