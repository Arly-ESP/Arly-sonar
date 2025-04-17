import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';

class AvatarNotifier extends ChangeNotifier {
  int _selectedAvatarIndex = 0;

  int get selectedAvatarIndex => _selectedAvatarIndex;

  Future<void> loadAvatarSelection() async {
    final prefs = await SharedPreferences.getInstance();
    _selectedAvatarIndex = prefs.getInt('selectedAvatarIndex') ?? 0;
    notifyListeners();
  }

  Future<void> saveAvatarSelection(int index) async {
    final prefs = await SharedPreferences.getInstance();
    await prefs.setInt('selectedAvatarIndex', index);
    _selectedAvatarIndex = index;
    notifyListeners();
  }

  void onAvatarTap(int index) {
    saveAvatarSelection(index);
  }
}
