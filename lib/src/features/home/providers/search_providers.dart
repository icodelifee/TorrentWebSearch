import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:itorweb/src/features/home/providers/query_provider.dart';

import '../../../../domain/entities/torrent/torrent.dart';
import '../../../../providers/torrent_client_provider.dart';

typedef Result = List<Torrent>;

final leetxProvider = FutureProvider<Result>((ref) async {
  final client = ref.watch(torrentClientProvider);
  final query = ref.watch(queryProvider);
  if (query.isEmpty) return [];

  final results = await client.get1337x(query);
  return results;
});

final tpbProvider = FutureProvider<Result>((ref) async {
  final client = ref.watch(torrentClientProvider);
  final query = ref.watch(queryProvider);
  if (query.isEmpty) return [];

  final results = await client.getTPB(query);
  return results;
});

final ytsProvider = FutureProvider<Result>((ref) async {
  final client = ref.watch(torrentClientProvider);
  final query = ref.watch(queryProvider);
  if (query.isEmpty) return [];

  final results = await client.getYTS(query);
  return results;
});

final katProvider = FutureProvider<Result>((ref) async {
  final client = ref.watch(torrentClientProvider);
  final query = ref.watch(queryProvider);
  if (query.isEmpty) return [];

  final results = await client.getKat(query);
  return results;
});

final limeProvider = FutureProvider<Result>((ref) async {
  final client = ref.watch(torrentClientProvider);
  final query = ref.watch(queryProvider);
  if (query.isEmpty) return [];

  final results = await client.getLime(query);
  return results;
});

final tgxProvider = FutureProvider<Result>((ref) async {
  final client = ref.watch(torrentClientProvider);
  final query = ref.watch(queryProvider);
  if (query.isEmpty) return [];

  final results = await client.getTGX(query);
  return results;
});

final rarbgProvider = FutureProvider<Result>((ref) async {
  final client = ref.watch(torrentClientProvider);
  final query = ref.watch(queryProvider);
  if (query.isEmpty) return [];

  final results = await client.getRarbg(query);
  return results;
});
