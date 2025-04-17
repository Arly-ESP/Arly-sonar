import 'package:Arly/core/style.dart';
import 'package:flutter/material.dart';
import 'package:Arly/core/textbox.dart';
import 'package:Arly/core/utils.dart';

import '../../core/buttons.dart';

class RegisterForm extends StatefulWidget {
  final TextEditingController emailController;
  final TextEditingController firstnameController;
  final TextEditingController lastnameController;
  final TextEditingController passwordController;
  final String errorMessage;
  final bool isLoading;
  final VoidCallback onRegister;
  final VoidCallback onNavigateToLogin;

  const RegisterForm({
    super.key,
    required this.emailController,
    required this.firstnameController,
    required this.lastnameController,
    required this.passwordController,
    required this.errorMessage,
    required this.isLoading,
    required this.onRegister,
    required this.onNavigateToLogin,
  });

  @override
  State<RegisterForm> createState() => _RegisterFormState();
}

class _RegisterFormState extends State<RegisterForm> {
  bool _obscurePassword = true;

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
          child: SingleChildScrollView(
            child: Column(
              mainAxisSize: MainAxisSize.min,
              children: [
                const Text(
                  'S’enregistrer',
                  style: TextStyle(
                    fontSize: 24,
                    fontWeight: FontWeight.bold,
                  ),
                ),
                smallSpacer(),
                Container(
                  height: 2,
                  width: 100,
                  color: AppColors.primaryGreen,
                ),
                largeSpacer(),
                CustomTextField(
                  controller: widget.emailController,
                  label: 'Adresse email',
                  hint: 'Entrer mon adresse email',
                  keyboardType: TextInputType.emailAddress,
                ),
                smallerSpacer(),
                CustomTextField(
                  controller: widget.firstnameController,
                  label: 'Prénom',
                  hint: 'Entrer mon prénom',
                ),
                smallerSpacer(),
                CustomTextField(
                  controller: widget.lastnameController,
                  label: 'Nom de famille',
                  hint: 'Entrer mon nom de famille',
                ),
                smallerSpacer(),
                TextField(
                  controller: widget.passwordController,
                  obscureText: !_obscurePassword,
                  decoration: InputDecoration(
                    labelText: 'Mot de passe',
                    floatingLabelStyle:
                        TextStyle(color: Colors.green), // Label focus
                    hintText: 'Entrer mon mot de passe',
                    suffixIcon: IconButton(
                      icon: Icon(
                        _obscurePassword
                            ? Icons.visibility
                            : Icons.visibility_off,
                        color: Colors.green, // Optionnel : couleur de l’icône
                      ),
                      onPressed: () {
                        setState(() {
                          _obscurePassword = !_obscurePassword;
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
                ),
                const SizedBox(height: 60),
                // Error Message
                if (widget.errorMessage.isNotEmpty)
                  Padding(
                    padding: const EdgeInsets.only(bottom: 10),
                    child: Text(
                      widget.errorMessage,
                      style: const TextStyle(color: AppColors.primaryRed),
                    ),
                  ),
                widget.isLoading
                    ? const CircularProgressIndicator()
                    : SizedBox(
                        width: double.infinity,
                        child: customButton(
                          text: 'S’enregistrer',
                          textColor: AppColors.primaryWhite,
                          backgroundColor: AppColors.primaryGreen,
                          onPressed: widget.onRegister,
                        ),
                      ),

                const SizedBox(height: 10),
                SizedBox(
                  width: double.infinity,
                  child: customButton(
                    text: 'Se connecter',
                    onPressed: widget.onNavigateToLogin,
                    backgroundColor: AppColors.primaryWhite,
                    textColor: AppColors.primaryGreen,
                    borderColor: AppColors.primaryGreen,
                    borderWidth: 3,
                  ),
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
