# unpack

Recursively extract files

## Install

1. Download the required release:
   - https://github.com/stephendotcarter/unpack/releases
1. Make the `unpack` binary executable and move it to your `PATH`:
   ```
   chmod +x ./unpack-darwin
   sudo mv ./unpack-darwin /usr/local/bin/unpack
   ```
1. You should now be able to use `unpack`:
   ```
   unpack
   v1.0
   Usage: unpack <file>...
   ```

## Usage

- Here is an example compressed file
  ```
  $ tar -tf mydata.tgz
  subdir1.tar.gz
  subdir2.zip
  ```
- Run `unpack`:
  ```
  $ unpack mydata.tgz
  Unpacking "mydata.tgz"
  + mydata.tgz
  + subdir1.tar.gz
  + subdir2.zip
  ```
- View the extracted files:
  ```
  $ find .
  .
  ./mydata.tgz
  ./mydata
  ./mydata/subdir1.tar.gz
  ./mydata/subdir1
  ./mydata/subdir1/file2.txt
  ./mydata/subdir1/file1.txt
  ./mydata/subdir2.zip
  ./mydata/subdir2
  ./mydata/subdir2/data3.dat
  ./mydata/subdir2/data4.csv
  ```
