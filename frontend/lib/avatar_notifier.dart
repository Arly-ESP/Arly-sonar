import 'package:flutter/material.dart';

class AvatarNotifier extends ValueNotifier<String> {
  AvatarNotifier(String value) : super(value);
}

final avatarNotifier = AvatarNotifier('assets/arly.png');
