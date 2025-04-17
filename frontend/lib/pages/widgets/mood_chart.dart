import 'package:Arly/core/style.dart';
import 'package:flutter/material.dart';

class MoodChart extends StatelessWidget {
  final List<dynamic> userActivities;

  const MoodChart({
    super.key,
    required this.userActivities,
  });

  final Map<String, Color> moodColors = const {
    "bad": Color(0xFFEF9A9A),
    "sad": Color(0xFFFFCC80),
    "poker": Color(0xFFFFF176),
    "nice": Color(0xFFA5D6A7),
    "happy": Color(0xFF66BB6A),
  };

  IconData _getMoodIcon(String mood) {
    switch (mood) {
      case "bad":
        return Icons.sentiment_very_dissatisfied;
      case "sad":
        return Icons.sentiment_dissatisfied;
      case "poker":
        return Icons.sentiment_neutral;
      case "nice":
        return Icons.sentiment_satisfied;
      case "happy":
        return Icons.sentiment_very_satisfied;
      default:
        return Icons.sentiment_neutral;
    }
  }

  List<Widget> _getMoodBars() {
    if (userActivities.isEmpty) {
      return [const CircularProgressIndicator()];
    }

    final now = DateTime.now();
    final lastSevenDaysActivities = userActivities.where((activity) {
      final activityDate = DateTime.parse(activity['date']);
      return activityDate.isAfter(now.subtract(const Duration(days: 7))) &&
          activityDate.isBefore(now.add(const Duration(days: 1)));
    }).toList();

    lastSevenDaysActivities.sort((a, b) {
      final dateA = DateTime.parse(a['date']);
      final dateB = DateTime.parse(b['date']);
      return dateA.compareTo(dateB);
    });

    return lastSevenDaysActivities.map((activity) {
      final mood = activity['mood'];
      final date = DateTime.parse(activity['date']);
      final color = moodColors[mood] ?? Colors.grey;

      return Column(
        mainAxisSize: MainAxisSize.min,
        children: [
          Container(
            padding: const EdgeInsets.all(6),
            decoration: BoxDecoration(
              color: color.withOpacity(0.15),
              shape: BoxShape.circle,
            ),
            child: Icon(
              _getMoodIcon(mood),
              color: color,
              size: 26,
            ),
          ),
          const SizedBox(height: 10),
          Container(
            width: 18,
            height: 60,
            decoration: BoxDecoration(
              gradient: LinearGradient(
                colors: [
                  color.withOpacity(0.9),
                  color.withOpacity(0.6),
                ],
                begin: Alignment.topCenter,
                end: Alignment.bottomCenter,
              ),
              borderRadius: BorderRadius.circular(20),
              boxShadow: [
                BoxShadow(
                  color: color.withOpacity(0.3),
                  blurRadius: 6,
                  offset: const Offset(0, 3),
                ),
              ],
            ),
          ),
          const SizedBox(height: 8),
          Text(
            "${date.day}/${date.month}",
            style: const TextStyle(
              fontSize: 11,
              color: Colors.grey,
              fontWeight: FontWeight.w500,
            ),
          )
        ],
      );
    }).toList();
  }

  @override
  Widget build(BuildContext context) {
    return Container(
      margin: const EdgeInsets.symmetric(horizontal: 20, vertical: 16),
      padding: const EdgeInsets.symmetric(vertical: 24, horizontal: 20),
      decoration: BoxDecoration(
        color: AppColors.primaryWhite,
        borderRadius: BorderRadius.circular(28),
      ),
      child: Column(
        children: [
          Row(
            children: [
              const Icon(Icons.bar_chart, color: AppColors.tealPrimary, size: 22),
              const SizedBox(width: 8),
              Text(
                "Mes humeurs de la semaine",
                style: Theme.of(context).textTheme.titleMedium!.copyWith(
                  fontSize: 17,
                  fontWeight: FontWeight.w700,
                  color: AppColors.primaryGreen,
                ),
              ),
            ],
          ),
          const SizedBox(height: 28),
          Row(
            mainAxisAlignment: MainAxisAlignment.spaceAround,
            crossAxisAlignment: CrossAxisAlignment.end,
            children: _getMoodBars(),
          ),
          const SizedBox(height: 24),
          Divider(color: Colors.grey[300], thickness: 1),
          const SizedBox(height: 10),
          Text(
            _getDateRange(),
            style: Theme.of(context).textTheme.bodySmall!.copyWith(
              fontSize: 12,
              color: Colors.grey[600],
              fontWeight: FontWeight.w500,
            ),
          ),
        ],
      ),
    );
  }

  String _getMonthName(int month) {
    const months = [
      "Janv", "Fév", "Mars", "Avr", "Mai", "Juin",
      "Juil", "Août", "Sept", "Oct", "Nov", "Déc"
    ];
    return months[month - 1];
  }

  String _getDateRange() {
    if (userActivities.isEmpty) return "";
    final sortedActivities = List.from(userActivities)
      ..sort((a, b) {
        final dateA = DateTime.parse(a['date']);
        final dateB = DateTime.parse(b['date']);
        return dateA.compareTo(dateB);
      });

    final reversed = sortedActivities.reversed.toList();
    final firstDate = DateTime.parse(reversed.last['date']);
    final lastDate = DateTime.parse(reversed.first['date']);

    return "${firstDate.day} ${_getMonthName(firstDate.month)} - ${lastDate.day} ${_getMonthName(lastDate.month)}";
  }
}
