# Downloads-Organizer

Organizes your download folder based on their type and extensions.

***If there is a file with unspecified extension, Downloads-Organizer will automatically create a new sub folder for the extension under Others. All future files with same extension will be saved under the same folder.***

- Documents
- Text Files
- Pictures
- Music
- Videos
- Compressed
- Programs
- Others
  - Other extension folders
  - files without extensions
-  Folders
   - All folders
  
 ## Usage
 > ### Set the environment variable `SORT_FOLDER_DESTINATION` to download destination folder and SORT_FOLDER_SLEEP_MINUTES to set sleep minutes to check download folder again in your operating system and run the executable file from releases or compile with `go build main.go`

## Instructions to run the executables from releases
- ### Windows
  - Make sure you have set environment variables as shown above in USAGE
  - Open file expolorer and navigate to `%appdata%\Microsoft\Windows\Start Menu\Programs\Startup`
  - paste the downloaded file in this folder
  - now Downloads-Organizer will start automatically on every boot
- ### Linux
  - Type `sudo nano /lib/systemd/system/downloads-organizer.service`
  - Paste following content in the file. Replace `<path-to-your-downloaded-file>`
    ```
    [Unit]
    Description=Sorts the downloads folder

    [Service]
    ExecStart=<path-to-your-downloaded-file>

    [Install]
    WantedBy=multi-user.target
    ```
   - Type `sudo systemctl edit downloads-organizer.service`
   - Paste following content in the file. Replace `<path-to-your-download-folder>`
     ```
     [Service]
      Environment="SORT_FOLDER_DESTINATION=<path-to-your-download-folder>"
      Environment="SORT_FOLDER_SLEEP_MINUTES=720"
      ```
   - Run `sudo systemctl enable downloads-organizer.service`
   - Then `sudo systemctl start downloads-organizer.service`
   - To verify if the downloads organizer is runnung
     - Run `sudo systemctl status downloads-organizer.service`
     - It will show Active as seen below
       ![image](https://user-images.githubusercontent.com/63334479/173121054-550a396d-b287-4a28-a9cb-544c98d46389.png)

---
### If your finding difficulty in setting up feel free to create an [issue](https://github.com/rabilrbl/Downloads-Organizer/issues)
---
