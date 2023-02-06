import 'package:flutter/material.dart';

import '../../../../../domain/enum/providers.dart';
import '../../providers/search_providers.dart';
import '../widgets/home_appbar.dart';
import '../widgets/result_page.dart';

class HomeScreen extends StatelessWidget {
  const HomeScreen({super.key});

  @override
  Widget build(BuildContext context) {
    return DefaultTabController(
      length: SearchProvider.values.length,
      child: Scaffold(
        appBar: const HomeAppBar(),
        body: TabBarView(
          children: [
            ResultPage(provider: leetxProvider),
            ResultPage(provider: tpbProvider),
            ResultPage(provider: ytsProvider),
            ResultPage(provider: katProvider),
            ResultPage(provider: limeProvider),
            ResultPage(provider: tgxProvider),
          ],
        ),
      ),
    );
  }
}
