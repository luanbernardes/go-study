#!/bin/bash

# Path to the file you want to clear
file_path="$1"

# Check if the file exists
if [ -f "$file_path" ]; then
    # Clear the file contents
    > "$file_path"
    echo "Cleared contents of $file_path"
else
    echo "File $file_path does not exist"
fi