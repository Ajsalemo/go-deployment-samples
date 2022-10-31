# go-deployment-samples
Various Go web framework examples

To run these examples locally and then deploy, do the following:

1. After cloning this repository, `cd` to any of these framework directories and run `go mod tidy`.
2. Run `go run main.go`.
3. To deploy, for example with Local Git:
  - Run `git init`
  - Add the remote repository with `git remote add azure <remote URI>`
  - Add and commit - `git add .` followed by `git commit "initial commit"`
  - Push to the remote - `git push azure master`
