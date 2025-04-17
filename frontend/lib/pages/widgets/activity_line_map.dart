import 'package:Arly/core/style.dart';
import 'package:fl_chart/fl_chart.dart';
import 'package:flutter/material.dart';
import 'package:intl/intl.dart';
import 'package:intl/date_symbol_data_local.dart';

class ActivityLineChart extends StatefulWidget {
  final List<dynamic> userActivities;

  const ActivityLineChart({super.key, required this.userActivities});

  @override
  State<ActivityLineChart> createState() => _ActivityLineChartState();
}

class _ActivityLineChartState extends State<ActivityLineChart> {
  bool _localeReady = false;

  @override
  @override
  void initState() {
    super.initState();
    _initializeLocale();
  }

  Future<void> _initializeLocale() async {
    await initializeDateFormatting('fr_FR', null);
    setState(() {
      _localeReady = true;
    });
  }

  @override
  Widget build(BuildContext context) {
    if (!_localeReady || widget.userActivities.isEmpty) {
      return const Center(child: CircularProgressIndicator());
    }

    final sortedActivities = List.from(widget.userActivities)
      ..sort((a, b) => DateTime.parse(a['date']).compareTo(DateTime.parse(b['date'])));

    final spots = <FlSpot>[];
    final xLabels = <String>[];
    double maxY = 0;

    for (int i = 0; i < sortedActivities.length; i++) {
      final activity = sortedActivities[i];
      final date = DateTime.tryParse(activity['date'].toString());
      final count = double.tryParse(activity['message_count'].toString()) ?? 0.0;

      if (date != null) {
        spots.add(FlSpot(i.toDouble(), count));
        final dayLetter = DateFormat('E', 'fr_FR').format(date).substring(0, 1).toUpperCase();
        xLabels.add(dayLetter);
        if (count > maxY) maxY = count;
      }
    }

    maxY = (maxY / 20).ceil() * 20;

    return Column(
      crossAxisAlignment: CrossAxisAlignment.start,
      children: [
        // Title
        Padding(
          padding: const EdgeInsets.symmetric(horizontal: 0, vertical: 0),
          child: Row(
            children: [
              const Icon(Icons.show_chart, color: AppColors.tealPrimary, size: 30),
              const SizedBox(width: 12),
              Text(
                "Carte d’activités",
                style: Theme.of(context).textTheme.titleLarge!.copyWith(
                  fontSize: 17,
                  fontWeight: FontWeight.w700,
                  color: AppColors.primaryGreen,
                ),
              ),
            ],
          ),
        ),
        const SizedBox(height: 24),

        // Line Chart
        SizedBox(
          height: 200,
          child: Padding(
            padding: const EdgeInsets.symmetric(horizontal: 0),
            child: LineChart(
              LineChartData(
                minY: 0,
                maxY: maxY,
                lineTouchData: LineTouchData(
                  handleBuiltInTouches: true,
                  touchTooltipData: LineTouchTooltipData(
                    tooltipBgColor: AppColors.primaryWhite,
                    tooltipRoundedRadius: 12,
                    tooltipMargin: 8,
                    getTooltipItems: (spots) => spots.map((spot) {
                      return LineTooltipItem(
                        '${spot.y.toInt()} messages',
                        const TextStyle(
                          fontWeight: FontWeight.bold,
                          color: Colors.black87,
                        ),
                      );
                    }).toList(),
                  ),
                ),
                gridData: FlGridData(
                  show: true,
                  drawVerticalLine: false,
                  horizontalInterval: 20,
                  getDrawingHorizontalLine: (value) => FlLine(
                    color: AppColors.gey01.withOpacity(0.1),
                    strokeWidth: 1,
                  ),
                ),
                borderData: FlBorderData(show: false),
                titlesData: FlTitlesData(
                  leftTitles: AxisTitles(
                    sideTitles: SideTitles(
                      showTitles: true,
                      reservedSize: 40,
                      interval: 20,
                      getTitlesWidget: (value, _) => Text(
                        value.toInt().toString(),
                        style: const TextStyle(fontSize: 12),
                      ),
                    ),
                  ),
                  bottomTitles: AxisTitles(
                    sideTitles: SideTitles(
                      showTitles: true,
                      getTitlesWidget: (value, _) {
                        final index = value.toInt();
                        if (index >= 0 && index < xLabels.length) {
                          return Transform.translate(
                            offset: const Offset(0, 6), // 👈 Push label down
                            child: Text(
                              xLabels[index],
                              style: const TextStyle(fontSize: 12),
                            ),
                          );
                        }
                        return const SizedBox.shrink();
                      },
                    ),
                  ),
                  topTitles: AxisTitles(sideTitles: SideTitles(showTitles: false)),
                  rightTitles: AxisTitles(sideTitles: SideTitles(showTitles: false)),
                ),
                lineBarsData: [
                  LineChartBarData(
                    spots: spots,
                    isCurved: true,
                    color: AppColors.primaryGreen,
                    barWidth: 4,
                    isStrokeCapRound: true,
                    dotData: FlDotData(
                      show: true,
                      getDotPainter: (spot, _, __, ___) => FlDotCirclePainter(
                        radius: 5,
                        color: AppColors.primaryGreen,
                        strokeColor: AppColors.primaryWhite,
                        strokeWidth: 2,
                      ),
                    ),
                    belowBarData: BarAreaData(
                      show: true,
                      color: const Color(0xFFCBECCF).withOpacity(0.2),
                    ),
                  ),
                ],
              ),
            ),
          ),
        ),
        const SizedBox(height: 24),
      ],
    );
  }
}
