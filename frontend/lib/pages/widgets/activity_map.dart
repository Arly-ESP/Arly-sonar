import 'package:flutter/material.dart';

class ActivityMap extends StatelessWidget {
  final List<dynamic> userActivities;

  const ActivityMap({
    super.key,
    required this.userActivities,
  });

  Color _getMessageCountColor(int messageCount) {
    if (messageCount > 5) return const Color(0xFFB3E5FC); // pastel blue
    if (messageCount > 3) return const Color(0xFFC8E6C9); // pastel green
    if (messageCount > 1) return const Color(0xFFFFF9C4); // pastel yellow
    if (messageCount == 1) return const Color(0xFFFFE0B2); // pastel peach
    return const Color(0xFFFFF8E1); // light cream
  }

  @override
  Widget build(BuildContext context) {
    if (userActivities.isEmpty) {
      return const Center(child: CircularProgressIndicator());
    }

    final sortedActivities = List.from(userActivities)
      ..sort((a, b) {
        final dateA = DateTime.parse(a['date']);
        final dateB = DateTime.parse(b['date']);
        return dateA.compareTo(dateB);
      });

    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        Padding(
          padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 12),
          child: Row(
            children: [
              const Icon(Icons.place, color: Colors.teal),
              const SizedBox(width: 8),
              Text(
                "Carte d’activités",
                style: Theme.of(context).textTheme.titleMedium!.copyWith(
                  fontWeight: FontWeight.bold,
                  color: Colors.black87,
                ),
              ),
            ],
          ),
        ),
        Padding(
          padding: const EdgeInsets.symmetric(horizontal: 16),
          child: GridView.builder(
            shrinkWrap: true,
            physics: const NeverScrollableScrollPhysics(),
            itemCount: sortedActivities.length,
            gridDelegate: const SliverGridDelegateWithFixedCrossAxisCount(
              crossAxisCount: 7,
              mainAxisSpacing: 8,
              crossAxisSpacing: 8,
              childAspectRatio: 1,
            ),
            itemBuilder: (context, index) {
              final activity = sortedActivities[index];
              final messageCount = activity['message_count'];
              final color = _getMessageCountColor(messageCount);

              return AnimatedContainer(
                duration: const Duration(milliseconds: 300),
                decoration: BoxDecoration(
                  color: color,
                  borderRadius: BorderRadius.circular(12),
                  boxShadow: [
                    BoxShadow(
                      color: Colors.black12,
                      blurRadius: 4,
                      offset: Offset(2, 2),
                    ),
                  ],
                ),
              );
            },
          ),
        ),
        const SizedBox(height: 24),
      ],
    );
  }
}
