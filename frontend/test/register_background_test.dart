import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:Arly/pages/widgets/register_background.dart';

void main() {
  group('RegisterBackground Widget Tests', () {
    testWidgets('displays the image correctly', (WidgetTester tester) async {
      await tester.pumpWidget(const MaterialApp(
        home: Scaffold(
          body: Stack(
            children: [
              RegisterBackground(),
            ],
          ),
        ),
      ));
      expect(find.byType(Image), findsOneWidget);

      final image = tester.widget<Image>(find.byType(Image));
      expect(image.image, const AssetImage('assets/panda.jpg'));

      final imageBoxFit = tester.widget<Image>(find.byType(Image)).fit;
      expect(imageBoxFit, BoxFit.cover);
    });

    testWidgets('covers the entire area', (WidgetTester tester) async {
      await tester.pumpWidget(const MaterialApp(
        home: Scaffold(
          body: Stack(
            children: [
              RegisterBackground(),
            ],
          ),
        ),
      ));

      expect(find.byType(Positioned), findsOneWidget);

      final positioned = tester.widget<Positioned>(find.byType(Positioned));
      expect(positioned.left, 0.0);
      expect(positioned.top, 0.0);
      expect(positioned.right, 0.0);
      expect(positioned.bottom, 0.0);
    });
  });
}

