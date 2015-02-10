# github-get

Simply get a list of files from a (possibly private) github repo. You
will need your github token for private repos.

## Usage

Concatenate a list of files to stdout:

```
REPO=myorg/myrepo TOKEN=123 github-get path/to/file
```

You may write to local files using the syntax `src:dest`:

```
REPO=myorg/myrepo TOKEN=123 github-get path/to/file1:/tmp/one path/to/file1:/tmp/two
```

It is fine to mix the `stdout` and `src:dest` syntax if you wish.
