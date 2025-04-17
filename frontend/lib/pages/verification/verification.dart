import 'package:Arly/core/style.dart';
import 'package:Arly/pages/detail/questionnaire.dart';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'package:shared_preferences/shared_preferences.dart';

import '../widgets/verification_background.dart';
import '../widgets/verification_form.dart';
import 'package:Arly/config.dart';

class VerificationPage extends StatefulWidget {
  final String email;
  const VerificationPage({super.key, required this.email});

  @override
  _VerificationPageState createState() => _VerificationPageState();
}

class _VerificationPageState extends State<VerificationPage> {
  final List<TextEditingController> _controllers =
      List.generate(6, (_) => TextEditingController());
  String _errorMessage = '';
  bool _isLoading = false;
  bool _isResending = false;

  Future<void> _verifyCode() async {
    final code = _controllers.map((controller) => controller.text).join();
    if (code.length != 6 || code.contains(RegExp(r'\D'))) {
      setState(() {
        _errorMessage = 'Veuillez entrer un code valide de 6 chiffres.';
      });
      return;
    }

    setState(() {
      _isLoading = true;
      _errorMessage = '';
    });

    try {
      final response = await http.post(
        Uri.parse('${HOST}/api/verify'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({
          'email': widget.email,
          'code': code,
        }),
      );

      if (response.statusCode == 200) {
        final responseData = jsonDecode(response.body);
        final token = responseData['token'];

        final prefs = await SharedPreferences.getInstance();
        await prefs.setString('authToken', token);

        Navigator.pushReplacement(
          context,
          MaterialPageRoute(builder: (context) => const QuestionnairePage()),
        );
      } else if (response.statusCode == 401) {
        setState(() {
          _errorMessage = 'Code de vérification invalide ou expiré.';
        });
      } else {
        final error = jsonDecode(response.body);
        setState(() {
          _errorMessage = error['message'] ?? 'Une erreur s\'est produite.';
        });
      }
    } catch (e) {
      setState(() {
        _errorMessage = 'Erreur de connexion: $e';
      });
    } finally {
      setState(() {
        _isLoading = false;
      });
    }
  }

  Future<void> _resendCode() async {
    setState(() {
      _isResending = true;
      _errorMessage = '';
    });

    try {
      final response = await http.post(
        Uri.parse('${HOST}/api/resend-code'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({'email': widget.email}),
      );

      if (response.statusCode == 200) {
        setState(() {
          _errorMessage = 'Code de vérification renvoyé avec succès.';
        });
      } else {
        final error = jsonDecode(response.body);
        setState(() {
          _errorMessage = error['message'] ?? 'Une erreur s\'est produite.';
        });
      }
    } catch (e) {
      setState(() {
        _errorMessage = 'Erreur réseau: $e';
      });
    } finally {
      setState(() {
        _isResending = false;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppColors.blackPrimary,
      body: Stack(
        children: [
          const VerificationBackground(),
          VerificationForm(
            email: widget.email,
            controllers: _controllers,
            errorMessage: _errorMessage,
            isLoading: _isLoading,
            isResending: _isResending,
            onVerify: _verifyCode,
            onResend: _resendCode,
          ),
        ],
      ),
    );
  }
}
