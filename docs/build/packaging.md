# MacOS: Packaging bmx for distribution

For distribution in MacOS, the common deployment model for this is as a package (`.pkg`). To create a new package of BMX for distribution you can perform the following steps:

1. Create a new release in GitHub to build the version
2. Unzip the binary into a separate directory

```
COMPANY="acme"
VERSION="x.y.z"
pkgbuild --root "./bmx" --identifier "${COMPANY}.aws.bmx" --version "${VERSION}" --install-location "/usr/local/bin" "bmx.pkg"
```

3. Place the package file in your intended distribution tool
