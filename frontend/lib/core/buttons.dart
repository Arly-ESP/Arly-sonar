import 'package:Arly/core/style.dart';
import 'package:flutter/material.dart';

ElevatedButton loginButton({
  required String text,
  required VoidCallback onPressed,
  Color backgroundColor = AppColors.primaryGreen,
  Color textColor = AppColors.primaryWhite,
  Color borderColor = Colors.transparent,
  double borderWidth = 0,
}) {
  return ElevatedButton(
    onPressed: onPressed,
    style: ElevatedButton.styleFrom(
      padding: const EdgeInsets.all(15),
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(10),
        side: BorderSide(
          color: borderColor,
          width: borderWidth,
        ),
      ),
      backgroundColor: backgroundColor,
    ),
    child: Text(
      text,
      style: TextStyle(fontSize: 16, color: textColor),
    ),
  );
}

ElevatedButton RegisterButton({
  required String text,
  required VoidCallback onPressed,
  Color backgroundColor = AppColors.primaryWhite,
  Color textColor = AppColors.primaryGreen,
  Color borderColor = Colors.green,
  double borderWidth = 3,
}) {
  return ElevatedButton(
    onPressed: onPressed,
    style: ElevatedButton.styleFrom(
      padding: const EdgeInsets.all(15),
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(10),
        side: BorderSide(
          color: borderColor,
          width: borderWidth,
        ),
      ),
      backgroundColor: backgroundColor,
    ),
    child: Text(
      text,
      style: TextStyle(fontSize: 16, color: textColor),
    ),
  );
}

ElevatedButton customButton({
  required String text,
  required VoidCallback onPressed,
  Color backgroundColor = AppColors.primaryGreen,
  Color textColor = AppColors.primaryWhite,
  Color borderColor = Colors.transparent,
  double borderWidth = 0,
  double borderRadius = 10, // Added parameter for border radius
}) {
  return ElevatedButton(
    onPressed: onPressed,
    style: ElevatedButton.styleFrom(
      padding: const EdgeInsets.all(15),
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(borderRadius),
        side: BorderSide(
          color: borderColor,
          width: borderWidth,
        ),
      ),
      backgroundColor: backgroundColor,
    ),
    child: Text(
      text,
      style: TextStyle(fontSize: 16, color: textColor),
    ),
  );
}
