import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:Arly/pages/widgets/message_input.dart';

void main() {
  group('MessageInput Widget Tests', () {
    late TextEditingController controller;

    setUp(() {
      controller = TextEditingController();
    });

    tearDown(() {
      controller.dispose();
    });

    testWidgets('renders correctly with initial state', (WidgetTester tester) async {
      await tester.pumpWidget(
        MaterialApp(
          home: Scaffold(
            body: MessageInput(
              controller: controller,
              isListening: false,
              isLoading: false,
              onMicTap: () {},
              onSendTap: () {},
            ),
          ),
        ),
      );

      expect(find.byType(TextField), findsOneWidget);
      expect(find.byType(IconButton), findsNWidgets(2)); // Mic and Send buttons
    });

    testWidgets('send button is disabled when loading', (WidgetTester tester) async {
      await tester.pumpWidget(
        MaterialApp(
          home: Scaffold(
            body: MessageInput(
              controller: controller,
              isListening: false,
              isLoading: true, // Send button should be disabled
              onMicTap: () {},
              onSendTap: () {},
            ),
          ),
        ),
      );

      final Finder sendButtonFinder = find.widgetWithIcon(IconButton, Icons.send);
      final IconButton sendButton = tester.widget(sendButtonFinder) as IconButton;

      expect(sendButton.onPressed, isNull); // Expect the button to be disabled
    });

    testWidgets('mic button changes color when listening', (WidgetTester tester) async {
      await tester.pumpWidget(
        MaterialApp(
          home: Scaffold(
            body: MessageInput(
              controller: controller,
              isListening: true,
              isLoading: false,
              onMicTap: () {},
              onSendTap: () {},
            ),
          ),
        ),
      );

      final Finder micButtonFinder = find.widgetWithIcon(IconButton, Icons.mic);
      final IconButton micButton = tester.widget(micButtonFinder) as IconButton;

      expect((micButton.icon as Icon).color, Colors.red);
    });

    testWidgets('mic button is green when not listening', (WidgetTester tester) async {
      await tester.pumpWidget(
        MaterialApp(
          home: Scaffold(
            body: MessageInput(
              controller: controller,
              isListening: false,
              isLoading: false,
              onMicTap: () {},
              onSendTap: () {},
            ),
          ),
        ),
      );

      final Finder micButtonFinder = find.widgetWithIcon(IconButton, Icons.mic_none);
      final IconButton micButton = tester.widget(micButtonFinder) as IconButton;

      expect((micButton.icon as Icon).color, Colors.green);
    });

    testWidgets('text input updates when typing', (WidgetTester tester) async {
      await tester.pumpWidget(
        MaterialApp(
          home: Scaffold(
            body: MessageInput(
              controller: controller,
              isListening: false,
              isLoading: false,
              onMicTap: () {},
              onSendTap: () {},
            ),
          ),
        ),
      );

      await tester.enterText(find.byType(TextField), 'Hello world');
      expect(controller.text, 'Hello world');
    });

    testWidgets('send button triggers onSendTap when pressed', (WidgetTester tester) async {
      bool wasSendTapped = false;

      await tester.pumpWidget(
        MaterialApp(
          home: Scaffold(
            body: MessageInput(
              controller: controller,
              isListening: false,
              isLoading: false,
              onMicTap: () {},
              onSendTap: () {
                wasSendTapped = true;
              },
            ),
          ),
        ),
      );

      await tester.tap(find.widgetWithIcon(IconButton, Icons.send));
      await tester.pump();

      expect(wasSendTapped, isTrue);
    });

    testWidgets('mic button triggers onMicTap when pressed', (WidgetTester tester) async {
      bool wasMicTapped = false;

      await tester.pumpWidget(
        MaterialApp(
          home: Scaffold(
            body: MessageInput(
              controller: controller,
              isListening: false,
              isLoading: false,
              onMicTap: () {
                wasMicTapped = true;
              },
              onSendTap: () {},
            ),
          ),
        ),
      );

      await tester.tap(find.widgetWithIcon(IconButton, Icons.mic_none));
      await tester.pump();

      expect(wasMicTapped, isTrue);
    });
  });
}
