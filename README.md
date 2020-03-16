# cvmfs_overlayfs
This repository is meant for GSoC 2020 project task for Podman integration to cvmfs. It implements the functionality required as per [this document](https://drive.google.com/file/d/1MR9ZgjHG8ILf-KXseGkC2YUNRuzoSI8Z/view?usp=sharing).

# Usage
- Mount the cvmfs unpacked.cern.ch directory using `sudo mount -t cvmfs unpacked.cern.ch /cvmfs/unpacked.cern.ch`
- Clone the repository and run `sudo go run *.go` from inside the directory


