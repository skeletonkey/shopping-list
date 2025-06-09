# List Item Logging

It is expected that Items will be added and removed from lists on a frequent basis that storing this in a relational database is not a useful.

We will be utilizing a poor-man's Kafka like eventing system.

## File

The file will reside next to the database file (as long as SQLite is being used.)

It will be a binary file containing fixed with entries.

Filename: list_history.dat

## Caching

The system prioritizes performance over data retention. Since I/O operations are expensive and stress the file system events are stored in a cache and bulk written to the data file. The interval for these writes is configurable and should be provided in seconds.

It is understood that this is a lists app and losing data is "ok." That being said all best practices will be taken to ensure no data lose.

To ensure system safety a locking mechanism needs to be used. If none exists this will be done by creating a "touch" file next to the binary's location with the same name with ".touch" appended to it.

There needs to be a configuration to start alerting if the file remains locked for certain amount of time which is set as seconds.

There is also an override time configuration which will break the file lock. This can be set to 0 to never brake the lock or some time interval expressed in seconds.

It is the user's responsibility to understand these settings, their implications, and set them to their tolerances for potential data lose:

- time_between_writes (seconds)
- data_file_lock_alert_threshold (seconds) - 0 meaning no alerts will be provided
- data_file_lock_override_threshold (seconds) - 0 meaning lock will never be broken

The service will not start if the file is locked.

## Entry

An entry will consist of a list_id, item_id, and a bit indicating true or false. The IDs are integers and should represent the integer of the architecture that it is being run on.

This means that a file created on a 32 bit system will not be readable on a 64 bit system. The data file should be considered not portable!

## Utilities

Utility program(s) will reside in a utility directory and properly labeled/documented.

The following utilities are provided - they may be separate programs or one program - the implementation is not set here.

### Backup

The backup utility allows the binary file to be translated into a ASCII file and back again. There is no merging these will be overwrite functions. It is the user's responsibility not to lose data.

The binary file will be changed into a '|' (pipe) delimited file with the true/false bit being represented by a 1 and 0. Entries will be delimited by a '|' so that the file can be transported between systems without having to worry about line endings.

This is an inefficient process prioritizing portability. It is recommended to run compression on the binary file before backing it up.

### Compression

The data file is not meant for long term historical data. It is meant to represent the current state of the list. Compressing that data file from time to time will save memory and storage space as well as increase processing time on start up.

Compression can be trigger manually through the utility, by time, time interval, or space. Setting these to 0 indicates that they are not being used (turned off.)

These setting need to be understood by the user and set to the needs of their use case:

- time (seconds since midnight) run once daily - midnight can not be chosen
- timeInterval (seconds) run every X seconds
- size (megabytes) min size before running compression
- sizeCheckInterval (seconds) how often to check the data file's size; default 86400 (once a day)
