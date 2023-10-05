# git-bird
Git Bird is a command-line tool that simplifies common Git operations. It provides a convenient way to perform tasks like committing changes, pulling updates from a remote repository, and more.

[![Go build](https://github.com/chrispeterjeyaraj/git-bird/actions/workflows/go.yml/badge.svg)](https://github.com/chrispeterjeyaraj/git-bird/actions/workflows/go.yml)

## Prerequisites :clipboard:

Before installing Git Bird, You will also need to have GoLang installed. To check if you have Go installed, open your terminal window and run the following command:
```sh
go --version
```

If you see a version number in the output, it means that Go is installed. If not, you can download and install go using the below documentation:
```sh
https://go.dev/doc/install
```
## Installation :inbox_tray:
To install Git Bird, please follow the steps below:
1. Clone the Git Bird repository to your local machine by running the following command in your terminal:
```sh
git clone https://github.com/chrispeterjeyaraj/git-bird
```
2. Navigate into the cloned repository by running the following command in your terminal:
```sh
cd git-bird
```
3. Run the install.sh script by running the following command in your terminal:
```sh
./install.sh
```
This will install the required Go packages, copy the gptcommit executable to your /usr/local/bin directory, and add an alias to your .bashrc file.

## Usage :zap:
To use Git Bird, Navigate to the directory of the Git repository you wish to commit changes to, and run the following command in your terminal:
```sh
gitbird

```
This will show the commands Git Bird supports and provide help documentation.

## Uninstallation :outbox_tray:
To uninstall Git Bird, please follow the steps below:
1. Navigate to the directory where you cloned the Git Bird repository by running the following command in your terminal:
```sh
cd <path-to-gitgird-repo>
```
2. Run the uninstall.sh script by running the following command in your terminal:
```sh
./uninstall.sh
```
This will remove the gitbird executable from your /usr/local/bin directory

## Contribution :raised_hands:
I welcome contributions from the community! If you find a bug 🐛 or have an idea for a new feature 💡, please open an issue or submit a pull request.

## LICENSE :scroll:
Git Bird is licensed under the MIT License. See the [LICENSE](./LICENSE) file for more information.

----------
- If you found Git Bird useful, please consider giving this repo a star ⭐️!

