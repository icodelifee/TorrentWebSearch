import 'package:json_annotation/json_annotation.dart';

part 'torrent.g.dart';

@JsonSerializable()
class Torrent {
  final String? title;
  final String? seeds;
  final String? leechs;
  final String? size;
  final String? added;
  final String? link;
  final String? magnet;

  const Torrent({
    this.title,
    this.seeds,
    this.leechs,
    this.size,
    this.added,
    this.link,
    this.magnet,
  });

  @override
  String toString() {
    return 'Torrent(title: $title, seeds: $seeds, leechs: $leechs, size: $size, added: $added, link: $link, magnet: $magnet)';
  }

  factory Torrent.fromJson(Map<String, dynamic> json) {
    return _$TorrentFromJson(json);
  }

  Map<String, dynamic> toJson() => _$TorrentToJson(this);
}
