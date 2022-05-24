# gfs_simple_implementation

# TO:DO 

Andy: 
- Explore using Bazel
- Add more complex synchronization mechanism to check for completion of Master/Chunkserver Initialization
- Security
- Add timing mechanisms
- How to properly close client connections? We should keep track of them so that proper garbage collection should happen
- Find some type of library to more efficiently process, read, and aggregate logs
- Fix Subtle Bug where if the number of successful Chunkservers is < NUMCHUNKSERVERS, then the error is properly handled
- Log to a file instead of to output. Only log errors and warnings to output
- Optimize generateChunkHandle() such that it is resilience and deterministic under heavy load.
- Add function to cleanly and safely destruct everything. (memory, files, etc)
- Upgrade Lock Manager to support directories
- Create introduction READ ME for GFS repo. 
- Update ChunkSize type from int64 -> uint64
- Insert some nice fun visual to show that tests are passing
- Have better naming than InitChunkserver() and InitMasterServer() to indicate the start process.
- Implement heart beat messages as part of failure recovery
- Load balancing for perf and storage
- Move IO calls such as create and mkdir() out of class construction calls such as NewChunkServer() -> separation of concerns.
- Potential Optimizations can be made in the decision making processs for serializing mutations.
- Still not really using lease manager (not sure we really need it if we maintain a designated primary)
- We never actually clear the write cache on the chunkserver 
- There's an error when writing to an existing file
- Organize the grpc messages members (repeated types, bytes, strings, int, uints, bool)
- We need to check on what happens when you do overlapping creates(), we should prevent this.
- We should have tests for every publicly exposed function
- FIX up IO usage in repo. (IO usage is very coupled into our code, it shouldn't be)
- Update transactionIds -> txids

Vincent:
- Optimization: async forwarding of commitReq from primaryCS to secondaryCS
- Optimization: async sending writeData from client to all replicas
- Observational/tedious: Not sure if our data types are sloppy but there seems to be a lot of type-casting between int, int32, uint64, and int64. Fixable?
- Benchmarking
- More consistent/systematic logging -> consider how logging IO affect benchmarking data when longers program paths have more logs (can we turn off with a flag?) [Mentioned by Andy] 
- Look into more sysematic workload/test generation in test.go
- Use dedicated testing package in stead of running in main
- Transaction ID Generation should be done on chunkserver side.
- Update prior code so that we aren't neglecting error prints. We shouldn't be returning -1 as a indication of error. Should be the error itself. 
At a first glance, client calls such as Read() and Write() should be fixed for this. localWriteToFile() in chunkserver_servicer.go had this issue too but I fixed 
in my latest PR.
- Add flag to default remove temporary generated files.
- The tests in the test suite aren't actually handling any of the returns from the client calls.
- We can probably remove the tests that aren't actually doing checking outputs. (e.g. the ones that are only WriteTest instead of ReadWriteTest)