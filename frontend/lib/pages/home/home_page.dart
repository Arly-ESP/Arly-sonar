import 'dart:convert';
import 'package:Arly/core/style.dart';
import 'package:Arly/pages/Login/login_page.dart';
import 'package:Arly/pages/Chat/chat_page.dart';
import 'package:Arly/pages/Profil/seeting_page.dart';
import 'package:Arly/pages/widgets/activity_line_map.dart';
import 'package:features_tour/features_tour.dart';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';

import '../widgets/home_app_bar.dart';
import '../widgets/mood_selection_row.dart';
import '../widgets/mood_chart.dart';
import '../widgets/mood_summary.dart';
import 'package:Arly/config.dart';
import 'package:Arly/pages/widgets/chart_card.dart';

class HomePage extends StatefulWidget {
  const HomePage({super.key});
  @override
  _HomePageState createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  final int _currentIndex = 0;
  int _expandedIndex = -1;
  final String _selectedFilter = 'Month';
  List<dynamic> _userActivities = [];
  final tourController = FeaturesTourController('HomePage');

  String _getMood(int index) {
    switch (index) {
      case 0:
        return 'bad';
      case 1:
        return 'sad';
      case 2:
        return 'poker';
      case 3:
        return 'nice';
      case 4:
        return 'happy';
      default:
        return 'neutral';
    }
  }

  final List<Widget> _pages = [
    const HomePage(),
    const ChatPage(),
    SettingsPage(),
  ];

  @override
  void initState() {
    super.initState();
    getUser(context);
    tourController.start(context);
  }

  Future<void> sendMoodRequest(String mood) async {
    try {
      final prefs = await SharedPreferences.getInstance();
      final token = prefs.getString('authToken');

      if (token == null) {
        _navigateToLogin(context);
        return;
      }

      final response = await http.post(
        Uri.parse('$HOST/api/mood'),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
        body: jsonEncode({"mood": mood}),
      );

      if (response.statusCode == 200 || response.statusCode == 201) {
        ScaffoldMessenger.of(context).showSnackBar(
          SnackBar(
            content: Text('Mood updated to "$mood" successfully!'),
            backgroundColor: AppColors.primaryGreen,
            duration: const Duration(seconds: 2),
          ),
        );
        fetchUserActivities();
      } else {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(
            content: Text('Failed to update mood. Please try again.'),
            backgroundColor: AppColors.primaryRed,
            duration: Duration(seconds: 2),
          ),
        );
      }
    } catch (_) {
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(
          content: Text('An error occurred while updating mood.'),
          backgroundColor: AppColors.primaryRed,
          duration: Duration(seconds: 2),
        ),
      );
    }
  }

  Future<void> fetchUserActivities() async {
    try {
      final prefs = await SharedPreferences.getInstance();
      final token = prefs.getString('authToken');

      if (token == null) {
        _navigateToLogin(context);
        return;
      }

      final response = await http.get(
        Uri.parse('$HOST/api/user/activities'),
        headers: {
          'accept': 'application/json',
          'Authorization': 'Bearer $token',
        },
      );
      if (response.statusCode == 200) {
        setState(() {
          _userActivities = jsonDecode(response.body);
        });
      }
    } catch (_) {}
  }

  Future<void> getUser(BuildContext context) async {
    try {
      fetchUserActivities();
      final prefs = await SharedPreferences.getInstance();
      final token = prefs.getString('authToken');

      if (token == null) {
        _navigateToLogin(context);
        return;
      }

      final response = await http.get(
        Uri.parse('$HOST/api/user'),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
      );

      if (response.statusCode == 401) {
        _navigateToLogin(context);
      }
    } catch (_) {
      _navigateToLogin(context);
    }
  }

  void _navigateToLogin(BuildContext context) {
    Navigator.of(context).pushAndRemoveUntil(
      MaterialPageRoute(builder: (context) => const LoginPage()),
          (Route<dynamic> route) => false,
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: const HomeAppBar(),
      body: Container(
        decoration: const BoxDecoration(
          gradient: LinearGradient(
            begin: Alignment.topLeft,
            end: Alignment.bottomRight,
            colors: [
              Color(0xFFF4FBF4),
              Color(0xFFFFF9F4),
            ],
          ),
        ),
        child: SingleChildScrollView(
          physics: const BouncingScrollPhysics(),
          padding: const EdgeInsets.symmetric(vertical: 24.0),
          child: Column(
            children: [
              Padding(
                padding: const EdgeInsets.symmetric(horizontal: 16),
                child: FeaturesTour(
                  controller: tourController,
                  index: 1,
                  introduce: const Text(
                    'Tu peux choisir ton humeur du jour ici en tapant sur l\'une de ces icônes !',
                    style: TextStyle(fontSize: 24, color: AppColors.primaryWhite),
                  ),
                  child: Container(
                    padding: const EdgeInsets.symmetric(vertical: 16, horizontal: 12),
                    decoration: BoxDecoration(
                      color: Colors.white.withOpacity(0.9),
                      borderRadius: BorderRadius.circular(24),
                      boxShadow: [
                        BoxShadow(
                          color: Colors.black.withOpacity(0.05),
                          blurRadius: 12,
                          offset: const Offset(0, 4),
                        ),
                      ],
                    ),
                    child: MoodSelectionRow(
                      expandedIndex: _expandedIndex,
                      onMoodTap: (index) {
                        setState(() {
                          sendMoodRequest(_getMood(index));
                          _expandedIndex = _expandedIndex == index ? -1 : index;
                        });
                      },
                    ),
                  ),
                ),
              ),
              const SizedBox(height: 24),
              Padding(
                padding: const EdgeInsets.symmetric(horizontal: 2),
                child: FeaturesTour(
                  controller: tourController,
                  index: 2,
                  introduce: const Text(
                    'Cette zone permet de visualiser tes humeurs au fil du temps...',
                    style: TextStyle(fontSize: 24, color: AppColors.primaryWhite),
                  ),
                  child: ChartCard(
                    icon: Icons.bar_chart,
                    color: AppColors.superLightSage,
                    child: MoodChart(userActivities: _userActivities),
                  ),
                ),
              ),
              const SizedBox(height: 24),
              Padding(
                padding: const EdgeInsets.symmetric(horizontal: 16),
                child: FeaturesTour(
                  controller: tourController,
                  index: 3,
                  introduce: const Text(
                    'Cette zone permet de visualiser tes activités passées.',
                    style: TextStyle(fontSize: 24, color: AppColors.primaryWhite),
                  ),
                  child: Container(
                    decoration: BoxDecoration(
                      color: AppColors.primaryWhite,
                      borderRadius: BorderRadius.circular(24),
                      boxShadow: [
                        BoxShadow(
                          color: Colors.black.withOpacity(0.05),
                          blurRadius: 12,
                          offset: const Offset(0, 4),
                        ),
                      ],
                    ),
                    child: Padding(
                      padding: const EdgeInsets.all(20),
                      child: ChartCard(
                        icon: Icons.show_chart,
                        color: Colors.transparent,
                        child: ActivityLineChart(userActivities: _userActivities, ),
                      ),
                    ),
                  ),
                ),
              ),
              const SizedBox(height: 24),
              Padding(
                padding: const EdgeInsets.symmetric(horizontal: 16),
                child: FeaturesTour(
                  controller: tourController,
                  index: 4,
                  introduce: const Text(
                    'Cette zone résume tes humeurs de la semaine.',
                    style: TextStyle(fontSize: 24, color: AppColors.primaryWhite),
                  ),
                  child: Container(
                    width: double.infinity,
                    decoration: BoxDecoration(
                      color: Colors.white.withOpacity(0.95),
                      borderRadius: BorderRadius.circular(24),
                      boxShadow: [
                        BoxShadow(
                          color: Colors.black.withOpacity(0.05),
                          blurRadius: 12,
                          offset: const Offset(0, 4),
                        ),
                      ],
                    ),
                    child: Padding(
                      padding: const EdgeInsets.all(20),
                      child: MoodSummary(userActivities: _userActivities),
                    ),
                  ),
                ),
              ),
              const SizedBox(height: 32),
            ],
          ),
        ),
      ),
      bottomNavigationBar: FeaturesTour(
        controller: tourController,
        index: 5,
        introduce: const Text(
          'Cette zone permet de naviguer entre les différentes sections de l\'application.',
          style: TextStyle(fontSize: 24, color: AppColors.primaryWhite),
        ),
        child: Container(
          decoration: BoxDecoration(
            color: Colors.white,
            borderRadius: const BorderRadius.only(
              topLeft: Radius.circular(24),
              topRight: Radius.circular(24),
            ),
            boxShadow: [
              BoxShadow(
                color: Colors.black.withOpacity(0.08),
                blurRadius: 10,
                offset: const Offset(0, -2),
              ),
            ],
          ),
          child: ClipRRect(
            borderRadius: const BorderRadius.only(
              topLeft: Radius.circular(24),
              topRight: Radius.circular(24),
            ),
            child: BottomNavigationBar(
              type: BottomNavigationBarType.fixed,
              backgroundColor: Colors.white,
              selectedItemColor: AppColors.primaryGreen,
              unselectedItemColor: Colors.grey[400],
              selectedLabelStyle: const TextStyle(
                fontWeight: FontWeight.bold,
                fontSize: 13,
              ),
              unselectedLabelStyle: const TextStyle(
                fontWeight: FontWeight.w500,
                fontSize: 12,
              ),
              currentIndex: _currentIndex,
              onTap: (index) async {
                if (_currentIndex != index) {
                  await Navigator.of(context).push(
                    MaterialPageRoute(builder: (context) => _pages[index]),
                  );
                  fetchUserActivities();
                }
              },
              items: const [
                BottomNavigationBarItem(
                  icon: Icon(Icons.home_outlined),
                  label: 'Accueil',
                ),
                BottomNavigationBarItem(
                  icon: Icon(Icons.chat_bubble_outline),
                  label: 'Chat',
                ),
                BottomNavigationBarItem(
                  icon: Icon(Icons.settings_outlined),
                  label: 'Paramètres',
                ),
              ],
            ),
          ),
        ),
      ),
    );
  }
}
