#!/usr/bin/env bash
echo "Checking for pre-requisites"
echo "==========================="
 # ANSI escape code for red text
RED='\033[0;31m'
# ANSI escape code to reset text color
RESET='\033[0m'
# Check if yq is installed
if ! command -v yq &> /dev/null; then
    echo "yq is not installed. Installing..."
    if [ "$(uname -s)" == "Darwin" ]; then
        # macOS (using Homebrew)
        brew install yq
    else
        # Linux (using snap)
        sudo snap install yq
    fi
fi

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
        echo "      Un-installing Git Bird $version       "
        echo "==============================================="
        echo ""
    fi
else
    echo -e "${RED}Something went wrong: unable to load configuration${RESET}"
    echo ""
    exit 1
fi
# Remove the go package
sudo rm /usr/local/bin/gitbird

echo "✓ Un-Installation Completed."