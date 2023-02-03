import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

import '../../../../../domain/enum/providers.dart';
import '../../providers/query_provider.dart';

class HomeAppBar extends ConsumerStatefulWidget with PreferredSizeWidget {
  const HomeAppBar({
    super.key,
  });

  @override
  ConsumerState<HomeAppBar> createState() => _HomeAppBarState();

  @override
  Size get preferredSize => const Size.fromHeight(170);
}

class _HomeAppBarState extends ConsumerState<HomeAppBar> {
  final searchController = TextEditingController();

  @override
  void dispose() {
    searchController.dispose();
    super.dispose();
  }

  Future<void> onSearch(String? value) async {
    final query = value ?? searchController.text;
    if (query.isEmpty) return;
    ref.read(queryProvider.notifier).state = query;
  }

  @override
  Widget build(BuildContext context) {
    return AppBar(
      toolbarHeight: widget.preferredSize.height,
      title: Column(
        children: [
          const Text('iTorrent Search', style: TextStyle(fontSize: 20, letterSpacing: 1.2)),
          const SizedBox(height: 12),
          TextField(
            onSubmitted: onSearch,
            controller: searchController,
            decoration: InputDecoration(
              border: const OutlineInputBorder(
                borderRadius: BorderRadius.all(Radius.circular(10)),
                borderSide: BorderSide.none,
              ),
              hintText: 'Search',
              filled: true,
              fillColor: Colors.white,
              prefixIcon: IconButton(
                icon: const Icon(Icons.search),
                onPressed: () => onSearch(null),
              ),
            ),
          ),
        ],
      ),
      bottom: TabBar(
        labelColor: Colors.white,
        indicatorColor: Colors.white,
        unselectedLabelColor: Colors.white,
        tabs: SearchProvider.values.map((e) => Tab(text: e.name)).toList(),
      ),
    );
  }
}
