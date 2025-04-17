import 'package:Arly/pages/Chat/chat_page.dart';
import 'package:Arly/pages/Login/login_page.dart';
import 'package:Arly/pages/Register/register_page.dart';
import 'package:Arly/pages/detail/questionnaire.dart';
import 'package:Arly/pages/verification/verification.dart';
import 'package:features_tour/features_tour.dart';
import 'package:flutter/material.dart';
import 'package:Arly/pages/home/home_page.dart';
import 'package:Arly/pages/payments/subscription_page.dart';
import 'package:Arly/core/style.dart';

void main() {
  FeaturesTour.setGlobalConfig(
    force: null,
    predialogConfig: PredialogConfig(
      enabled: false,
    ),
    childConfig: ChildConfig(
      backgroundColor: AppColors.primaryWhite,
      isAnimateChild: true,
      zoomScale: 1.05,
      animationDuration: const Duration(milliseconds: 500),
    ),
    skipConfig: SkipConfig(
      text: 'Passer',
    ),
    nextConfig: NextConfig(text: 'Suivant'),
    doneConfig: DoneConfig(text: 'Terminer'),
    introduceConfig: IntroduceConfig(
      backgroundColor: AppColors.textPrimary.withOpacity(0.84),
    ),
  );

  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: "Arly",
      debugShowCheckedModeBanner: false,
      initialRoute: '/home',
      routes: {
        '/home': (context) => const HomePage(),
        '/login': (context) => const LoginPage(),
        '/register': (context) => const RegisterPage(),
        '/chat': (context) => const ChatPage(),
        '/questions': (context) => const QuestionnairePage(),
        '/verification': (context) => const VerificationPage(
              email: '',
            ),
        '/subscription': (context) => const SubscriptionPage(),
      },
    );
  }
}
