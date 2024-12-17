# Input and output file paths
$inputFile = "slip39 (trezor).txt"    # Path to your input text file
$outputFile = "output.go"   # Path to save the Go array

# Read the lines from the text file
$lines = Get-Content -Path $inputFile

# Convert to Go array format
$goArray = "var english_slip39 := []string{`"" + ($lines -join "`",`"") + "`"}"

# Save the result to a file
$goArray | Out-File -FilePath $outputFile -Encoding utf8

# Output to console
Write-Host "Go array generated:"
Write-Host $goArray