import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:Arly/pages/widgets/mood_summary.dart';

void main() {
  group('MoodSummary Widget Tests', () {
    testWidgets('displays CircularProgressIndicator when userActivities is empty', (WidgetTester tester) async {
      await tester.pumpWidget(MaterialApp(
        home: Scaffold(
          body: MoodSummary(userActivities: []),
        ),
      ));

      expect(find.byType(CircularProgressIndicator), findsOneWidget);
    });

    testWidgets('displays mood counts correctly', (WidgetTester tester) async {
  final userActivities = [
    {'mood': 'happy'},
    {'mood': 'sad'},
    {'mood': 'happy'},
    {'mood': 'neutral'},
  ];

  await tester.pumpWidget(MaterialApp(
    home: Scaffold(
      body: MoodSummary(userActivities: userActivities),
    ),
  ));

  expect(find.text('HAPPY'), findsOneWidget);
  expect(find.text('2'), findsOneWidget);
  expect(find.text('SAD'), findsOneWidget);
  expect(find.text('1'), findsNWidgets(2));
  expect(find.text('0'), findsNWidgets(2));
});


    testWidgets('handles missing moods gracefully', (WidgetTester tester) async {
      final userActivities = [
        {'mood': 'happy'},
        {'mood': 'unknown'},
      ];

      await tester.pumpWidget(MaterialApp(
        home: Scaffold(
          body: MoodSummary(userActivities: userActivities),
        ),
      ));

      expect(find.text('HAPPY'), findsOneWidget);
      expect(find.text('1'), findsOneWidget);
      expect(find.text('UNKNOWN'), findsNothing);
    });
  });
}
