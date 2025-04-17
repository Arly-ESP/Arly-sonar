import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:Arly/pages/widgets/register_form.dart';

void main() {
  group('RegisterForm Widget Tests', () {
    final emailController = TextEditingController();
    final firstnameController = TextEditingController();
    final lastnameController = TextEditingController();
    final passwordController = TextEditingController();
    final errorMessage = 'Error message';
    bool isLoading = false;
    VoidCallback onRegister = () {};
    VoidCallback onNavigateToLogin = () {};

    testWidgets('displays all form fields', (WidgetTester tester) async {
      await tester.pumpWidget(MaterialApp(
        home: Scaffold(
          body: RegisterForm(
            emailController: emailController,
            firstnameController: firstnameController,
            lastnameController: lastnameController,
            passwordController: passwordController,
            errorMessage: '',
            isLoading: false,
            onRegister: onRegister,
            onNavigateToLogin: onNavigateToLogin,
          ),
        ),
      ));

      expect(find.text('Adresse email'), findsOneWidget);
      expect(find.text('Prénom'), findsOneWidget);
      expect(find.text('Nom de famille'), findsOneWidget);
      expect(find.text('Mot de passe'), findsOneWidget);
    });

    testWidgets('displays error message when not empty', (WidgetTester tester) async {
      await tester.pumpWidget(MaterialApp(
        home: Scaffold(
          body: RegisterForm(
            emailController: emailController,
            firstnameController: firstnameController,
            lastnameController: lastnameController,
            passwordController: passwordController,
            errorMessage: errorMessage,
            isLoading: false,
            onRegister: onRegister,
            onNavigateToLogin: onNavigateToLogin,
          ),
        ),
      ));

      expect(find.text(errorMessage), findsOneWidget);
    });

    testWidgets('shows loading indicator when isLoading is true', (WidgetTester tester) async {
      isLoading = true;
      await tester.pumpWidget(MaterialApp(
        home: Scaffold(
          body: RegisterForm(
            emailController: emailController,
            firstnameController: firstnameController,
            lastnameController: lastnameController,
            passwordController: passwordController,
            errorMessage: '',
            isLoading: isLoading,
            onRegister: onRegister,
            onNavigateToLogin: onNavigateToLogin,
          ),
        ),
      ));

      expect(find.byType(CircularProgressIndicator), findsOneWidget);
      expect(find.widgetWithText(ElevatedButton, 'S’enregistrer'), findsNothing);
    });

    testWidgets('triggers onRegister callback when register button is pressed', (WidgetTester tester) async {
      bool registerCalled = false;
      onRegister = () {
        registerCalled = true;
      };

      await tester.pumpWidget(MaterialApp(
        home: Scaffold(
          body: RegisterForm(
            emailController: emailController,
            firstnameController: firstnameController,
            lastnameController: lastnameController,
            passwordController: passwordController,
            errorMessage: '',
            isLoading: false,
            onRegister: onRegister,
            onNavigateToLogin: onNavigateToLogin,
          ),
        ),
      ));

      await tester.tap(find.widgetWithText(ElevatedButton, 'S’enregistrer'));
      await tester.pump();

      expect(registerCalled, isTrue);
    });

    testWidgets('triggers onNavigateToLogin callback when navigate to login button is pressed', (WidgetTester tester) async {
      bool navigateToLoginCalled = false;
      onNavigateToLogin = () {
        navigateToLoginCalled = true;
      };

      await tester.pumpWidget(MaterialApp(
        home: Scaffold(
          body: RegisterForm(
            emailController: emailController,
            firstnameController: firstnameController,
            lastnameController: lastnameController,
            passwordController: passwordController,
            errorMessage: '',
            isLoading: false,
            onRegister: onRegister,
            onNavigateToLogin: onNavigateToLogin,
          ),
        ),
      ));

      await tester.tap(find.text('Se connecter'));
      await tester.pump();

      expect(navigateToLoginCalled, isTrue);
    });
  });
}
