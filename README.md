# cursor-update

Checks the installed Cursor AppImage against the latest Linux x64 stable release. If it is behind, downloads the new image, backs up the current one as `cursor.AppImage.bak`, and replaces it.

Requires `7z` (p7zip) to read the version from the installed AppImage.

The AppImage path is set in `main.go`.
