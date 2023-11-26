# Download the 1990 BRFSS SAS transport file data
curl https://www.cdc.gov/brfss/annual_data/1990/files/CDBRFS90XPT.zip -o .\CDBRFS90.zip

Expand-Archive -Path .\CDBRFS90.zip -DestinationPath .\CDBRFS90
Move-Item .\CDBRFS90\CDBRFS90.XPT .
Remove-Item .\CDBRFS90 -Recurse
Remove-Item .\CDBRFS90.zip
