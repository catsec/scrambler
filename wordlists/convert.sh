#!/bin/bash

for input_file in *.txt; do
    if [ -f "$input_file" ]; then
        output_file="${input_file%.txt}.go"

        var_name="$(basename "$input_file" .txt)"
        var_name="$(echo "$var_name" | sed -E 's/[^a-zA-Z0-9_]/_/g')"

        {
            echo "package main"
            echo
            echo "var $var_name = []string{"

            while IFS= read -r line; do
                echo "    \"$line\","
            done < "$input_file"

            echo "}"
        } > "$output_file"

        echo "Go file '$output_file' created successfully."
    else
        echo "No .txt files found in the current directory."
    fi
done