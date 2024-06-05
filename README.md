# file_splitter_n_reconstructor-golang
file splitter  and reconstructor written by golang.

### split command
```shell
go run main.go splitter.go reconstructor.go utils.go split ./temp.zip ./temp_100MB 100
```
### reconstruct command
```shell
go run main.go splitter.go reconstructor.go utils.go reconstruct ./temp_100MB ./original_file.zip
```