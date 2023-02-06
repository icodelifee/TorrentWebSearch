import 'package:diox/diox.dart';
import 'package:flutter/material.dart';

const apiURL = '';
const primaryColor = Color(0xFFFF4151);

Dio get dio {
  return Dio()
    ..options.baseUrl = apiURL
    ..options.headers['Access-Control-Allow-Origin'] = '*'
    ..options.headers['Access-Control-Allow-Methods'] = 'GET, OPTIONS'
    ..options.headers['Access-Control-Allow-Headers'] = 'Origin, Content-Type, X-Auth-Token'
    ..options.headers['Access-Control-Allow-Credentials'] = 'true'
    ..options.headers['Content-Type'] = 'application/json';
}
