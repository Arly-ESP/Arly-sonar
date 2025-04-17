import 'package:Arly/core/style.dart';
import 'package:flutter/material.dart';

class MessageInput extends StatelessWidget {
  final TextEditingController controller;
  final bool isListening;
  final bool isLoading;
  final VoidCallback onMicTap;
  final VoidCallback onSendTap;

  const MessageInput({
    super.key,
    required this.controller,
    required this.isListening,
    required this.isLoading,
    required this.onMicTap,
    required this.onSendTap,
  });

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(vertical: 8.0, horizontal: 8.0),
      child: Row(
        children: [
          Expanded(
            child: TextField(
              controller: controller,
              maxLines: null,
              decoration: InputDecoration(
                hintText: 'Tapez votre message...',
                filled: true,
                fillColor: AppColors.primaryWhite,
                border: OutlineInputBorder(
                  borderRadius: BorderRadius.circular(20),
                  borderSide: BorderSide.none,
                ),
                contentPadding: const EdgeInsets.symmetric(
                    vertical: 10.0, horizontal: 15.0),
              ),
            ),
          ),
          IconButton(
            icon: Icon(
              isListening ? Icons.mic : Icons.mic_none,
              color: isListening ? AppColors.primaryRed : AppColors.primaryGreen,
            ),
            onPressed: onMicTap,
          ),
          IconButton(
            icon: const Icon(Icons.send, color: AppColors.primaryGreen),
            onPressed: isLoading ? null : onSendTap,
          ),
        ],
      ),
    );
  }
}
