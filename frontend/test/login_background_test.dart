import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:Arly/pages/widgets/login_background.dart';

void main() {
  testWidgets('LoginBackground displays the correct image', (WidgetTester tester) async {
    await tester.pumpWidget(
      const MaterialApp(
        home: Scaffold(
          body: Stack(
            children: [
              LoginBackground(),
            ],
          ),
        ),
      ),
    );

    // Verify that the background image is displayed
    final imageFinder = find.byType(Image);
    expect(imageFinder, findsOneWidget);

    final imageWidget = tester.widget<Image>(imageFinder);
    expect((imageWidget.image as AssetImage).assetName, 'assets/panda.jpg');
  });
}