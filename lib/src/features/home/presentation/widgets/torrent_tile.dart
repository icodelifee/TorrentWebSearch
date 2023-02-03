import 'package:flutter/material.dart';
import 'package:flutter/services.dart';

import '../../../../../domain/entities/torrent/torrent.dart';

class TorrentTile extends StatelessWidget {
  const TorrentTile({super.key, required this.torrent});
  final Torrent torrent;

  @override
  Widget build(BuildContext context) {
    return Card(
      child: ListTile(
        onTap: () async {
          await Clipboard.setData(ClipboardData(text: torrent.magnet!));
          if (context.mounted) {
            ScaffoldMessenger.of(context).showSnackBar(
              const SnackBar(
                content: Text('Magnet copied to clipboard'),
              ),
            );
          }
        },
        title: Text(torrent.title ?? ''),
        subtitle: Row(
          children: [
            if (torrent.size != null) ...[
              const Icon(Icons.folder),
              const SizedBox(width: 4),
              Text(torrent.size!),
            ],
            if (torrent.seeds != null) ...[
              const SizedBox(width: 8),
              const Icon(Icons.arrow_upward),
              const SizedBox(width: 4),
              Text(torrent.seeds!),
            ],
            if (torrent.leechs != null) ...[
              const SizedBox(width: 8),
              const Icon(Icons.arrow_downward),
              const SizedBox(width: 4),
              Text(torrent.leechs!),
            ]
          ],
        ),
        trailing: const Icon(Icons.link),
      ),
    );
  }
}
