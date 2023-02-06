import 'package:flutter/material.dart';

void showSnackbar(BuildContext context, String content) {
  if (context.mounted) {
    ScaffoldMessenger.of(context).showSnackBar(SnackBar(content: Text(content)));
  }
}
