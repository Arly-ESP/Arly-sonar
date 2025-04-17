import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:Arly/pages/widgets/seeting_app_bar.dart';

void main() {
  group('SettingsAppBar Widget Tests', () {
    testWidgets('displays the image correctly', (WidgetTester tester) async {
      await tester.pumpWidget(const MaterialApp(
        home: Scaffold(
          appBar: SettingsAppBar(),
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

    testWidgets('applies the semi-transparent black overlay', (WidgetTester tester) async {
      await tester.pumpWidget(const MaterialApp(
        home: Scaffold(
          appBar: SettingsAppBar(),
        ),
      ));

      // Check if the Container with the overlay is present
      expect(find.byType(Container), findsWidgets);

      // Check if the overlay color is correct
      final container = tester.widgetList<Container>(find.byType(Container)).firstWhere(
        (c) => c.color == Colors.black.withOpacity(0.4),
        orElse: () => throw Exception('Overlay not found'),
      );
      expect(container.color, Colors.black.withOpacity(0.4));
    });

    testWidgets('displays the text correctly', (WidgetTester tester) async {
      await tester.pumpWidget(const MaterialApp(
        home: Scaffold(
          appBar: SettingsAppBar(),
        ),
      ));

      // Check if the Text widget is present
      expect(find.text("Configuration"), findsOneWidget);

      // Check if the text style is correct
      final text = tester.widget<Text>(find.text("Configuration"));
      expect(text.style?.color, Colors.white);
      expect(text.style?.fontSize, 24);
      expect(text.style?.fontWeight, FontWeight.bold);
    });

    testWidgets('has the correct preferred size', (WidgetTester tester) async {
      await tester.pumpWidget(const MaterialApp(
        home: Scaffold(
          appBar: SettingsAppBar(),
        ),
      ));

      // Check if the AppBar has the preferred size
      final appBar = tester.widget<SettingsAppBar>(find.byType(SettingsAppBar));
      expect(appBar.preferredSize, const Size.fromHeight(100.0));
    });
  });
}
