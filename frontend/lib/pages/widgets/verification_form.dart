import 'package:Arly/core/style.dart';
import 'package:flutter/material.dart';

class VerificationForm extends StatelessWidget {
  final String email;
  final List<TextEditingController> controllers;
  final String errorMessage;
  final bool isLoading;
  final bool isResending;
  final VoidCallback onVerify;
  final VoidCallback onResend;

  const VerificationForm({
    super.key,
    required this.email,
    required this.controllers,
    required this.errorMessage,
    required this.isLoading,
    required this.isResending,
    required this.onVerify,
    required this.onResend,
  });

  @override
  Widget build(BuildContext context) {
    return Align(
      alignment: Alignment.bottomCenter,
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Container(
          padding:
              const EdgeInsets.symmetric(vertical: 80, horizontal: 25),
          decoration: BoxDecoration(
            color: AppColors.primaryWhite,
            borderRadius: BorderRadius.circular(25),
          ),
          child: Column(
            mainAxisSize: MainAxisSize.min,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              const Text(
                'Validation',
                style: TextStyle(
                  fontSize: 24,
                  fontWeight: FontWeight.bold,
                ),
              ),
              const SizedBox(height: 5),
              Container(
                height: 2,
                width: 100,
                color: AppColors.primaryGreen,
              ),
              const SizedBox(height: 15),
              Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  const Text(
                    'Adresse email',
                    style: TextStyle(fontWeight: FontWeight.bold),
                  ),
                  Text(email),
                  TextButton(
                    onPressed: isResending ? null : onResend,
                    child: isResending
                        ? const CircularProgressIndicator()
                        : const Text(
                            'Renvoyer',
                            style: TextStyle(color: AppColors.primaryGreen),
                          ),
                  ),
                ],
              ),
              const SizedBox(height: 15),
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: List.generate(6, (index) {
                  return SizedBox(
                    width: 40,
                    child: TextField(
                      controller: controllers[index],
                      keyboardType: TextInputType.number,
                      maxLength: 1,
                      textAlign: TextAlign.center,
                      decoration: InputDecoration(
                        counterText: '',
                        border: OutlineInputBorder(
                          borderRadius: BorderRadius.circular(10),
                        ),
                      ),
                      onChanged: (value) {
                        if (value.length == 1 && index < 5) {
                          FocusScope.of(context).nextFocus();
                        } else if (value.isEmpty && index > 0) {
                          FocusScope.of(context).previousFocus();
                        }
                      },
                    ),
                  );
                }),
              ),
              const SizedBox(height: 15),
              if (errorMessage.isNotEmpty)
                Text(
                  errorMessage,
                  style: const TextStyle(color: AppColors.primaryRed),
                ),
              const SizedBox(height: 20),
              SizedBox(
                width: double.infinity,
                child: ElevatedButton(
                  onPressed: isLoading ? null : onVerify,
                  style: ElevatedButton.styleFrom(
                    padding: const EdgeInsets.all(15),
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(10),
                    ),
                    backgroundColor: AppColors.primaryGreen,
                  ),
                  child: isLoading
                      ? const CircularProgressIndicator(
                          color: AppColors.primaryWhite,
                        )
                      : const Text(
                          'Valider',
                          style: TextStyle(fontSize: 16),
                        ),
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
