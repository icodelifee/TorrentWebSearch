# Web Torrent Search

A simple, fast and reliable torrent search engine built with Flutter and Golang to find and download torrents.
This repository contains both the client and the server.

## Getting Started



##### Prerequisites
- Golang version 1.18.
- Flutter version ^3.0

## Setup
1. Clone the repository: `https://github.com/icodelifee/TorrentWebSearch.git`

#### Server Setup
1. Navigate to the Go server: `cd server/`
2. Install the Go dependencies: `go get`
3. Run the Go server: `go run main.go`

#### Client Setup
1. Install the Flutter dependencies: `flutter pub get`
2. Edit ```lib/config.dart``` and add the API url
    ``` dart
    ...
    const apiURL = '<YOUR-SERVER-URL>';
    ...
    ```
    
3. Run the Flutter app: `flutter run`


## Usage
1. Open the Flutter app on your browser
2. Enter a search query in the search bar
3. Browse through the results and click on a torrent to download

## Contributing
We welcome contributions to this repository. If you have any suggestions or bug reports, feel free to open an issue or create a pull request.