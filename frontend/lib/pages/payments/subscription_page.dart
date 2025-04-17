import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:http/http.dart' as http;
import 'package:Arly/pages/Login/login_page.dart';
import 'package:Arly/pages/Profil/seeting_page.dart';
import 'package:Arly/pages/home/home_page.dart';
import 'package:Arly/pages/Chat/chat_page.dart';
import 'package:Arly/pages/widgets/webview_screen.dart'; // Import the WebView page
import 'package:Arly/config.dart';


class SubscriptionPage extends StatefulWidget {
  const SubscriptionPage({Key? key}) : super(key: key);

  @override
  _SubscriptionPageState createState() => _SubscriptionPageState();
}

class _SubscriptionPageState extends State<SubscriptionPage> {
  final int _currentIndex = 2;
  List<dynamic> _subscriptionTypes = [];
  Map<String, dynamic>? _currentSubscription;

  @override
  void initState() {
    super.initState();
    _fetchSubscriptions();
  }

  Future<void> _fetchSubscriptions() async {
    final prefs = await SharedPreferences.getInstance();
    final token = prefs.getString('authToken');
    if (token == null) {
      _navigateToLogin();
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
        final data = jsonDecode(response.body);
        setState(() {
          _currentSubscription = data['subscription_type'];
          _subscriptionTypes = data['subscription_types']; //id, name,price
        });
      } else {
        print('Failed to fetch subscriptions: ${response.body}');
      }
    } catch (e) {
      print('Error fetching subscriptions: $e');
    }
  }

  void _navigateToLogin() {
    Navigator.of(context).pushReplacement(
      MaterialPageRoute(builder: (context) => const LoginPage()),
    );
  }

  Future<void> _upgradeSubscription(Map<String, dynamic> subscription) async {
    final prefs = await SharedPreferences.getInstance();
    final token = prefs.getString('authToken');
    if (token == null) {
      _navigateToLogin();
      return;
    }

    try {
      final response = await http.post(
        Uri.parse('${HOST}/api/payment/subscription'),
        headers: {
          'Authorization': 'Bearer $token',
          'Content-Type': 'application/json',
        },
        body: jsonEncode({'price_id': subscription['price_id']}),
      );

      if (response.statusCode == 200) {
        final data = jsonDecode(response.body);
        final checkoutSessionUrl = data['checkout_session_url'];

        if (checkoutSessionUrl != null && checkoutSessionUrl is String) {
          Navigator.of(context).push(
            MaterialPageRoute(
              builder: (context) => WebViewScreen(url: checkoutSessionUrl),
            ),
          );
        } else {
          ScaffoldMessenger.of(context).showSnackBar(
            const SnackBar(content: Text('Invalid payment session.')),
          );
        }
      } else {
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(content: Text('Failed to upgrade subscription.')),
        );
      }
    } catch (e) {
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Error upgrading subscription.')),
      );
    }
  }

   Widget _buildSubscriptionCard(Map<String, dynamic> subscription) {
    final subscriptionName = subscription['subscription_name'];
    final subscriptionTypeId = subscription['id'].toString();
    final subscriptionPrice = subscription['price'];

    bool isCurrent = _currentSubscription != null &&
        (_currentSubscription!['id'].toString() == subscriptionTypeId);

    bool isDowngrade = _currentSubscription != null &&
        (_currentSubscription!['price'] > subscriptionPrice);

    return Card(
      margin: const EdgeInsets.symmetric(vertical: 8.0),
      elevation: 3,
      shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12.0)),
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: Row(
          children: [
            Expanded(
              child: Column(
                crossAxisAlignment: CrossAxisAlignment.start,
                children: [
                  Text(
                    subscriptionName,
                    style: const TextStyle(fontSize: 18, fontWeight: FontWeight.bold),
                  ),
                  const SizedBox(height: 8),
                  Text("Subscription Plan", style: const TextStyle(fontSize: 14)),
                  Text("Price: \$${subscriptionPrice.toString()}",
                      style: const TextStyle(fontSize: 14, fontWeight: FontWeight.w500)),
                ],
              ),
            ),
            ElevatedButton(
              onPressed: (isCurrent || isDowngrade) ? null : () => _upgradeSubscription(subscription),
              child: Text(isCurrent ? 'Current Plan' : (isDowngrade ? 'Downgrade Not Allowed' : 'Upgrade')),
              style: ElevatedButton.styleFrom(
                shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(8.0)),
              ),
            ),
          ],
        ),
      ),
    );
  }


  void _onBottomNavTap(int index) {
    if (index == _currentIndex) return;
    Navigator.of(context).pushReplacement(
      MaterialPageRoute(
        builder: (context) {
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
        },
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(title: const Text('Subscriptions'), centerTitle: true),
      body: Padding(
        padding: const EdgeInsets.all(16.0),
        child: _subscriptionTypes.isEmpty
            ? const Center(child: CircularProgressIndicator())
            : ListView.builder(
                itemCount: _subscriptionTypes.length,
                itemBuilder: (context, index) {
                  final subscription = _subscriptionTypes[index];
                  return _buildSubscriptionCard(subscription);
                },
              ),
      ),
      bottomNavigationBar: BottomNavigationBar(
        currentIndex: _currentIndex,
        onTap: _onBottomNavTap,
        items: const [
          BottomNavigationBarItem(icon: Icon(Icons.home), label: 'Home'),
          BottomNavigationBarItem(icon: Icon(Icons.chat), label: 'Chat'),
          BottomNavigationBarItem(icon: Icon(Icons.settings), label: 'Settings'),
        ],
      ),
    );
  }
}
