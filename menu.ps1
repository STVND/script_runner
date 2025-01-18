Write-Host("Choose an option:`n")

while($true) {
Write-Host("[0] - Leave")
[string]$userInt = Read-Host

if ($userInt -eq "0") {
    break
} else {
    continue
}

}