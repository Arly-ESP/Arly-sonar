import 'package:Arly/core/style.dart';
import 'package:flutter/material.dart';
import 'package:Arly/core/utils.dart';
import 'package:Arly/core/buttons.dart';

class LoginForm extends StatefulWidget {
  final TextEditingController emailController;
  final TextEditingController passwordController;
  final String errorMessage;
  final bool isLoading;
  final VoidCallback onLogin;
  final VoidCallback onRegister;

  const LoginForm({
    super.key,
    required this.emailController,
    required this.passwordController,
    required this.errorMessage,
    required this.isLoading,
    required this.onLogin,
    required this.onRegister,
  });

  @override
  State<LoginForm> createState() => _LoginFormState();
}

class _LoginFormState extends State<LoginForm> {
  bool isPasswordVisible = false;

  @override
  Widget build(BuildContext context) {
    return Align(
      alignment: Alignment.bottomCenter,
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Container(
          padding: const EdgeInsets.symmetric(vertical: 30, horizontal: 25),
          decoration: BoxDecoration(
            color: AppColors.primaryWhite,
            borderRadius: BorderRadius.circular(25),
          ),
          child: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              const Text(
                'Connexion',
                style: TextStyle(
                  fontSize: 24,
                  fontWeight: FontWeight.bold,
                ),
              ),
              const SizedBox(height: 25),
              textFieldEmail(),
              const SizedBox(height: 20),
              buildTextFieldPassword(),
              largeSpacer(),
              if (widget.errorMessage.isNotEmpty)
                Padding(
                  padding: const EdgeInsets.only(bottom: 10),
                  child: Text(
                    "Nom d’utilisateur ou mot de passe incorrect",
                    style: const TextStyle(color: AppColors.primaryRed),
                  ),
                ),
              // Login Button
              widget.isLoading
                  ? const CircularProgressIndicator()
                  : SizedBox(
                      width: double.infinity,
                      child: loginButton(
                        text: 'Se connecter',
                        onPressed: widget.onLogin,
                      ),
                    ),
              smallSpacer(),
              SizedBox(
                width: double.infinity,
                child: RegisterButton(
                  text: 'S’enregistrer',
                  onPressed: widget.onRegister,
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }

  TextField buildTextFieldPassword() {
    return TextField(
      controller: widget.passwordController,
      obscureText: !isPasswordVisible,
      decoration: InputDecoration(
        labelText: 'Mot de passe',
        hintText: 'Entrer mon mot de passe',
        floatingLabelStyle: TextStyle(color: Colors.green), // Label focus
        suffixIcon: IconButton(
          icon: Icon(
            isPasswordVisible ? Icons.visibility : Icons.visibility_off,
            color: Colors.green, // Optionnel : couleur de l’icône
          ),
          onPressed: () {
            setState(() {
              isPasswordVisible = !isPasswordVisible;
            });
          },
        ),
        enabledBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(10),
          borderSide: BorderSide(
            color: Colors.grey,
          ),
        ),
        focusedBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(10),
          borderSide: BorderSide(
            color: Colors.green,
            width: 2.0,
          ),
        ),
      ),
    );
  }

  TextField textFieldEmail() {
    return TextField(
      controller: widget.emailController,
      keyboardType: TextInputType.emailAddress,
      decoration: InputDecoration(
        labelText: 'Adresse email',
        floatingLabelStyle:
            TextStyle(color: Colors.green), // Couleur du label quand focus
        hintText: 'Entrer mon adresse email',
        enabledBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(10),
          borderSide: BorderSide(
            color: Colors.grey, // couleur du border non focus
          ),
        ),
        focusedBorder: OutlineInputBorder(
          borderRadius: BorderRadius.circular(10),
          borderSide: BorderSide(
            color: Colors.green, // couleur du border quand focus
            width: 2.0, // épaisseur si tu veux accentuer
          ),
        ),
      ),
    );
  }
}
