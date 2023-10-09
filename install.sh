#!/usr/bin/env bash
 echo "Checking for pre-requisites:"
 # ANSI escape code for red text
RED='\033[0;31m'
# ANSI escape code to reset text color
RESET='\033[0m'
# Check if yq is installed
if ! command -v yq &> /dev/null; then
     echo ""
     echo "yq is not installed. Installing..."
    if [ "$(uname -s)" == "Darwin" ]; then
        # macOS (using Homebrew)
        brew install yq
    else
        # Linux (using snap)
        sudo snap install yq
    fi
  else
    echo ""
    echo "✓ yq is installed."
fi

# Install the required Go packages
go get -u github.com/go-git/go-git
if [ $? -ne 0 ]; then
  echo -e "${RED}Failed to install Go packages.${RESET}"
  exit 1
fi
echo ""
echo "✓ Go packages installed."
echo ""

# Now that yq is installed or already present, proceed with YAML extraction
if [ -f "./config/config.yaml" ]; then
    # Read the YAML file and extract the "port" value
    version=$(yq e '.version' ./config/config.yaml)
    # Check if the command produced an error
    if [ $? -ne 0 ]; then
        echo -e "${RED}Error running yq command.${RESET}"
        exit 1
    elif [ "$version" = "null" ]; then
        echo -e "${RED}Valid version could not be found for installation${RESET}"
        exit 1
    else
        echo "==============================================="
        echo "      Installing Git Bird $version       "
        echo "==============================================="
        echo ""
    fi
else
    echo -e "${RED}Something went wrong: unable to load configuration${RESET}"
    echo ""
    exit 1
fi
# Find the script file in the repository
script_path=$(find $(pwd) -name "gitbird.go" | head -1)

if [ -z "$script_path" ]
then
  echo -e "${RED}Script file not found. Please make sure the script is named 'gitbird.go' and is located in the repository.${RESET}"
  exit 1
fi

# Build the Go executable
echo "Building Git Bird Executables"
go build -o gitbird "$script_path"
if [ $? -ne 0 ]; then
  echo -e "${RED}Failed to build Go executable.${RESET}"
  exit 1
fi
echo ""
echo "✓ Build Complete"

# Copy the executable to /usr/local/bin
echo "Copying executable to /usr/local/bin..."
sudo cp gitbird /usr/local/bin/

echo ""
echo "✓ Installation Completed."
echo ""
echo "==============================================="
echo "    You can now use the 'gitbird' command."
echo "==============================================="