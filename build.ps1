$platforms = @("windows/amd64", "linux/amd64", "darwin/amd64", "darwin/arm64")
$outputNames = @("scrambler.exe", "scrambler-linux-amd64", "scrambler-macos-intel", "scrambler-macos-arm")

for ($i = 0; $i -lt $platforms.Count; $i++) {
    $platform = $platforms[$i]
    $output = $outputNames[$i]
    $split = $platform -split "/"
    $env:GOOS = $split[0]
    $env:GOARCH = $split[1]
    
    if (Test-Path -Path $output) {
        Remove-Item -Path $output -Force
        Write-Host "Removed existing file: $output"
    }
    
    & go build -o $output
    Write-Host "Built $output"
    
    $signatureFile = "$output.asc"
    if (Test-Path -Path $signatureFile) {
        Remove-Item -Path $signatureFile -Force
        Write-Host "Removed existing signature: $signatureFile"
    }
    
    Write-Host "Signing $output..."
    & gpg --detach-sign --armor -o $signatureFile $output
    if ($LASTEXITCODE -eq 0) {
        Write-Host "Successfully signed $output -> $signatureFile"
    } else {
        Write-Host "Failed to sign $output" -ForegroundColor Red
        break
    }
}