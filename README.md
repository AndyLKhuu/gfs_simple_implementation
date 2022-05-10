# gfs_simple_implementation

# TO:DO 
- Explore using Bazel
- Move from fmt.println -> log.printf()
- Add more complex synchronization mechanism to check for completion of Master/Chunkserver Initialization
- Security
- Better variable/function naming
- Add timing mechanisms
- How to properly close client connections? We should keep track of them so that proper garbage collection should happen
- Find some type of library to more efficiently process, read, and aggregate logs
- Figure out how many chunkservers we want to spin up and the timeouts we need to do so.