#!/bin/bash

# Directory containing the .sol files
baseFolder="./contracts"
folders=("/evm" "/zevm")

# Output directory for the flattened files
output_dir="./contracts"

# Directory for the class diagrams
diagram_folder="./contract-class-diagrams"

# Loop through all .sol files in the contracts_dir
for folder in "${folders[@]}"; do
    contracts_dir="$baseFolder$folder"
    echo "Processing folder: $contracts_dir"

    for file in "$contracts_dir"/*.sol; do
        # Check if any .sol files exist in the directory
        if [[ ! -e "$file" ]]; then
            echo "No .sol files found in $contracts_dir"
            continue
        fi

        # Get the base name of the file (without path)
        base_name=$(basename "$file")
        echo "Processing file: $file"

        # Construct the output file path
        output_file="$output_dir/Flattened_$base_name"

        # Run the flatten command
        npx hardhat flatten "$file" > "$output_file"

        # Check if the command was successful
        if [ $? -eq 0 ]; then
            echo "Successfully flattened $file to $output_file"
        else
            echo "Error flattening $file"
        fi

        # Run the sol2uml command
        npx sol2uml class "$output_file" -f png -o "$diagram_folder/$folder/$base_name.png"

        # Remove the flattened file
        rm "$output_file"
    done
done
