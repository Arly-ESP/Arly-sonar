import 'dart:convert';
import 'package:Arly/pages/Chat/chat_page.dart';
import 'package:Arly/pages/Login/login_page.dart';
import 'package:Arly/pages/detail/questionnaire.dart';
import 'package:Arly/pages/home/home_page.dart';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';
import '../../config.dart';
import '../widgets/seeting_app_bar.dart';
import '../widgets/user_seetings_card.dart';
import 'package:Arly/core/style.dart';

class SettingsPage extends StatefulWidget {
  const SettingsPage({super.key});

  @override
  _SettingsPageState createState() => _SettingsPageState();
}

class _SettingsPageState extends State<SettingsPage> {
  final int _currentIndex = 2;
  String _email = '';
  String _name = '';
  String _avatar = '';
  String _subscriptionType = '';
  String _subscriptionId = '';
  bool _isRecurringSubscription = false;

  Future<void> deleteAccount(BuildContext context) async {
    final prefs = await SharedPreferences.getInstance();
    final token = prefs.getString('authToken');

    if (token == null) {
      _navigateToLogin(context);
      return;
    }

    try {
      final response = await http.delete(
        Uri.parse('${HOST}/api/user'),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-Type': 'application/json',
        },
      );

      if (response.statusCode == 200 || response.statusCode == 204) {
        await prefs.remove('authToken');
        _navigateToLogin(context);
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(
            content: Text("Votre compte a été supprimé avec succès."),
            backgroundColor: AppColors.primaryGreen,
            duration: Duration(seconds: 2),
          ),
        );
      } else {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(
            content: Text("Échec de la suppression du compte."),
            backgroundColor: AppColors.primaryRed,
            duration: Duration(seconds: 2),
          ),
        );
      }
    } catch (e) {
      print('Error: $e');
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(
          content: Text("Une erreur est survenue. Veuillez réessayer."),
          backgroundColor: AppColors.primaryRed,
          duration: Duration(seconds: 2),
        ),
      );
    }
  }


  Future<void> changeAvatar(BuildContext context) async {
    final prefs = await SharedPreferences.getInstance();

    // Par exemple, tu proposes plusieurs avatars prédéfinis (chemins locaux)
    List<String> avatars = [
      'assets/baguette.png',
      'assets/chat.png',
      'assets/sushi.png',
      'assets/tagin.png',
    ];

    // Tu peux afficher un dialog pour que l'utilisateur choisisse un avatar
    String? selectedAvatar = await showDialog<String>(
      context: context,
      builder: (context) => SimpleDialog(
        title: const Text("Choisissez un avatar"),
        children: avatars.map((path) {
          return SimpleDialogOption(
            onPressed: () {
              Navigator.pop(context, path);
            },
            child: Row(
              children: [
                Image.asset(path, height: 40, width: 40),
                const SizedBox(width: 10),
                Text(path.split('/').last),
              ],
            ),
          );
        }).toList(),
      ),
    );

    if (selectedAvatar != null) {
      await prefs.setString('pandAvatar', selectedAvatar);
      setState(() {
        _avatar = selectedAvatar;
      });
    }
  }

  void _navigateToSubscription(BuildContext context) {
    Navigator.of(context).pushNamed('/subscription');
  }

  void _navigateToLogin(BuildContext context) {
    Navigator.of(context).pushAndRemoveUntil(
      MaterialPageRoute(builder: (context) => LoginPage()),
          (Route<dynamic> route) => false,
    );
  }

  Future<void> takeQuestionnaryBack(BuildContext context) async {
    Navigator.of(context).pushAndRemoveUntil(
      MaterialPageRoute(builder: (context) => QuestionnairePage()),
          (Route<dynamic> route) => true,
    );
  }

  Future<void> logoutAndRedirect(BuildContext context) async {
    final prefs = await SharedPreferences.getInstance();
    await prefs.remove('authToken');
    Navigator.of(context).pushAndRemoveUntil(
      MaterialPageRoute(builder: (context) => LoginPage()),
          (Route<dynamic> route) => false,
    );
  }

  Future<void> getUserSubscriptions() async {
    final prefs = await SharedPreferences.getInstance();
    final token = prefs.getString('authToken');

    if (token == null) {
      _navigateToLogin(context);
      return;
    }

    try {
      final response = await http.get(
        Uri.parse('${HOST}/api/payment/subscription'),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-Type': 'application/json',
        },
      );

      if (response.statusCode == 200) {
        final subscriptions = jsonDecode(response.body);

        setState(() {
          _subscriptionType =
          subscriptions['subscription_type']['subscription_name'];
          _subscriptionId = subscriptions['subscription']['id'].toString();
          _isRecurringSubscription =
          subscriptions['subscription_type']['is_recurring'];
        });
        print('Subscriptions: $subscriptions');
      } else {
        print('Error: ${response.body}');
      }
    } catch (e) {
      print('Error: $e');
    }
  }

  Future<void> getUser(BuildContext context) async {
    try {
      final prefs = await SharedPreferences.getInstance();
      final token = prefs.getString('authToken');

      if (token == null) {
        _navigateToLogin(context);
        return;
      }

      final response = await http.get(
        Uri.parse('${HOST}/api/user'),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
      );

      if (response.statusCode == 200) {
        final userData = jsonDecode(response.body);
        setState(() {
          _email = userData['email'];
          _name = userData['first_name'];
        });
        print('User is connected');
        await getUserSubscriptions();
      } else {
        _navigateToLogin(context);
      }
    } catch (e) {
      print('Error: $e');
      _navigateToLogin(context);
    }
  }

  @override
  void initState() {
    super.initState();
    getUser(context);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: const SettingsAppBar(),

      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: UserSettingsCard(
          name: _name,
          email: _email,
          avatar: _avatar,
          subscriptionType: _subscriptionType,
          subscriptionId: _subscriptionId,
          isRecurringSubscription: _isRecurringSubscription,
          onTakeQuestionnaire: () => takeQuestionnaryBack(context),
          onContact: () {},
          onDeleteAccount: () async => await deleteAccount(context),
          onLogout: () => logoutAndRedirect(context),
          navigateToSubscription: () => _navigateToSubscription(context),
        ),
      ),
      bottomNavigationBar: BottomNavigationBar(
        currentIndex: _currentIndex,
        onTap: (index) {
          if (_currentIndex != index) {
            Navigator.of(context).pushReplacement(
              MaterialPageRoute(builder: (context) {
                switch (index) {
                  case 0:
                    return const HomePage();
                  case 1:
                    return const ChatPage();
                  case 2:
                    return const SettingsPage();
                  default:
                    return const HomePage();
                }
              }),
            );
          }
        },
        items: const [
          BottomNavigationBarItem(
            icon: Icon(Icons.home),
            label: 'Accueil',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.chat),
            label: 'Chat',
          ),
          BottomNavigationBarItem(
            icon: Icon(Icons.settings),
            label: 'Paramètres',
          ),
        ],
      ),
    );
  }
}
