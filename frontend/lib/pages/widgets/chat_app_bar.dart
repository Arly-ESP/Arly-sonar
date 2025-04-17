import 'dart:ui';
import 'package:flutter/material.dart';

class ChatAppBar extends StatelessWidget implements PreferredSizeWidget {
  final String selectedTheme;
  final ValueChanged<String> onThemeChanged;

  const ChatAppBar({
    super.key,
    required this.selectedTheme,
    required this.onThemeChanged,
  });

  static final List<Map<String, dynamic>> _themes = [
    {"label": "Famille", "value": "famille", "icon": Icons.home_rounded, "color": Colors.pinkAccent},
    {"label": "Amis", "value": "amis", "icon": Icons.people_alt_rounded, "color": Colors.lightBlueAccent},
    {"label": "Santé", "value": "santé", "icon": Icons.favorite_rounded, "color": Colors.greenAccent},
    {"label": "Ennemis", "value": "ennemis", "icon": Icons.warning_amber_rounded, "color": Colors.deepOrange},
    {"label": "Méditations", "value": "méditations", "icon": Icons.self_improvement_rounded, "color": Colors.deepPurpleAccent},
  ];

  void _showThemeSelector(BuildContext context) {
    showModalBottomSheet(
      context: context,
      backgroundColor: Colors.transparent,
      isScrollControlled: true,
      builder: (_) => ThemeDropSheet(
        selectedTheme: selectedTheme,
        onThemeSelected: onThemeChanged,
      ),
    );
  }

  @override
  @override
  Widget build(BuildContext context) {
    return AppBar(
      backgroundColor: Colors.transparent,
      elevation: 0,
      automaticallyImplyLeading: false,
      leading: Padding(
        padding: const EdgeInsets.only(left: 8.0),
        child: IconButton(
          icon: const Icon(Icons.arrow_back_ios_new_rounded, color: Colors.white),
          onPressed: () => Navigator.of(context).maybePop(),
          splashRadius: 24,
        ),
      ),
      flexibleSpace: Stack(
        fit: StackFit.expand,
        children: [
          Image.asset('assets/dodo.png', fit: BoxFit.cover),
          Container(
            decoration: BoxDecoration(
              color: Colors.black.withOpacity(0.4),
              borderRadius: const BorderRadius.vertical(
                bottom: Radius.circular(24),
              ),
            ),
          ),
        ],
      ),
      centerTitle: true,
      title: Padding(
        padding: const EdgeInsets.only(top: 12),
        child: InkWell(
          onTap: () => _showThemeSelector(context),
          borderRadius: BorderRadius.circular(32),
          child: Container(
            padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 8),
            decoration: BoxDecoration(
              color: Colors.white.withOpacity(0.12),
              borderRadius: BorderRadius.circular(32),
              border: Border.all(color: Colors.white.withOpacity(0.2)),
            ),
            child: Row(
              mainAxisSize: MainAxisSize.min,
              children: [
                const Icon(Icons.water_drop_rounded, color: Colors.white, size: 18),
                const SizedBox(width: 8),
                Text(
                  selectedTheme,
                  style: const TextStyle(color: Colors.white, fontWeight: FontWeight.w500),
                ),
                const SizedBox(width: 4),
                const Icon(Icons.expand_more_rounded, color: Colors.white, size: 20),
              ],
            ),
          ),
        ),
      ),
    );
  }

  @override
  Size get preferredSize => const Size.fromHeight(72);
}

// ------------------------------
// DROPLET MENU MODAL
// ------------------------------

class ThemeDropSheet extends StatelessWidget {
  final String selectedTheme;
  final ValueChanged<String> onThemeSelected;

  const ThemeDropSheet({
    super.key,
    required this.selectedTheme,
    required this.onThemeSelected,
  });

  @override
  Widget build(BuildContext context) {
    final List<Map<String, dynamic>> themes = ChatAppBar._themes;

    return BackdropFilter(
      filter: ImageFilter.blur(sigmaX: 15, sigmaY: 15),
      child: Container(
        padding: const EdgeInsets.symmetric(vertical: 24),
        decoration: BoxDecoration(
          color: Colors.white.withOpacity(0.9),
          borderRadius: const BorderRadius.vertical(top: Radius.circular(30)),
        ),
        child: Column(
          mainAxisSize: MainAxisSize.min,
          children: themes.map((theme) {
            final bool isSelected = theme["value"] == selectedTheme;

            return AnimatedSlide(
              duration: const Duration(milliseconds: 300),
              offset: const Offset(0, 0),
              child: GestureDetector(
                onTap: () {
                  onThemeSelected(theme["value"]);
                  Navigator.pop(context);
                },
                child: Container(
                  margin: const EdgeInsets.symmetric(vertical: 8, horizontal: 24),
                  padding: const EdgeInsets.symmetric(vertical: 16, horizontal: 20),
                  decoration: BoxDecoration(
                    color: theme["color"].withOpacity(0.9),
                    borderRadius: BorderRadius.circular(32),
                    boxShadow: [
                      BoxShadow(
                        color: theme["color"].withOpacity(0.3),
                        blurRadius: 14,
                        offset: const Offset(0, 6),
                      ),
                    ],
                  ),
                  child: Row(
                    children: [
                      Icon(theme["icon"], color: Colors.white, size: 22),
                      const SizedBox(width: 12),
                      Text(
                        theme["label"],
                        style: TextStyle(
                          color: Colors.white,
                          fontSize: 17,
                          fontWeight: isSelected ? FontWeight.bold : FontWeight.w500,
                        ),
                      ),
                    ],
                  ),
                ),
              ),
            );
          }).toList(),
        ),
      ),
    );
  }
}
