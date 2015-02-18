# github-get

Simply get a list of files from a (possibly private) github repo. You
will need your github token for private repos.

## Usage

Concatenate a list of files to `stdout`:

```
REPO=myorg/myrepo \
TOKEN=123 \
github-get path/to/file
```

You may write to local files using the syntax `src:dest`, where `dest`
is a directory:

```
REPO=myorg/myrepo \
TOKEN=123 \
github-get \
  path/to/file1:/tmp/ \
  path/to/file2:/tmp/
```

It is fine to mix the `stdout` and `src:dest` syntax if you wish.

If `src` ends with `/` it is interpreted as a directory, and all first
level files in the directory are downloaded. This is not recursive.

```
REPO=myorg/myrepo \
TOKEN=123 \
github-get path/to/dir/:/tmp/
```

## Docker

Comes with `Dockerfile`. I build image as follows:

```
CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' github-get.go
docker build -r rlister/github-get .
```

You can pull ready-made image from docker hub:

```
docker pull rlister/github-get
```

and run it:

```
docker run \
  -e REPO=myorg/myrepo \
  -e TOKEN=123 \
  rlister/github-get path/to/file1
```
