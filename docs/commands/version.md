# Version

BMX allows displaying the version of the command by running `version`

```
> bmx version
bmx/2.0.1 git/bd60b68 
```

If you installed by source, when you run you should see:

```
> bmx version
bmx/nostamp
```

The version and commit of a build is stamped by the bazel build process. Installation using `go get` does not perform this stamping, which results in the response `nostamp`. 