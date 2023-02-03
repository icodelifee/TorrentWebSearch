import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:itorweb/src/features/home/providers/query_provider.dart';

import '../../../../domain/entities/torrent/torrent.dart';
import '../../../../providers/torrent_client_provider.dart';

typedef Result = List<Torrent>;

final leetxProvider = FutureProvider.autoDispose<Result>((ref) async {
  final client = ref.watch(torrentClientProvider);
  final query = ref.watch(queryProvider);
  if (query.isEmpty) return [];

  final results = await client.get1337x(query);
  return results;
});

final tpbProvider = FutureProvider.autoDispose<Result>((ref) async {
  final client = ref.watch(torrentClientProvider);
  final query = ref.watch(queryProvider);
  if (query.isEmpty) return [];

  final results = await client.getTPB(query);
  return results;
});

final ytsProvider = FutureProvider.autoDispose<Result>((ref) async {
  final client = ref.watch(torrentClientProvider);
  final query = ref.watch(queryProvider);
  if (query.isEmpty) return [];

  final results = await client.getYTS(query);
  return results;
});

final katProvider = FutureProvider.autoDispose<Result>((ref) async {
  final client = ref.watch(torrentClientProvider);
  final query = ref.watch(queryProvider);
  if (query.isEmpty) return [];

  final results = await client.getKat(query);
  return results;
});

final limeProvider = FutureProvider.autoDispose<Result>((ref) async {
  final client = ref.watch(torrentClientProvider);
  final query = ref.watch(queryProvider);
  if (query.isEmpty) return [];

  final results = await client.getLime(query);
  return results;
});
