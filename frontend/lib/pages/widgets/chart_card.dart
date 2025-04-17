import 'package:Arly/core/style.dart';
import 'package:flutter/material.dart';

class ChartCard extends StatelessWidget {
  final Widget child;
  final String? title;
  final IconData? icon;
  final Color color;
  final Color textColor;

  const ChartCard({
    super.key,
    required this.child,
    this.title,
    this.icon,
    this.color = AppColors.primaryWhite,
    this.textColor = Colors.black87,
  });

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.symmetric(horizontal: 16.0, vertical: 10),
      child: Container(
        decoration: BoxDecoration(
          color: AppColors.primaryWhite,
          borderRadius: BorderRadius.circular(24),
        ),
        padding: const EdgeInsets.symmetric(horizontal: 20, vertical: 24),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            if (title != null)
              Row(
                children: [
                  if (icon != null)
                    Icon(icon, color: textColor, size: 20),
                  if (icon != null) const SizedBox(width: 8),
                  Text(
                    title!,
                    style: Theme.of(context).textTheme.titleMedium!.copyWith(
                      fontWeight: FontWeight.w600,
                      color: textColor,
                    ),
                  ),
                ],
              ),
            if (title != null) const SizedBox(height: 16),
            child,
          ],
        ),
      ),
    );
  }
}
