package templates

import (
	"bytes"
	"html/template"
)

type EmailVerificationCode struct {
	VerificationCode string
}

func GenerateVerificationCodeEmail(code string) (string, error) {
	const verificationCodeTemplate = `
	<!DOCTYPE html>
	<html>
	<head>
	    <meta charset="UTF-8">
	    <meta name="viewport" content="width=device-width, initial-scale=1.0">
	    <title>Vérifiez Votre Email ✅</title>
	    <style>
	        body {
	            margin: 0;
	            padding: 0;
	            font-family: 'Helvetica Neue', Arial, sans-serif;
	            background-color: #f4f4f9;
	            color: #333333;
	        }
	        .email-container {
	            max-width: 600px;
	            margin: 30px auto;
	            background: #ffffff;
	            border-radius: 10px;
	            box-shadow: 0 5px 15px rgba(0, 0, 0, 0.1);
	            overflow: hidden;
	        }
	        .email-header {
	            background-color: #2D964B;
	            padding: 25px;
	            text-align: center;
	            color: #ffffff;
	            font-size: 26px;
	            font-weight: 600;
	        }
	        .email-body {
	            padding: 30px 20px;
	            text-align: center;
	        }
	        .email-body p {
	            font-size: 16px;
	            line-height: 1.8;
	            margin: 15px 0;
	        }
	        .verification-code {
	            display: inline-block;
	            font-size: 28px;
	            font-weight: 700;
	            color: #2D964B;
	            background: #f1f8ff;
	            padding: 12px 25px;
	            border-radius: 8px;
	            margin: 25px 0;
	            border: 1px solid #2D964B;
	        }
	        .email-footer {
	            padding: 20px;
	            text-align: center;
	            font-size: 14px;
	            color: #666666;
	            border-top: 1px solid #dddddd;
	        }
	        .email-footer a {
	            color: #2D964B;
	            text-decoration: none;
	        }
	    </style>
	</head>
	<body>
	    <div class="email-container">
	        <div class="email-header">
	            Vérifiez Votre Email ✅
	        </div>
	        <div class="email-body">
	            <p>Bonjour 👋,</p>
	            <p>Bienvenue chez <strong>Arly</strong> ! 🎉 Pour continuer, veuillez vérifier votre adresse e-mail en utilisant le code ci-dessous :</p>
	            <div class="verification-code">{{.VerificationCode}}</div>
	            <p>Ce code est valide pendant <strong>10 minutes</strong>. ⏳</p>
	            <p>Si vous n'avez pas demandé cela, ignorez simplement cet e-mail. Si vous avez des questions, n'hésitez pas à contacter notre équipe de support.</p>
	        </div>
	        <div class="email-footer">
	            <p>&copy; 2024 Arly. Tous droits réservés.</p>
	            <p>Besoin d'aide ? Contactez-nous à <a href="mailto:support@arly.com">support@arly.com</a></p>
	        </div>
	    </div>
	</body>
	</html>
	`

	tmpl, _ := template.New("verificationCode").Parse(verificationCodeTemplate)

	data := EmailVerificationCode{
		VerificationCode: code,
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
