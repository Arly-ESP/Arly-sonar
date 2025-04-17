import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:Arly/pages/widgets/login_form.dart';

void main() {
  group('LoginForm Widget Tests', () {
    late TextEditingController emailController;
    late TextEditingController passwordController;
    late VoidCallback onLogin;
    late VoidCallback onRegister;

    setUp(() {
      emailController = TextEditingController();
      passwordController = TextEditingController();
      onLogin = () {};
      onRegister = () {};
    });

    testWidgets('displays all essential elements', (WidgetTester tester) async {
      await tester.pumpWidget(
        MaterialApp(
          home: Scaffold(
            body: LoginForm(
              emailController: emailController,
              passwordController: passwordController,
              errorMessage: '',
              isLoading: false,
              onLogin: onLogin,
              onRegister: onRegister,
            ),
          ),
        ),
      );

      expect(find.text('Connexion'), findsOneWidget);
      expect(
          find.byType(TextField), findsNWidgets(2)); // Email & Password fields
      expect(
          find.widgetWithText(ElevatedButton, 'Se connecter'), findsOneWidget);
      expect(find.widgetWithText(TextButton, 'S’enregistrer'), findsOneWidget);
    });

    testWidgets('displays error message when provided',
        (WidgetTester tester) async {
      await tester.pumpWidget(
        MaterialApp(
          home: Scaffold(
            body: LoginForm(
              emailController: emailController,
              passwordController: passwordController,
              errorMessage: 'Nom d’utilisateur ou mot de passe incorrect',
              isLoading: false,
              onLogin: onLogin,
              onRegister: onRegister,
            ),
          ),
        ),
      );

      expect(find.text('Nom d’utilisateur ou mot de passe incorrect'),
          findsOneWidget);
    });

    testWidgets('shows CircularProgressIndicator when isLoading is true',
        (WidgetTester tester) async {
      await tester.pumpWidget(
        MaterialApp(
          home: Scaffold(
            body: LoginForm(
              emailController: emailController,
              passwordController: passwordController,
              errorMessage: '',
              isLoading: true,
              onLogin: onLogin,
              onRegister: onRegister,
            ),
          ),
        ),
      );

      expect(find.byType(CircularProgressIndicator), findsOneWidget);
      expect(find.byType(ElevatedButton), findsNothing);
    });

    testWidgets('triggers onLogin callback when login button is tapped',
        (WidgetTester tester) async {
      bool loginTapped = false;

      await tester.pumpWidget(
        MaterialApp(
          home: Scaffold(
            body: LoginForm(
              emailController: emailController,
              passwordController: passwordController,
              errorMessage: '',
              isLoading: false,
              onLogin: () {
                loginTapped = true;
              },
              onRegister: onRegister,
            ),
          ),
        ),
      );

      await tester.tap(find.widgetWithText(ElevatedButton, 'Se connecter'));
      await tester.pump();

      expect(loginTapped, isTrue);
    });

    testWidgets('triggers onRegister callback when register button is tapped',
        (WidgetTester tester) async {
      bool registerTapped = false;

      await tester.pumpWidget(
        MaterialApp(
          home: Scaffold(
            body: LoginForm(
              emailController: emailController,
              passwordController: passwordController,
              errorMessage: '',
              isLoading: false,
              onLogin: onLogin,
              onRegister: () {
                registerTapped = true;
              },
            ),
          ),
        ),
      );

      await tester.tap(find.widgetWithText(TextButton, 'S’enregistrer'));
      await tester.pump();

      expect(registerTapped, isTrue);
    });
  });
}
