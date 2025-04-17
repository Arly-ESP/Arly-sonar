import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:Arly/pages/widgets/activity_map.dart';

void main() {
  group('ActivityMap Widget Tests', () {
    testWidgets('displays loading indicator when userActivities is empty', (WidgetTester tester) async {
      await tester.pumpWidget(const MaterialApp(
        home: Scaffold(
          body: ActivityMap(userActivities: []),
        ),
      ));
      
      expect(find.byType(CircularProgressIndicator), findsOneWidget);
    });

    testWidgets('displays grid items based on userActivities length', (WidgetTester tester) async {
      final userActivities = [
        {'date': '2024-03-01', 'message_count': 2},
        {'date': '2024-03-02', 'message_count': 5},
        {'date': '2024-03-03', 'message_count': 0},
      ];
      
      await tester.pumpWidget(MaterialApp(
        home: Scaffold(
          body: ActivityMap(userActivities: userActivities),
        ),
      ));

      expect(find.byType(Container), findsNWidgets(userActivities.length));
    });

    testWidgets('applies correct colors based on message count', (WidgetTester tester) async {
      final userActivities = [
        {'date': '2024-03-01', 'message_count': 6}, // Green
        {'date': '2024-03-02', 'message_count': 4}, // Light Green
        {'date': '2024-03-03', 'message_count': 2}, // Lime
        {'date': '2024-03-04', 'message_count': 1}, // Yellow
        {'date': '2024-03-05', 'message_count': 0}, // Grey
      ];
      
      await tester.pumpWidget(MaterialApp(
        home: Scaffold(
          body: ActivityMap(userActivities: userActivities),
        ),
      ));

      final containers = tester.widgetList<Container>(find.byType(Container)).toList();
      expect((containers[0].decoration as BoxDecoration).color, Colors.green);
      expect((containers[1].decoration as BoxDecoration).color, Colors.lightGreen);
      expect((containers[2].decoration as BoxDecoration).color, Colors.lime);
      expect((containers[3].decoration as BoxDecoration).color, Colors.yellow);
      expect((containers[4].decoration as BoxDecoration).color, Colors.grey[300]);
    });
  });
}
