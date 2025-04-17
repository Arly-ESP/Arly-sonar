import 'package:Arly/core/style.dart';
import 'package:flutter/material.dart';

class MoodSelectionRow extends StatelessWidget {
  final int expandedIndex;
  final ValueChanged<int> onMoodTap;

  const MoodSelectionRow({
    super.key,
    required this.expandedIndex,
    required this.onMoodTap,
  });

  final List<IconData> moodIcons = const [
    Icons.sentiment_very_dissatisfied,
    Icons.sentiment_dissatisfied,
    Icons.sentiment_neutral,
    Icons.sentiment_satisfied,
    Icons.sentiment_very_satisfied,
  ];

  final List<Color> moodColors = const [
    Color(0xFFEF9A9A),
    Color(0xFFFFCC80),
    Color(0xFFFFF176),
    Color(0xFFA5D6A7),
    Color(0xFF66BB6A),
  ];

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        // Header: icon + title
        Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: const [
            Icon(Icons.emoji_emotions, color: AppColors.tealPrimary, size: 22),
            SizedBox(width: 8),
            Text(
              "Comment tu te sens aujourd’hui ?",
              style: TextStyle(
                fontSize: 16,
                fontWeight: FontWeight.w600,
                color: AppColors.primaryGreen,
              ),
            ),
          ],
        ),
        const SizedBox(height: 20),
        // Mood Icons Row
        Row(
          mainAxisAlignment: MainAxisAlignment.center,
          children: List.generate(5, (index) {
            final isSelected = index == expandedIndex;

            return Expanded(
              child: GestureDetector(
                onTap: () => onMoodTap(index),
                child: AnimatedContainer(
                  duration: const Duration(milliseconds: 300),
                  curve: Curves.easeInOut,
                  height: 100,
                  decoration: BoxDecoration(
                    border: Border.all(
                      color: isSelected ? Colors.blueAccent : Colors.transparent,
                      width: 2,
                    ),
                    borderRadius: BorderRadius.only(
                      bottomLeft: index == 0 ? const Radius.circular(50) : Radius.zero,
                      bottomRight: index == 4 ? const Radius.circular(50) : Radius.zero,
                    ),
                    gradient: isSelected
                        ? LinearGradient(
                      colors: [moodColors[index], moodColors[index].withOpacity(0.6)],
                      begin: Alignment.topLeft,
                      end: Alignment.bottomRight,
                    )
                        : null,
                  ),
                  child: Center(
                    child: AnimatedContainer(
                      duration: const Duration(milliseconds: 300),
                      width: isSelected ? 70 : 50,
                      height: isSelected ? 70 : 50,
                      decoration: BoxDecoration(
                        shape: BoxShape.circle,
                        color: moodColors[index],
                        boxShadow: isSelected
                            ? [
                          const BoxShadow(
                            color: Colors.black26,
                            blurRadius: 8,
                            offset: Offset(0, 4),
                          ),
                        ]
                            : [],
                      ),
                      child: Icon(
                        moodIcons[index],
                        color: Colors.white,
                        size: isSelected ? 40 : 30,
                      ),
                    ),
                  ),
                ),
              ),
            );
          }),
        ),
      ],
    );
  }
}
