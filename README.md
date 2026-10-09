# cursor-update

Checks the installed Cursor AppImage latest official release. If its version is behind, downloads the new image, backs up the current one as `cursor.AppImage.bak`, and replaces it.

Requires `7z` cli command (p7zip) to read the version from the installed AppImage.

## How to use

Install the command:

```bash
go install github.com/raswonders/cursor-update@latest
```

The binary is placed in `$(go env GOPATH)/bin`, or in `GOBIN` when that is set. Pass the installed AppImage with `-path`, or set `CURSOR_APPIMAGE`:

```bash
cursor-update -path /path/to/cursor.AppImage
```

`-platform` defaults to `linux-x64` and `-track` defaults to `stable`.
