import 'package:Arly/core/style.dart';
import 'package:Arly/pages/verification/verification.dart';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'package:shared_preferences/shared_preferences.dart';

import '../widgets/register_background.dart';
import '../widgets/register_form.dart';
import 'package:Arly/config.dart';

class RegisterPage extends StatefulWidget {
  const RegisterPage({super.key});

  @override
  _RegisterPageState createState() => _RegisterPageState();
}

class _RegisterPageState extends State<RegisterPage> {
  final TextEditingController _emailController = TextEditingController();
  final TextEditingController _firstnameController = TextEditingController();
  final TextEditingController _lastnameController = TextEditingController();
  final TextEditingController _passwordController = TextEditingController();
  String _errorMessage = '';
  bool _isLoading = false;

  Future<void> _register() async {
    final email = _emailController.text.trim();
    final firstname = _firstnameController.text.trim();
    final lastname = _lastnameController.text.trim();
    final password = _passwordController.text.trim();

    if (email.isEmpty ||
        firstname.isEmpty ||
        lastname.isEmpty ||
        password.isEmpty) {
      setState(() {
        _errorMessage = 'Tous les champs sont obligatoires.';
      });
      return;
    }

    if (!RegExp(r'^[^@]+@[^@]+\.[^@]+').hasMatch(email)) {
      setState(() {
        _errorMessage = 'Adresse email invalide.';
      });
      return;
    }

    if (password.length < 6) {
      setState(() {
        _errorMessage = 'Le mot de passe doit contenir au moins 6 caractères.';
      });
      return;
    }

    final Map<String, dynamic> requestBody = {
      "email": email,
      "first_name": firstname,
      "last_name": lastname,
      "password": password
    };

    setState(() {
      _isLoading = true;
      _errorMessage = '';
    });

    try {
      final response = await http.post(
        Uri.parse('${HOST}/api/register'),
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
            MaterialPageRoute(
              builder: (context) => VerificationPage(email: email),
            ),
          );
        } else {
          setState(() {
            _errorMessage = 'Token non reçu, vérifiez votre réponse API.';
          });
        }
      } else {
        final errorResponse = jsonDecode(response.body);
        setState(() {
          _errorMessage = errorResponse['error'] ??
              'Erreur lors de l\'enregistrement. Veuillez réessayer.';
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

  void _navigateToLogin() {
    Navigator.pushNamed(context, '/login');
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppColors.blackPrimary,
      body: Stack(
        children: [
          const RegisterBackground(),
          RegisterForm(
            emailController: _emailController,
            firstnameController: _firstnameController,
            lastnameController: _lastnameController,
            passwordController: _passwordController,
            errorMessage: _errorMessage,
            isLoading: _isLoading,
            onRegister: _register,
            onNavigateToLogin: _navigateToLogin,
          ),
        ],
      ),
    );
  }
}
