Create a new bonzai branch

{{cmd .Name}} is a wrapper around the gh-cli commands used to create a new repository based on a template. The newly created repo will be cloned into the current working directory. {{cmd .Name}} will prompt the user to provide a `bonzaiBranchName` and `templateName`.

(Requires that `gh-cli` is installed and configured correctly.)
