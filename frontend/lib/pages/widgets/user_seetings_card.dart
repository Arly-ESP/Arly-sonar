import 'package:Arly/core/style.dart';
import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';
import '../../avatar_notifier.dart';


class UserSettingsCard extends StatefulWidget {
  final String name;
  final String email;
  final String avatar;
  final String subscriptionType;
  final String subscriptionId;
  final bool isRecurringSubscription;
  final VoidCallback onTakeQuestionnaire;
  final VoidCallback onContact;
  final VoidCallback onDeleteAccount;
  final VoidCallback onLogout;
  final VoidCallback navigateToSubscription;

  const UserSettingsCard({
    super.key,
    required this.name,
    required this.email,
    required this.avatar,
    required this.subscriptionType,
    required this.subscriptionId,
    required this.isRecurringSubscription,
    required this.onTakeQuestionnaire,
    required this.onContact,
    required this.onDeleteAccount,
    required this.onLogout,
    required this.navigateToSubscription,
  });

  @override
  _UserSettingsCardState createState() => _UserSettingsCardState();
}

class _UserSettingsCardState extends State<UserSettingsCard> {
  int? selectedAvatarIndex;

  @override
  void initState() {
    super.initState();
    _loadAvatarSelection();
  }

  Future<void> _loadAvatarSelection() async {
    final prefs = await SharedPreferences.getInstance();
    setState(() {
      selectedAvatarIndex = prefs.getInt('selectedAvatarIndex') ?? 0;
    });
  }

  Future<void> _saveAvatarSelection(int index) async {
    final prefs = await SharedPreferences.getInstance();
    await prefs.setInt('selectedAvatarIndex', index);
  }

  void _onAvatarTap(int index) {
    setState(() {
      selectedAvatarIndex = index;
    });
    _saveAvatarSelection(index);

    final avatarAssets = [
      'assets/baguette.png',
      'assets/sushi.png',
      'assets/arly.png',
      'assets/tagin.png',
    ];
    avatarNotifier.value = avatarAssets[index];
  }

  @override
  Widget build(BuildContext context) {
    return Card(
      elevation: 4,
      child: Padding(
        padding: const EdgeInsets.all(16.0),
        child: SingleChildScrollView(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              const Text("Mon compte", style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold)),
              const SizedBox(height: 8),
              Text("Nom d'utilisateur : ${widget.name}"),
              Text("Email : ${widget.email}"),
              const SizedBox(height: 16),
              const Text("Mon avatar", style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold)),
              const SizedBox(height: 16),
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: List.generate(4, (index) {
                  final List<String> avatarAssets = [
                    'assets/tagin.png',
                    'assets/sushi.png',
                    'assets/arly.png',
                    'assets/baguette.png',
                  ];
                  return GestureDetector(
                    onTap: () => _onAvatarTap(index),
                    child: Container(
                      decoration: BoxDecoration(
                        shape: BoxShape.circle,
                        border: Border.all(
                          color: selectedAvatarIndex == index ? AppColors.blackPrimary : Colors.transparent,
                          width: 2,
                        ),
                      ),
                      child: CircleAvatar(
                        backgroundImage: AssetImage(avatarAssets[index]),
                        radius: 25,
                      ),
                    ),
                  );
                }),
              ),
              const SizedBox(height: 12),
              const Text("Mon abonnement", style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold)),
              const SizedBox(height: 8),
              Row(
                mainAxisAlignment: MainAxisAlignment.spaceBetween,
                children: [
                  const Text("Renouvellement automatique"),
                  Switch(
                    value: widget.isRecurringSubscription,
                    onChanged: (value) {}, // Handle locally or ignore
                    activeColor: AppColors.primaryGreen,
                  ),
                ],
              ),
              Text("Type d'abonnement : ${widget.subscriptionType}"),
              Text("Date d'abonnement : ${widget.subscriptionId}"),
              const SizedBox(height: 5),
              Row(
                children: [
                  Expanded(
                    child: OutlinedButton(
                      onPressed: widget.navigateToSubscription,
                      style: OutlinedButton.styleFrom(
                        shape: const RoundedRectangleBorder(
                          borderRadius: BorderRadius.all(Radius.circular(0)),
                        ),
                      ),
                      child: const Row(
                        mainAxisAlignment: MainAxisAlignment.center,
                        children: [
                          Icon(Icons.credit_card),
                          SizedBox(width: 8),
                          Text("Gérer mon abonnement"),
                        ],
                      ),
                    ),
                  ),
                ],
              ),
              const SizedBox(height: 16),
              const Text("Général", style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold)),
              const SizedBox(height: 8),
              ListTile(
                title: const Text("Reprendre le questionnaire", style: TextStyle(color: AppColors.greyPrimary)),
                onTap: widget.onTakeQuestionnaire,
              ),
              ListTile(
                title: const Text("Nous contacter", style: TextStyle(color: AppColors.greyPrimary)),
                onTap: widget.onContact,
              ),
              ListTile(
                title: const Text("Supprimer mon compte", style: TextStyle(color: AppColors.primaryGreen)),
                onTap: widget.onDeleteAccount,
              ),
              ListTile(
                title: const Text("Se déconnecter", style: TextStyle(color: AppColors.primaryGreen)),
                onTap: widget.onLogout,
              ),
            ],
          ),
        ),
      ),
    );
  }
}
