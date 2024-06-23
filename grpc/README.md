# Setup Repos
- x-proto: private SSH key in secret
- x-go-proto: public SSH key in Deploy Keys

# Publish x-go-proto packages
Run these commands inside the `x-go-proto` project's root folder:
```bash
go mod tidy

git tag v0.0.6
git push origin v0.0.6
GOPROXY=proxy.golang.org go list -m github.com/isd-sgcu/rpkm67-go-proto@v0.0.6
```
In this example, the package is `github.com/isd-sgcu/rpkm67-go-proto`

Note that the version number should be the latest version.
- v0.0.5 -> v0.0.6: patches
- v0.0.5 -> v0.1.0: minor changes
- v0.0.5 -> v1.0.0: major changes with breaking changes

or just do whatever you want with the version number.