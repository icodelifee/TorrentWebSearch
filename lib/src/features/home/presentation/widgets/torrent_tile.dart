import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:itorweb/src/utils/show_snackbar.dart';

import '../../../../../domain/entities/torrent/torrent.dart';

class TorrentTile extends StatelessWidget {
  const TorrentTile({super.key, required this.torrent});
  final Torrent torrent;

  @override
  Widget build(BuildContext context) {
    return MouseRegion(
      cursor: SystemMouseCursors.click,
      child: Card(
        child: ListTile(
          shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(8)),
          onTap: () async {
            await Clipboard.setData(ClipboardData(text: torrent.magnet!));
            if (context.mounted) {
              showSnackbar(context, 'Magnet copied to clipboard');
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
              ],
              if (torrent.added?.isNotEmpty ?? false) ...[
                const SizedBox(width: 8),
                const Icon(Icons.access_time),
                const SizedBox(width: 4),
                Flexible(
                  child: Text(
                    torrent.added!,
                    maxLines: 1,
                  ),
                ),
              ]
            ],
          ),
          trailing: const Icon(Icons.link),
        ),
      ),
    );
  }
}
