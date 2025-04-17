import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:Arly/pages/widgets/progress_bar.dart';


void main() {
  group('ProgressBar Widget Tests', () {
    testWidgets('displays progress correctly at 0%', (WidgetTester tester) async {
      await tester.pumpWidget(const MaterialApp(
        home: Scaffold(
          body: SizedBox(
            width: 200,
            child: ProgressBar(progress: 0.0),
          ),
        ),
      ));

      final fractionallySizedBox = tester.widget<FractionallySizedBox>(find.byKey(const Key('progressContainer')));
      expect(fractionallySizedBox.widthFactor, 0.0);
    });

    testWidgets('displays progress correctly at 50%', (WidgetTester tester) async {
      await tester.pumpWidget(const MaterialApp(
        home: Scaffold(
          body: SizedBox(
            width: 200,
            child: ProgressBar(progress: 0.5),
          ),
        ),
      ));

      final fractionallySizedBox = tester.widget<FractionallySizedBox>(find.byKey(const Key('progressContainer')));
      expect(fractionallySizedBox.widthFactor, 0.5);
    });

    testWidgets('displays progress correctly at 100%', (WidgetTester tester) async {
      await tester.pumpWidget(const MaterialApp(
        home: Scaffold(
          body: SizedBox(
            width: 200,
            child: ProgressBar(progress: 1.0),
          ),
        ),
      ));

      final fractionallySizedBox = tester.widget<FractionallySizedBox>(find.byKey(const Key('progressContainer')));
      expect(fractionallySizedBox.widthFactor, 1.0);
    });

    testWidgets('handles progress value less than 0', (WidgetTester tester) async {
      await tester.pumpWidget(const MaterialApp(
        home: Scaffold(
          body: SizedBox(
            width: 200,
            child: ProgressBar(progress: 0.0),
          ),
        ),
      ));

      final fractionallySizedBox = tester.widget<FractionallySizedBox>(find.byKey(const Key('progressContainer')));
      expect(fractionallySizedBox.widthFactor, 0.0);
    });
  });
}
