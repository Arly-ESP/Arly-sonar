import 'package:features_tour/features_tour.dart';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'package:speech_to_text/speech_to_text.dart';
import 'package:shared_preferences/shared_preferences.dart';

import '../widgets/chat_app_bar.dart';
import '../widgets/chat_bubble.dart';
import '../widgets/message_input.dart';
import 'package:Arly/config.dart';
import 'package:Arly/core/style.dart';

class ChatPage extends StatefulWidget {
  const ChatPage({super.key});

  @override
  State<ChatPage> createState() => _ChatPageState();
}

class _ChatPageState extends State<ChatPage> {
  final TextEditingController _messageController = TextEditingController();
  final List<Map<String, String>> _messages = [];

  final SpeechToText _speech = SpeechToText();
  final tourController = FeaturesTourController('ChatPage');

  bool _isLoading = false;
  bool _isListening = false;
  bool _isSpeechInitialized = false;
  String _selectedTheme = "famille";

  @override
  void initState() {
    super.initState();
    _initializeSpeech();
    tourController.start(context);
  }

  Future<void> _initializeSpeech() async {
    _isSpeechInitialized = await _speech.initialize(
      onStatus: (status) => debugPrint('Speech status: $status'),
      onError: (error) => debugPrint('Speech error: $error'),
      debugLogging: true,
    );

    if (!_isSpeechInitialized) {
      debugPrint("Speech recognition not available");
    }
  }

  void _startListening() async {
    if (!_isSpeechInitialized) return;

    await _speech.listen(
      localeId: 'fr_FR',
      onResult: (result) {
        if (mounted) {
          setState(() {
            _messageController.text = result.recognizedWords;
          });
        }
      },
    );

    if (mounted) {
      setState(() => _isListening = true);
    }
  }

  void _stopListening() async {
    await _speech.stop();
    if (mounted) {
      setState(() => _isListening = false);
    }
  }

  void _addUserMessage(String text) {
    _messages.add({'sender': 'user', 'text': text});
  }

  void _addAIMessage(String text) {
    _messages.add({'sender': 'ai', 'text': text});
  }

  Future<void> _sendMessage() async {
    final message = _messageController.text.trim();
    if (message.isEmpty) return;

    setState(() {
      _addUserMessage(message);
      _messageController.clear();
      _isLoading = true;
    });

    try {
      final prefs = await SharedPreferences.getInstance();
      final token = prefs.getString('authToken');

      final response = await http.post(
        Uri.parse('$HOST/api/chat'),
        headers: {
          'Content-Type': 'application/json',
          if (token != null) 'Authorization': 'Bearer $token',
        },
        body: jsonEncode({'message': message}),
      );

      final decodedBody = utf8.decode(response.bodyBytes);

      if (response.statusCode == 200) {
        final responseData = jsonDecode(decodedBody);
        final aiResponse = responseData['response'] ?? 'Pas de réponse de l’IA.';
        if (mounted) {
          setState(() => _addAIMessage(aiResponse));
        }
      } else {
        if (mounted) {
          setState(() => _addAIMessage("Erreur: Impossible d’obtenir une réponse de l’IA."));
        }
      }
    } catch (e) {
      if (mounted) {
        setState(() => _addAIMessage("Erreur réseau ou serveur: $e"));
      }
    } finally {
      if (mounted) {
        setState(() => _isLoading = false);
      }
    }
  }

  Widget _buildMessageList() {
    return Expanded(
      child: ListView.builder(
        reverse: true,
        itemCount: _messages.length,
        itemBuilder: (context, index) {
          final message = _messages[_messages.length - 1 - index];
          return ChatBubble(
            isUser: message['sender'] == 'user',
            text: message['text'] ?? '',
          );
        },
      ),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: PreferredSize(
        preferredSize: const Size.fromHeight(kToolbarHeight),
        child: FeaturesTour(
          controller: tourController,
          index: 1,
          introduce: const Text(
            'Bienvenue dans la page de chat ! Cette section permet de sélectionner le thème de la conversation. Il te dit également combien de temps tu peux passer pour compléter ta quête !',
            style: TextStyle(fontSize: 24, color: AppColors.primaryWhite),
          ),
          child: ChatAppBar(
            selectedTheme: _selectedTheme,
            onThemeChanged: (value) {
              setState(() => _selectedTheme = value);
            },
          ),
        ),
      ),
      body: Column(
        children: [
          _buildMessageList(),
          if (_isLoading)
            const Padding(
              padding: EdgeInsets.all(8.0),
              child: CircularProgressIndicator(),
            ),
          FeaturesTour(
            controller: tourController,
            index: 2,
            introduce: const Text(
              'Ici, tu peux écrire tes messages à Arly pour discuter autour du thème que tu as choisi. Tu peux aussi lui parler directement en appuyant sur le micro !',
              style: TextStyle(fontSize: 24, color: AppColors.primaryWhite),
            ),
            child: Padding(
              padding: const EdgeInsets.only(bottom: 18.0),
              child: MessageInput(
                controller: _messageController,
                isListening: _isListening,
                isLoading: _isLoading,
                onMicTap: () => _isListening ? _stopListening() : _startListening(),
                onSendTap: _sendMessage,
              ),
            ),
          ),
        ],
      ),
    );
  }
}
