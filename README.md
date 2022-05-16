# gfs_simple_implementation

# TO:DO 
- Explore using Bazel
- Add more complex synchronization mechanism to check for completion of Master/Chunkserver Initialization
- Security
- Better variable/function naming
- Add timing mechanisms
- How to properly close client connections? We should keep track of them so that proper garbage collection should happen
- Find some type of library to more efficiently process, read, and aggregate logs
- Figure out how many chunkservers we want to spin up and the timeouts we need to do so.
- Clean up communication set up calls, maybe it's too many nested goroutines
- Add a flag which supresses log calls (different levels of logs perhaps?)
- Fix Subtle Bug where if the number of successful Chunkservers is < NUMCHUNKSERVERS, then the error is properly handled
- Optimize generateChunkHandle() such that it is resilience and deterministic under heavy load.
- Add function to cleanly and safely destruct everything. (memory, files, etc)