import 'package:flutter/material.dart';
import 'package:Arly/core/style.dart';

class MoodSummary extends StatelessWidget {
  final List<dynamic> userActivities;

  MoodSummary({
    super.key,
    required this.userActivities,
  });

  final List<String> moods = const ["sad", "angry", "neutral", "anxious", "happy"];

  final Map<String, IconData> moodIcons = const {
    "sad": Icons.sentiment_dissatisfied,
    "angry": Icons.sentiment_very_dissatisfied,
    "neutral": Icons.sentiment_neutral,
    "anxious": Icons.sentiment_satisfied,
    "happy": Icons.sentiment_very_satisfied,
  };

  final Map<String, Color> moodColors = const {
    "sad": Color(0xFFEF9A9A),
    "angry": Color(0xFFFFCC80),
    "neutral": Color(0xFFFFF176),
    "anxious": Color(0xFFA5D6A7),
    "happy": Color(0xFF66BB6A),
  };

  @override
  Widget build(BuildContext context) {
    if (userActivities.isEmpty) {
      return const Center(child: CircularProgressIndicator());
    }

    Map<String, int> moodCounts = {};
    for (var activity in userActivities) {
      final mood = activity['mood'];
      moodCounts[mood] = (moodCounts[mood] ?? 0) + 1;
    }

    return Container(
      margin: const EdgeInsets.symmetric(horizontal: 0, vertical: 0),
      padding: const EdgeInsets.symmetric(vertical: 0, horizontal: 0),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.center,
        children: [
          Row(
            mainAxisAlignment: MainAxisAlignment.center,
            crossAxisAlignment: CrossAxisAlignment.center,
            children: const [
                Icon(Icons.insights, color: Colors.teal, size: 22),
                SizedBox(width: 18),
                Text(
                  "Résumé de ton humeur",
                  style: TextStyle(
                    fontSize: 17,
                    fontWeight: FontWeight.w600,
                    color: AppColors.primaryGreen,
                  ),
                ),
            ],
          ),
          const SizedBox(height: 20),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceAround,
            children: moods.map((mood) {
              final count = moodCounts[mood] ?? 0;
              final color = moodColors[mood] ?? Colors.grey;

              return Column(
                mainAxisSize: MainAxisSize.min,
                children: [
                  Container(
                    padding: const EdgeInsets.all(6),
                    decoration: BoxDecoration(
                      color: color.withOpacity(0.2),
                      shape: BoxShape.circle,
                    ),
                    child: Icon(
                      moodIcons[mood],
                      color: color,
                      size: 28,
                    ),
                  ),
                  const SizedBox(height: 15),
                  Text(
                    mood.toUpperCase(),
                    style: const TextStyle(
                      fontSize: 11,
                      fontWeight: FontWeight.w600,
                      color: Colors.black87,
                    ),
                  ),
                  const SizedBox(height: 10),
                  Container(
                    padding: const EdgeInsets.symmetric(horizontal: 10, vertical: 4),
                    decoration: BoxDecoration(
                      color: color.withOpacity(0.15),
                      borderRadius: BorderRadius.circular(20),
                    ),
                    child: Text(
                      count.toString(),
                      style: TextStyle(
                        color: color,
                        fontWeight: FontWeight.bold,
                        fontSize: 13,
                      ),
                    ),
                  ),
                ],
              );
            }).toList(),
          ),
        ],
      ),
    );
  }
}
