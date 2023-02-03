import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:itorweb/config.dart';

import '../data/service/torrent_api_client.dart';

final torrentClientProvider = Provider<TorrentAPIClient>(
  (ref) => TorrentAPIClient(dio),
);
