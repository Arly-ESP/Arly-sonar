import 'package:flutter/material.dart';

class VerificationBackground extends StatelessWidget {
  const VerificationBackground({super.key});

  @override
  Widget build(BuildContext context) {
    return Positioned.fill(
      child: Image.asset(
        'assets/panda.jpg',
        fit: BoxFit.cover,
      ),
    );
  }
}
