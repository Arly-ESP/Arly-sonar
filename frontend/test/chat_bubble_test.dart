// test/chat_bubble_test.dart
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:Arly/pages/widgets/chat_bubble.dart';

void main() {
  group('ChatBubble Widget', () {
    testWidgets('displays user message correctly without avatar', (WidgetTester tester) async {
      const messageText = "Hello from the user!";
      
      await tester.pumpWidget(
        const MaterialApp(
          home: Scaffold(
            body: ChatBubble(isUser: true, text: messageText),
          ),
        ),
      );
      
      expect(find.text(messageText), findsOneWidget);
      expect(find.byType(Image), findsNothing);
    });

    testWidgets('displays AI message correctly with avatar', (WidgetTester tester) async {
      const messageText = "Hello from the AI!";
      await tester.pumpWidget(
        const MaterialApp(
          home: Scaffold(
            body: ChatBubble(isUser: false, text: messageText),
          ),
        ),
      );
      
      expect(find.text(messageText), findsOneWidget);
      expect(find.byType(Image), findsOneWidget);
    });
  });
}
