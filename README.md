# Simple Implementation of Google File System

Simple Implementation of Google File System is a simplified version of GFS written in Golang. Our implementation emulates the normal operations of the GFS system with the assumption that nodes would communicate over gRpc. 


For a more indepth description of our implementation and its performance, refer to


[Insert link to our paper]

# Architecture

Our implementation of GFS consists of a Master Server, a configurable number of Chunkservers, as well as varying numbers of clients. Each of these components communicate with each other through gRpc. In our implementation, we provide an client interface which supports connecting to GFS as well as basic interactions with the filesystem 

Main Features
- Support for Create()/Read()/Write()/Remove() of plaintext files
- Support for Concurrent Mutations


# Testing
For those interested in running our system and seeing text files actually being managed and created, please feel free to fork our repo and test it by running the following command from the gfs directory.

<code> go run main.go </code>

We also welcome you to look through our client interface and try to stress test our system.

# Authors

Andy Khuu (andykhuu@stanford.edu) 

Vincent Heng (vheng@stanford.edu)
