const Map<String, dynamic> surveyData = {
  "survey_name": "Onboarding : Personnalité et Préférences",
  "survey_slug": "onboarding-personnalite-preferences",
  "survey_description":
  "Un questionnaire pour comprendre votre personnalité, vos préférences et vos centres d'intérêt.",
  "survey_questions": [
    {
      "question":
      "Quel est ton prénom, ton sexe (homme, femme, autre) et ton âge ?",
      "question_type": "text",
      "question_options": {},
      "order": 1
    },
    {
      "question": "Parmi les traits suivants, lequel te correspond le plus ?",
      "question_type": "radio",
      "question_options": {
        "options": ["Rebelle", "Organisé(e)", "Calme", "Aventureux(se)", "Créatif(ve)"]
      },
      "order": 2
    },
    {
      "question": "Aimes-tu prendre des décisions rapidement ?",
      "question_type": "radio",
      "question_options": {
        "options": ["Oui", "Non"]
      },
      "order": 3
    },
    {
      "question": "Es-tu à l’aise avec les imprévus ?",
      "question_type": "radio",
      "question_options": {
        "options": ["Oui", "Non"]
      },
      "order": 4
    },
    {
      "question": "Préfères-tu travailler seul(e) plutôt qu’en équipe ?",
      "question_type": "radio",
      "question_options": {
        "options": ["Oui", "Non"]
      },
      "order": 5
    },
    {
      "question":
      "Es-tu du genre à finir ce que tu as commencé, même si c’est difficile ?",
      "question_type": "radio",
      "question_options": {
        "options": ["Oui", "Non"]
      },
      "order": 6
    },
    {
      "question":
      "Prends-tu le temps de te déconnecter de ton téléphone chaque jour ?",
      "question_type": "radio",
      "question_options": {
        "options": ["Oui", "Non"]
      },
      "order": 7
    },
    {
      "question": "Quels sont tes centres d’intérêt parmi les suivants ?",
      "question_type": "checkbox",
      "question_options": {
        "options": [
          "Technologie",
          "Nature",
          "Sport",
          "Lecture",
          "Cinéma",
          "Musique",
          "Cuisine",
          "Voyage",
          "Gaming",
          "Mode",
          "Art",
          "Sciences",
          "Philosophie",
          "Activisme",
          "Méditation"
        ]
      },
      "order": 8
    }
  ]
};
