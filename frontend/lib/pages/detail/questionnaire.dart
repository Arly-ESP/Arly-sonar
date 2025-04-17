import 'package:Arly/core/style.dart';
import 'package:flutter/material.dart';
import 'package:Arly/pages/home/home_page.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:Arly/config.dart';
import 'package:Arly/core/constants/survey_data.dart';

class QuestionnairePage extends StatefulWidget {
  const QuestionnairePage({super.key});

  @override
  State<QuestionnairePage> createState() => _QuestionnairePageState();
}

class _QuestionnairePageState extends State<QuestionnairePage> {
  final PageController _pageController = PageController();
  int _currentPage = 0;
  final int _totalQuestions = 8;





  // Survey data
  final Map<String, dynamic> _surveyData = surveyData;


  // To manage the state of user answers
  final Map<int, dynamic> answers = {};

  // Method to submit answers
Future<void> _submitAnswers() async {
  // Prepare the formatted answers as expected by the server
  final formattedAnswers = {
    "answers": {
      for (var entry in answers.entries)
        (entry.key + 1).toString(): {
          "question": surveyData['survey_questions'][entry.key]['question'],
          "value": entry.value,
        }
    }
  };


  final prefs = await SharedPreferences.getInstance();
  final token = prefs.getString('authToken');

  try {
    final response = await http.post(
      Uri.parse('${HOST}/api/surveys/1/responses'),
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $token',
      },
      body: jsonEncode(formattedAnswers), // Correct payload
    );

    if (response.statusCode == 200 || response.statusCode == 201) {
      // Success: Navigate to HomePage
      Navigator.pushReplacement(
        context,
        MaterialPageRoute(builder: (context) => const HomePage()),
      );
    } else {
      // Show server error message
      print("Server Error: ${response.body}");
      _showErrorSnackBar('Échec de l\'envoi des réponses. Vérifiez vos données.');
    }
  } catch (e) {
    // Connection error
    print("Connection Error: $e");
    _showErrorSnackBar('Erreur de connexion au serveur.');
  }
}

  // Method to show a snackbar with an error message
  void _showErrorSnackBar(String message) {
    ScaffoldMessenger.of(context).showSnackBar(
      SnackBar(content: Text(message)),
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: AppColors.green600,
      body: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          const SizedBox(height: 50),
          const Padding(
            padding: EdgeInsets.all(16.0),
            child: Text(
              'Questionnaire de personnalité',
              style: TextStyle(
                color: AppColors.primaryWhite,
                fontSize: 24,
                fontWeight: FontWeight.bold,
              ),
            ),
          ),
          Padding(
            padding: const EdgeInsets.symmetric(horizontal: 16.0),
            child: Text(
              '${_currentPage + 1}/$_totalQuestions',
              style: const TextStyle(
                color: AppColors.primaryWhite,
                fontSize: 18,
              ),
            ),
          ),
          const SizedBox(height: 30),
          Expanded(
            child: PageView.builder(
              controller: _pageController,
              onPageChanged: (index) {
                setState(() {
                  _currentPage = index;
                });
              },
              itemCount: _totalQuestions,
              itemBuilder: (context, index) {
                return _buildQuestionPage(index);
              },
            ),
          ),
        ],
      ),
    );
  }
Widget _buildQuestionPage(int index) {
  final question = surveyData['survey_questions'][index];
  final questionType = question['question_type'];
  final questionOptions = question['question_options'];

  return Padding(
    padding: const EdgeInsets.all(16.0),
    child: Card(
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(20),
      ),
      child: Padding(
        padding: const EdgeInsets.all(20.0),
        child: SingleChildScrollView( // Make the page scrollable
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Text(
                question['question'],
                style: const TextStyle(
                  fontSize: 20,
                  fontWeight: FontWeight.bold,
                ),
              ),
              const SizedBox(height: 20),
              if (questionType == 'text')
                TextField(
                  onChanged: (value) {
                    answers[index] = value;
                  },
                  decoration: const InputDecoration(
                    border: OutlineInputBorder(),
                    hintText: 'Entrez votre réponse',
                  ),
                )
              else if (questionType == 'radio')
                Column(
                  children: (questionOptions['options'] as List<dynamic>)
                      .map((option) => RadioListTile(
                            title: Text(option),
                            value: option,
                            groupValue: answers[index],
                            onChanged: (value) {
                              setState(() {
                                answers[index] = value;
                              });
                            },
                          ))
                      .toList(),
                )
              else if (questionType == 'checkbox')
                Column(
                  children: (questionOptions['options'] as List<dynamic>)
                      .map((option) => CheckboxListTile(
                            title: Text(option),
                            value: answers[index]?.contains(option) ?? false,
                            onChanged: (value) {
                              setState(() {
                                if (value == true) {
                                  answers[index] = (answers[index] ?? [])..add(option);
                                } else {
                                  answers[index]?.remove(option);
                                }
                              });
                            },
                          ))
                      .toList(),
                )
              else
                const SizedBox.shrink(),
              const SizedBox(height: 20),
            ElevatedButton(
  onPressed: () {
    final currentAnswer = answers[_currentPage];

    // Check if the current question is answered
    if (currentAnswer == null ||
        (currentAnswer is String && currentAnswer.trim().isEmpty) ||
        (currentAnswer is List && currentAnswer.isEmpty)) {
      _showErrorSnackBar('Veuillez répondre à la question avant de continuer.');
    } else {
      if (_currentPage < _totalQuestions - 1) {
        _pageController.nextPage(
          duration: const Duration(milliseconds: 300),
          curve: Curves.easeInOut,
        );
      } else {
        _submitAnswers();
      }
    }
  },
  style: ElevatedButton.styleFrom(
    backgroundColor: AppColors.primaryGreen,
    padding: const EdgeInsets.symmetric(horizontal: 50, vertical: 15),
    shape: RoundedRectangleBorder(
      borderRadius: BorderRadius.circular(10),
    ),
  ),
  child: Text(
    _currentPage < _totalQuestions - 1 ? 'Suivant' : 'Terminer',
    style: const TextStyle(fontSize: 18),
  ),
)

            ],
          ),
        ),
      ),
    ),
  );
}
}
