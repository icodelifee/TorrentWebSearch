import 'package:flutter/material.dart';
import 'package:flutter_riverpod/flutter_riverpod.dart';

final themeProvider = StateProvider<Brightness>(
  (ref) => MediaQueryData.fromWindow(WidgetsBinding.instance.window).platformBrightness,
);
