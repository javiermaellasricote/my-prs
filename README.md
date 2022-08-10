# My PRs
Track all the PRs from a GitHub project that are relevant to you with a single command.

## Installation:
This CLI works as a wrapper of the official GitHub CLI `gh`, so you first need to install that one for `my-prs` to work.
```
brew install gh
```

Since this is a Go project, you will need to have the go installed in order to retrieve and compile the correct version of `my-prs`.
```
brew install go
```

Do not forget to add Go's `bin` path to your paths. For that, go to your `~/.bashrc` or `~/.zshrc` and add the following line:
```
export PATH=$PATH:$GOPATH/bin
```

Finally, to install the CLI, just do:
```
go install https://github.com/javiermaellasricote/my-prs@v0.3.0
```

## How it works:
Run `my-prs <YOUR_PROJECT_OR_USER_NAME>` to get a description of the PRs that you opened and the ones that are waiting for your review.

The information shown for each PR is the following:
* Repo where it was created.
* Branch name
* CICD Workflow status
* URL to the PR in GitHub

Running the command as is will only retrieve information from the last 10 active repos. If you want to increase the number of repos where the search is perform, add the number to the end of the CLI call: `my-prs <YOUR_PROJECT_OR_USER_NAME> <NUMBER_OF_REPOS>`. Take into account that the more repos there is, the slower the CLI call will get.
