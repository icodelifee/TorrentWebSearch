# Web Torrent Search

A simple, fast and reliable torrent search engine built with Flutter and Golang to find and download torrents.
This repository contains both the client and the server.

## Getting Started



##### Prerequisites
- Golang version 1.18.
- Flutter version ^3.0

##### Setup

Clone the repository:
```
git clone https://github.com/icodelifee/TorrentWebSearch.git
```
###### Server

``` bash
cd /server
go get
go run main.go
```

###### App

1. Edit ```lib/config.dart``` and add the API url
    ``` dart
    ...
    const apiURL = '<YOUR-SERVER-URL>';
    ...
    ```
2. Install the Flutter dependencies:
    `flutter pub get`
2. Run the app
    ``` sh
    flutter run -d chrome
    ```
3. Build the app
    ``` sh 
    flutter build web
    ```

### Contributing
We welcome contributions to this repository. If you have any suggestions or bug reports, feel free to open an issue or create a pull request.