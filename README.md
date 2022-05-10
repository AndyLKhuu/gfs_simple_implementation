# gfs_simple_implementation

# TO:DO 
- Explore using Bazel
- Move from fmt.println -> glog
- Add more complex synchronization mechanism to check for completion of Master/Chunkserver Initialization
- Security
- Better variable/function naming
- Add timing mechanisms
- Implement some type of repolling mechanism from master until chunkservers are up. Add a timeout to avoid long waits
- Fix master to chunk server connection, it only works for 1 chunkserver right now
