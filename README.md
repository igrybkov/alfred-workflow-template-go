# Alfred-Workflow Alfred Workflow

Alfred workflow for ...

## How to use this template

- Replace all occurrences of `Alfred-Workflow` with your workflow name
- Run these commands (or `make init`) to set the workflow metadata (feel free to change the values):

```shell
/usr/libexec/PlistBuddy info.plist -c "Set webaddress $(gh repo view --json url -q ".url")"
/usr/libexec/PlistBuddy info.plist -c "Set bundleid com.grybkov.$(gh repo view --json name -q ".name")"
/usr/libexec/PlistBuddy info.plist -c "Set name $(gh repo view --json name -q ".name" | sed -e 's/^alfred-//g')"
/usr/libexec/PlistBuddy info.plist -c "Set description $(gh repo view --json description -q ".description")"
/usr/libexec/PlistBuddy info.plist -c "Set version 1.0.0"
/usr/libexec/PlistBuddy info.plist -c "Set createdby Illia Grybkov"
go run github.com/nishanths/license/v5@latest -o LICENSE mit
```

- Update the `README.md` file
- Delete this section from the `README.md` file

## Features

- ...

## Installation

1. Download the latest release of the workflow.
2. Double-click the downloaded `.alfredworkflow` file to install it.

## Usage

Invoke Alfred and type `{keyword}`.

## Makefile targets

For list of available targets, run `make help`.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License - see the `LICENSE` file for details.
