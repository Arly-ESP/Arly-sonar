import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:Arly/pages/widgets/verification_background.dart';

void main() {
  group('VerificationBackground Widget Tests', () {
    testWidgets('displays the image correctly', (WidgetTester tester) async {
      await tester.pumpWidget(const MaterialApp(
        home: Scaffold(
          body: Stack(
            children: [
              VerificationBackground(),
            ],
          ),
        ),
      ));

      // Check if the Image widget is present
      expect(find.byType(Image), findsOneWidget);

      // Check if the image asset is loaded correctly
      final image = tester.widget<Image>(find.byType(Image));
      expect(image.image, const AssetImage('assets/panda.jpg'));

      // Check if the image fits the entire area
      final imageBoxFit = tester.widget<Image>(find.byType(Image)).fit;
      expect(imageBoxFit, BoxFit.cover);
    });

    testWidgets('covers the entire area', (WidgetTester tester) async {
      await tester.pumpWidget(const MaterialApp(
        home: Scaffold(
          body: Stack(
            children: [
              VerificationBackground(),
            ],
          ),
        ),
      ));

      // Check if the Positioned.fill widget is present
      expect(find.byType(Positioned), findsOneWidget);

      // Check if the Positioned.fill widget fills the entire area
      final positioned = tester.widget<Positioned>(find.byType(Positioned));
      expect(positioned.left, 0.0);
      expect(positioned.top, 0.0);
      expect(positioned.right, 0.0);
      expect(positioned.bottom, 0.0);
    });
  });
}
