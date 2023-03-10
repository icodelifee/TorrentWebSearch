// ignore: avoid_web_libraries_in_flutter
import 'dart:html' as html;

import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';
import 'package:itorweb/config.dart';

import '../../../../../domain/enum/providers.dart';
import '../../../../../providers/theme_provider.dart';
import '../../../../utils/bubble_tab_indicator.dart';
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
    final theme = ref.watch(themeProvider);
    final isLight = theme == Brightness.light;

    return AppBar(
      toolbarHeight: widget.preferredSize.height,
      title: Column(
        children: [
          Stack(
            alignment: Alignment.center,
            children: [
              const Text(
                appName,
                style: TextStyle(fontSize: 20, letterSpacing: 1.2),
              ),
              Align(
                alignment: Alignment.centerRight,
                child: Row(
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    const _GithubIconButton(),
                    const SizedBox(width: 12),
                    IconButton(
                      icon: isLight ? const Icon(Icons.dark_mode) : const Icon(Icons.light_mode, color: Colors.white),
                      onPressed: () {
                        ref.read(themeProvider.notifier).state = isLight ? Brightness.dark : Brightness.light;
                      },
                    ),
                  ],
                ),
              )
            ],
          ),
          const SizedBox(height: 12),
          TextField(
            onSubmitted: onSearch,
            controller: searchController,
            style: const TextStyle(color: Colors.black),
            cursorColor: Theme.of(context).primaryColor,
            decoration: InputDecoration(
              border: const OutlineInputBorder(
                borderRadius: BorderRadius.all(Radius.circular(10)),
                borderSide: BorderSide.none,
              ),
              hintText: 'Search',
              hintStyle: const TextStyle(color: Colors.grey),
              filled: true,
              fillColor: Colors.white,
              prefixIcon: IconButton(
                icon: const Icon(
                  Icons.search,
                  color: Colors.black,
                ),
                onPressed: () => onSearch(null),
              ),
            ),
          ),
        ],
      ),
      bottom: _TabBar(isLight: isLight),
    );
  }
}

class _TabBar extends StatelessWidget with PreferredSizeWidget {
  const _TabBar({
    required this.isLight,
  });

  final bool isLight;

  @override
  Widget build(BuildContext context) {
    return LayoutBuilder(
      builder: (context, constraints) => TabBar(
        isScrollable: constraints.maxWidth < 700,
        splashFactory: NoSplash.splashFactory,
        labelColor: isLight ? primaryColor : Colors.white,
        indicatorColor: Colors.white,
        padding: const EdgeInsets.symmetric(horizontal: 12),
        indicator: BubbleTabIndicator(
          indicatorHeight: 25.0,
          indicatorColor: isLight ? Colors.white : primaryColor,
          indicatorRadius: 12,
          insets: const EdgeInsets.symmetric(horizontal: 5.0),
          padding: const EdgeInsets.symmetric(horizontal: 10.0),
        ),
        unselectedLabelColor: Colors.white,
        tabs: SearchProvider.values.map((e) => Tab(text: e.name)).toList(),
      ),
    );
  }

  @override
  Size get preferredSize => Size.infinite;
}

class _GithubIconButton extends StatelessWidget {
  const _GithubIconButton();

  @override
  Widget build(BuildContext context) {
    return IconButton(
      onPressed: () {
        html.window.open('https://github.com/icodelifee/TorrentWebSearch', 'repo');
      },
      icon: Image.asset(
        'assets/github.png',
        width: 30,
        color: Colors.white,
      ),
    );
  }
}
