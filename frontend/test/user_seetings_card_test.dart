import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:Arly/pages/widgets/user_seetings_card.dart';

void main() {
  group('UserSettingsCard Widget Tests', () {
    const name = 'John Doe';
    const email = 'john.doe@example.com';
    VoidCallback onTakeQuestionnaire = () {};
    VoidCallback onContact = () {};
    VoidCallback onDeleteAccount = () {};
    VoidCallback onLogout = () {};

    testWidgets('displays user information correctly', (WidgetTester tester) async {
      await tester.pumpWidget(MaterialApp(
        home: Scaffold(
          body: UserSettingsCard(
            name: name,
            email: email,
            onTakeQuestionnaire: onTakeQuestionnaire,
            onContact: onContact,
            onDeleteAccount: onDeleteAccount,
            onLogout: onLogout,
          ),
        ),
      ));

      expect(find.text('Mon compte'), findsOneWidget);
      expect(find.text('Nom d\'utilisateur : $name'), findsOneWidget);
      expect(find.text('Email : $email'), findsOneWidget);
    });

    testWidgets('displays subscription information correctly', (WidgetTester tester) async {
      await tester.pumpWidget(MaterialApp(
        home: Scaffold(
          body: UserSettingsCard(
            name: name,
            email: email,
            onTakeQuestionnaire: onTakeQuestionnaire,
            onContact: onContact,
            onDeleteAccount: onDeleteAccount,
            onLogout: onLogout,
          ),
        ),
      ));

      expect(find.text('Mon abonnement'), findsOneWidget);
      expect(find.text('Renouvellement automatique'), findsOneWidget);
      expect(find.text('Type d\'abonnement : Premium'), findsOneWidget);
      expect(find.text('Jours restant avant la fin : 46'), findsOneWidget);
    });

    testWidgets('displays general settings correctly', (WidgetTester tester) async {
      await tester.pumpWidget(MaterialApp(
        home: Scaffold(
          body: UserSettingsCard(
            name: name,
            email: email,
            onTakeQuestionnaire: onTakeQuestionnaire,
            onContact: onContact,
            onDeleteAccount: onDeleteAccount,
            onLogout: onLogout,
          ),
        ),
      ));

      expect(find.text('Général'), findsOneWidget);
      expect(find.text('Reprendre le questionnaire'), findsOneWidget);
      expect(find.text('Nous contacter'), findsOneWidget);
      expect(find.text('Supprimer mon compte'), findsOneWidget);
      expect(find.text('Se déconnecter'), findsOneWidget);
    });

    testWidgets('triggers onTakeQuestionnaire callback when tapped', (WidgetTester tester) async {
      bool takeQuestionnaireCalled = false;
      onTakeQuestionnaire = () {
        takeQuestionnaireCalled = true;
      };

      await tester.pumpWidget(MaterialApp(
        home: Scaffold(
          body: UserSettingsCard(
            name: name,
            email: email,
            onTakeQuestionnaire: onTakeQuestionnaire,
            onContact: onContact,
            onDeleteAccount: onDeleteAccount,
            onLogout: onLogout,
          ),
        ),
      ));

      await tester.tap(find.text('Reprendre le questionnaire'));
      await tester.pump();

      expect(takeQuestionnaireCalled, isTrue);
    });

    testWidgets('triggers onContact callback when tapped', (WidgetTester tester) async {
      bool contactCalled = false;
      onContact = () {
        contactCalled = true;
      };

      await tester.pumpWidget(MaterialApp(
        home: Scaffold(
          body: UserSettingsCard(
            name: name,
            email: email,
            onTakeQuestionnaire: onTakeQuestionnaire,
            onContact: onContact,
            onDeleteAccount: onDeleteAccount,
            onLogout: onLogout,
          ),
        ),
      ));

      await tester.tap(find.text('Nous contacter'));
      await tester.pump();

      expect(contactCalled, isTrue);
    });

    testWidgets('triggers onDeleteAccount callback when tapped', (WidgetTester tester) async {
      bool deleteAccountCalled = false;
      onDeleteAccount = () {
        deleteAccountCalled = true;
      };

      await tester.pumpWidget(MaterialApp(
        home: Scaffold(
          body: UserSettingsCard(
            name: name,
            email: email,
            onTakeQuestionnaire: onTakeQuestionnaire,
            onContact: onContact,
            onDeleteAccount: onDeleteAccount,
            onLogout: onLogout,
          ),
        ),
      ));

      await tester.tap(find.text('Supprimer mon compte'));
      await tester.pump();

      expect(deleteAccountCalled, isTrue);
    });

    testWidgets('triggers onLogout callback when tapped', (WidgetTester tester) async {
      bool logoutCalled = false;
      onLogout = () {
        logoutCalled = true;
      };

      await tester.pumpWidget(MaterialApp(
        home: Scaffold(
          body: UserSettingsCard(
            name: name,
            email: email,
            onTakeQuestionnaire: onTakeQuestionnaire,
            onContact: onContact,
            onDeleteAccount: onDeleteAccount,
            onLogout: onLogout,
          ),
        ),
      ));

      await tester.tap(find.text('Se déconnecter'));
      await tester.pump();

      expect(logoutCalled, isTrue);
    });

    testWidgets('displays the switch correctly', (WidgetTester tester) async {
      await tester.pumpWidget(MaterialApp(
        home: Scaffold(
          body: UserSettingsCard(
            name: name,
            email: email,
            onTakeQuestionnaire: onTakeQuestionnaire,
            onContact: onContact,
            onDeleteAccount: onDeleteAccount,
            onLogout: onLogout,
          ),
        ),
      ));

      final switchWidget = tester.widget<Switch>(find.byType(Switch));
      expect(switchWidget.value, isTrue);
    });
  });
}
