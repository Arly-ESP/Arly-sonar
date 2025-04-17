import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:Arly/pages/widgets/verification_form.dart';

void main() {
  group('VerificationForm Widget Tests', () {
    const email = 'test@example.com';
    final controllers = List.generate(6, (_) => TextEditingController());
    const errorMessage = 'Error message';
    bool isLoading = false;
    VoidCallback onVerify = () {};
    VoidCallback onResend = () {};

    testWidgets('displays all text fields', (WidgetTester tester) async {
      await tester.pumpWidget(MaterialApp(
        home: Scaffold(
          body: VerificationForm(
            email: email,
            controllers: controllers,
            errorMessage: '',
            isLoading: false,
            isResending: false,
            onVerify: onVerify,
            onResend: onResend,
          ),
        ),
      ));

      expect(find.text('Validation'), findsOneWidget);
      expect(find.text('Adresse email'), findsOneWidget);
      expect(find.text(email), findsOneWidget);
      expect(find.text('Renvoyer'), findsOneWidget);
      expect(find.text('Valider'), findsOneWidget);
    });

    testWidgets('displays error message when not empty', (WidgetTester tester) async {
      await tester.pumpWidget(MaterialApp(
        home: Scaffold(
          body: VerificationForm(
            email: email,
            controllers: controllers,
            errorMessage: errorMessage,
            isLoading: false,
            isResending: false,
            onVerify: onVerify,
            onResend: onResend,
          ),
        ),
      ));

      expect(find.text(errorMessage), findsOneWidget);
    });

    testWidgets('shows loading indicator when isLoading is true', (WidgetTester tester) async {
      isLoading = true;
      await tester.pumpWidget(MaterialApp(
        home: Scaffold(
          body: VerificationForm(
            email: email,
            controllers: controllers,
            errorMessage: '',
            isLoading: isLoading,
            isResending: false,
            onVerify: onVerify,
            onResend: onResend,
          ),
        ),
      ));

      expect(find.byType(CircularProgressIndicator), findsOneWidget);
      expect(find.widgetWithText(ElevatedButton, 'Valider'), findsNothing);
    });

    testWidgets('triggers onVerify callback when verify button is pressed', (WidgetTester tester) async {
      bool verifyCalled = false;
      onVerify = () {
        verifyCalled = true;
      };

      await tester.pumpWidget(MaterialApp(
        home: Scaffold(
          body: VerificationForm(
            email: email,
            controllers: controllers,
            errorMessage: '',
            isLoading: false,
            isResending: false,
            onVerify: onVerify,
            onResend: onResend,
          ),
        ),
      ));

      await tester.tap(find.widgetWithText(ElevatedButton, 'Valider'));
      await tester.pump();

      expect(verifyCalled, isTrue);
    });

    testWidgets('triggers onResend callback when resend button is pressed', (WidgetTester tester) async {
      bool resendCalled = false;
      onResend = () {
        resendCalled = true;
      };

      await tester.pumpWidget(MaterialApp(
        home: Scaffold(
          body: VerificationForm(
            email: email,
            controllers: controllers,
            errorMessage: '',
            isLoading: false,
            isResending: false,
            onVerify: onVerify,
            onResend: onResend,
          ),
        ),
      ));

      await tester.tap(find.text('Renvoyer'));
      await tester.pump();

      expect(resendCalled, isTrue);
    });

    testWidgets('focus moves correctly between text fields', (WidgetTester tester) async {
      await tester.pumpWidget(MaterialApp(
        home: Scaffold(
          body: VerificationForm(
            email: email,
            controllers: controllers,
            errorMessage: '',
            isLoading: false,
            isResending: false,
            onVerify: onVerify,
            onResend: onResend,
          ),
        ),
      ));

      await tester.enterText(find.byType(TextField).first, '1');
      await tester.pumpAndSettle();

      expect(controllers[0].text, '1');
      expect(find.byType(TextField).at(1), findsOneWidget);
    });
  });
}
