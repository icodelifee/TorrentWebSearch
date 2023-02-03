import 'dart:convert';

import 'package:diox/diox.dart';
import 'package:flutter/material.dart';

const apiUNAME = '';
const apiPWD = '';
const apiURL = '';
const primaryColor = Color(0xFFFF4151);

Dio get dio {
  return Dio()
    ..options.baseUrl = apiURL
    ..options.headers['Authorization'] = 'Basic ${base64Encode(utf8.encode('$apiUNAME:$apiPWD'))}';
}
