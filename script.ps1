
Write-Host "Step 1: Creating a process"
Start-Process "notepad.exe"


Write-Host "Step 2: Listing processes"
Get-Process
Start-Sleep -Seconds 3

Write-Host "Step 3: Killing the process created in Step 1"
Stop-Process -Name "notepad.exe"