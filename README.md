# retarget-github-prs

A quick example using golang and github.com/google/go-github/github to retarget pull requests on github to a different branch.

## Setup

You can install golang on Mac using Homebrew:

Make sure XCode is installed:

`xcode-select --install`

Install Brew:

`/usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"`

Install Go:

`brew install go`

Add golang env vars to your `~/.bash_profile` then source it.

```bash
echo 'export GOPATH="$HOME/code/go"' >> ~/.bash_profile
echo 'export GOBIN="$GOPATH/bin"' >> ~/.bash_profile
echo 'export PATH="$GOBIN:$PATH"' >> ~/.bash_profile
source ~/.bash_profile
```

Clone this repo into your $GOPATH/src dir:

```bash
mkdir -p $GOPATH/src/github.com/clintmod
cd $GOPATH/src/github.com/clintmod
git clone git@github.com:clintmod/retarget-github-prs.git
cd retarget-github-prs
```

Build the project:

```bash
  go get
  go build
```

You should now have a binary in the dir `./retarget-github-prs`

## Usage

Make sure you GITHUB_USERNAME and GITHUB_PASSWORD have env vars set:

```bash
echo 'export GITHUB_USERNAME="YOUR_GITHUB_USERNAME"' >> ~/.bash_profile
echo 'export GITHUB_PASSWORD="YOUR_GITHUB_PASSWORD"' >> ~/.bash_profile
source ~/.bash_profile
```

Then you can call the executable like:

`./retarget-github-prs <github account> <old branch> <new branch> <comma separated list of repos>`

Example:

`./retarget-github-prs clintmod master development repo1,repo2,repo3`

This will retarget all prs in repo1, repo2 and repo3 that are targeted at master to the development branch.

## Notes

Github Account is seperate argument to specify a repository that you contribute to that's not your own.
