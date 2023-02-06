import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:itorweb/src/features/home/presentation/widgets/torrent_tile.dart';
import 'package:itorweb/src/utils/show_snackbar.dart';

import '../../providers/search_providers.dart';

class ResultPage extends ConsumerWidget {
  const ResultPage({super.key, required this.provider});
  final FutureProvider<Result> provider;

  @override
  Widget build(BuildContext context, ref) {
    final result = ref.watch(provider);
    return result.when(
      data: (data) {
        if (data.isEmpty) {
          return const Center(child: Text('No results'));
        }

        return ListView.builder(
          padding: const EdgeInsets.all(8),
          itemCount: data.length,
          itemBuilder: (context, index) {
            return TorrentTile(torrent: data[index]);
          },
        );
      },
      loading: () => const Center(child: CircularProgressIndicator()),
      skipLoadingOnRefresh: false,
      error: (error, stack) {
        debugPrint(error.toString());
        return Center(
          child: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              InkWell(
                onTap: () => showSnackbar(context, error.toString()),
                child: const Text('Error'),
              ),
              const SizedBox(height: 8),
              ElevatedButton(
                onPressed: () => ref.refresh(provider),
                child: Row(
                  mainAxisSize: MainAxisSize.min,
                  children: const [
                    Text('Search again'),
                    SizedBox(width: 8),
                    Icon(Icons.refresh),
                  ],
                ),
              ),
            ],
          ),
        );
      },
    );
  }
}
