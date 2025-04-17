import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:Arly/pages/widgets/chat_app_bar.dart';

void main() {
  testWidgets('ChatAppBar renders correctly and updates theme', (WidgetTester tester) async {
    String selectedTheme = "famille";
    String newTheme = "";
    
    await tester.pumpWidget(
      MaterialApp(
        home: Scaffold(
          appBar: ChatAppBar(
            selectedTheme: selectedTheme,
            onThemeChanged: (String value) {
              newTheme = value;
            },
          ),
        ),
      ),
    );

    // Verify initial title and dropdown presence
    expect(find.text("Thème :"), findsOneWidget);
    expect(find.text("famille"), findsOneWidget);
    expect(find.byType(DropdownButton<String>), findsOneWidget);
    expect(find.text("Quête : discuter au moins 5 minutes"), findsOneWidget);

    // Open the dropdown menu
    await tester.tap(find.byType(DropdownButton<String>));
    await tester.pumpAndSettle();

    // Verify dropdown items
    expect(find.text("amis"), findsOneWidget);
    expect(find.text("santé"), findsOneWidget);
    expect(find.text("ennemis"), findsOneWidget);
    expect(find.text("méditations"), findsOneWidget);

    // Select a new theme
    await tester.tap(find.text("santé"));
    await tester.pumpAndSettle();

    // Verify the callback was triggered
    expect(newTheme, "santé");
  });
}
