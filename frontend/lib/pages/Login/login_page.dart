import 'package:Arly/core/style.dart';
import 'package:Arly/pages/home/home_page.dart';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'package:shared_preferences/shared_preferences.dart';

import '../widgets/login_background.dart';
import '../widgets/login_form.dart';
import 'package:Arly/config.dart';

class LoginPage extends StatefulWidget {
  const LoginPage({super.key});

  @override
  _LoginPageState createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  final TextEditingController _emailController = TextEditingController();
  final TextEditingController _passwordController = TextEditingController();
  String _errorMessage = '';
  bool _isLoading = false;

  Future<void> _login() async {
    final email = _emailController.text.trim();
    final password = _passwordController.text.trim();

    if (email.isEmpty || password.isEmpty) {
      setState(() {
        _errorMessage = 'Veuillez remplir tous les champs.';
      });
      return;
    }

    setState(() {
      _isLoading = true;
      _errorMessage = '';
    });

    final Map<String, dynamic> requestBody = {
      "email": email,
      "password": password,
    };

    try {
      final response = await http.post(
        Uri.parse('${HOST}/api/login'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode(requestBody),
      );

      if (response.statusCode == 200 || response.statusCode == 201) {
        final responseData = jsonDecode(response.body);

        if (responseData['token'] != null) {
          final token = responseData['token'];
          SharedPreferences prefs = await SharedPreferences.getInstance();
          await prefs.setString('authToken', token);

          Navigator.pushReplacement(
            context,
            MaterialPageRoute(builder: (context) => const HomePage()),
          );
        } else {
          setState(() {
            _errorMessage = 'La réponse de l’API est invalide.';
          });
        }
      } else {
        final errorResponse = jsonDecode(response.body);
        setState(() {
          _errorMessage = errorResponse['error'] ??
              'Erreur lors de la connexion. Veuillez réessayer.';
        });
      }
    } catch (e) {
      setState(() {
        _errorMessage =
            'Erreur réseau ou problème de serveur. Veuillez réessayer.';
      });
    } finally {
      setState(() {
        _isLoading = false;
      });
    }
  }

  void _navigateToRegister() {
    Navigator.pushNamed(context, '/register');
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppColors.blackPrimary,
      body: Stack(
        children: [
          const LoginBackground(),
          LoginForm(
            emailController: _emailController,
            passwordController: _passwordController,
            errorMessage: _errorMessage,
            isLoading: _isLoading,
            onLogin: _login,
            onRegister: _navigateToRegister,
          ),
        ],
      ),
    );
  }
}
