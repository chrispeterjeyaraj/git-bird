#!/usr/bin/env bash
# Install the required Go packages
echo "==============================================="
echo "      Installing Git Bird v0.1.1       "
echo "==============================================="
echo ""
go get -u github.com/go-git/go-git
if [ $? -ne 0 ]; then
  echo "Failed to install Go packages."
  exit 1
fi
echo ""
echo "✓ Go packages installed."

# Find the script file in the repository
script_path=$(find $(pwd) -name "gitbird.go" | head -1)

if [ -z "$script_path" ]
then
  echo "Script file not found. Please make sure the script is named 'gitbird.go' and is located in the repository."
  exit 1
fi

# Build the Go executable
echo "Building Git Bird Executables"
go build -o gitbird "$script_path"
if [ $? -ne 0 ]; then
  echo "Failed to build Go executable."
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