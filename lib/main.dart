import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:itorweb/config.dart';
import 'package:itorweb/data/service/torrent_api_client.dart';
import 'package:itorweb/src/app.dart';

import 'domain/entities/torrent/torrent.dart';
import 'domain/enum/providers.dart';

void main() {
  runApp(const ProviderScope(child: MyApp()));
}
