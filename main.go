package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"gopkg.in/gomail.v2"
)

type EmailRequest struct {
	Email      string `json:"email"`
	Username   string `json:"username"`
	VerifyCode string `json:"verifyCode"`
}

// SendVerificationEmail sends a verification email
func SendVerificationEmail(email string, username string, verifyCode string) (map[string]interface{}, error) {
	// SMTP server configuration
	smtpServer := "smtp.gmail.com"              // Replace with your SMTP server
	smtpPort := 587                             // Replace with your SMTP port
	smtpUser := "ashutosh.linkedin14@gmail.com" // Replace with your SMTP username
	smtpPassword := "gyeiiqhxrzwjnqtc"          // Replace with your SMTP password

	// Create a new email message
	m := gomail.NewMessage()
	m.SetHeader("From", "onboarding@mystrify.com")
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Verification Code | Mysterify")
	m.SetBody("text/html", VerificationEmail(username, verifyCode))

	// Dial and send the email
	d := gomail.NewDialer(smtpServer, smtpPort, smtpUser, smtpPassword)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("Error Sending Verification Email: %v", err)
		return map[string]interface{}{
			"success": false,
			"message": "Failed to send Verification Email",
		}, err
	}

	return map[string]interface{}{
		"success": true,
		"message": "Verification Email Sent Successfully",
	}, nil
}

//handler starts

func handleSendEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	response, err := SendVerificationEmail(req.Email, req.Username, req.VerifyCode)
	if err != nil {
		http.Error(w, response["message"].(string), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// Verification Email generates the email body for verification
func VerificationEmail(username string, verifyCode string) string {
	return fmt.Sprintf(`
        <html lang="en">
        <head>
            <meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Verification Card</title>
            <style>
                body {
                    margin: 0;
                    padding: 0;
                    display: flex;
                    justify-content: center;
                    align-items: center;
                    height: 100vh;
                    background: #333; /* Dark background color */
                }
                .card-container {
                    max-width: 640px;
                    width: 100%%;
                    margin: 0 auto;
                    padding: 16px;
                    border-radius: 12px;
                    border: 1px solid rgba(255, 255, 255, 0.2); /* Light border for glassmorphism */
                    background: rgba(255, 255, 255, 0.2); /* Semi-transparent background */
                    backdrop-filter: blur(15px); /* Frosted glass effect */
                    -webkit-backdrop-filter: blur(15px); /* Safari compatibility */
                    color: #fff;
                    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.3); /* Shadow for depth */
                    position: relative;
                    display: flex;
                    align-items: center;
                }
                .card-container .logo {
                    position: absolute;
                    top: 16px; /* Adjust top padding as needed */
                    right: 16px; /* Adjust right padding as needed */
                    width: 60px;
                    height: 40px;
                    object-fit: contain; /* Ensures the logo scales correctly */
                }
                .card-content {
                    width: 100%%;
                    text-align: center;
                }
                .card-container table {
                    width: 100%%;
                    border-collapse: collapse;
                }
                .card-container td {
                    padding-right: 16px;
                }
                .card-container p {
                    margin: 0;
                }
                .card-container .header {
                    color: #FFD700;
                    font-family: 'Arial', sans-serif;
                    font-size: 18px;
                    font-weight: 700;
                    line-height: 24px;
                }
                .card-container .details {
                    color: #ddd;
                    font-family: 'Arial', sans-serif;
                    font-size: 14px;
                    font-weight: 400;
                    line-height: 20px;
                    margin-top: 4px;
                }
                .card-container .verification-code {
                    font-size: 24px;
                    font-weight: 700;
                    color: #FFD700;
                    text-shadow: 0 0 5px #FFD700, 0 0 10px #FFD700;
                    text-align: center;
                    margin: 20px 0;
                }
                .card-container .cardowner {
                    font-size: 16px;
                    color: #aaa;
                    text-align: center;
                }
            </style>
        </head>
        <body>
            <div class="card-container">
                <div class="card-content">
                    <p class="header">Verification Card</p>
                    <p class="verification-code">%s</p>
                    <p class="cardowner">%s</p>
                </div>
                <img src="https://i.imgur.com/QJL5bL2.png" class="logo" alt="Logo">
            </div>
        </body>
        </html>
    `, verifyCode, username)
}

// func VerificationEmail(username string, verifyCode string) string {
// 	return fmt.Sprintf(`
//         <html lang="en">
//         <head>
//             <meta charset="UTF-8">
//             <meta name="viewport" content="width=device-width, initial-scale=1.0">
//             <title>Verification Code | MYSTERIFY </title>
//             <style>
//                 body {
//                     margin: 0;
//                     padding: 0;
//                     display: flex;
//                     justify-content: center;
//                     align-items: center;
//                     height: 100vh;
//                     background: #333; /* Dark background color */
//                 }
//                 .card-container {
//                     max-width: 640px;
//                     margin: 0 auto;
//                     box-sizing: border-box;
//                     padding: 16px;
//                     border-radius: 12px;
//                     border: 1px solid rgba(255, 255, 255, 0.2); /* Light border for glassmorphism */
//                     background: rgba(255, 255, 255, 0.1); /* Semi-transparent background */
//                     backdrop-filter: blur(10px); /* Frosted glass effect */
//                     -webkit-backdrop-filter: blur(10px); /* Safari compatibility */
//                     color: #fff;
//                 }
//                 .card-container table {
//                     width: 100%%;
//                     border-collapse: collapse;
//                 }
//                 .card-container td {
//                     padding-right: 16px;
//                 }
//                 .card-container p {
//                     margin: 0;
//                 }
//                 .card-container .header {
//                     color: #7449c4;
//                     font-family: 'Arial', sans-serif;
//                     font-size: 18px;
//                     font-weight: 700;
//                     line-height: 24px;
//                 }
//                 .card-container .details {
//                     color: #ddd;
//                     font-family: 'Arial', sans-serif;
//                     font-size: 14px;
//                     font-weight: 400;
//                     line-height: 20px;
//                     margin-top: 4px;
//                 }
//                 .card-container .verification-code {
//                     font-size: 24px;
//                     font-weight: 700;
//                     color: #FFD700;
//                     text-shadow: 0 0 5px #FFD700, 0 0 10px #FFD700;
//                     text-align: center;
//                     margin: 20px 0;
//                 }
//                 .card-container .cardowner {
//                     font-size: 16px;
//                     color: #aaa;
//                     text-align: center;
//                 }
//             </style>
//         </head>
//         <body>
//             <div class="card-container">
//                 <table>
//                     <tbody>
//                         <tr>
//                             <td>
//                                 <p class="header">Verification Code - MYSTERIFY</p>
//                                 <p class="verification-code">%s</p>
//                                 <p class="cardowner">%s</p>
//                             </td>
//                         </tr>
//                     </tbody>
//                 </table>
//             </div>
//         </body>
//         </html>
//     `, verifyCode, username)
// }

// func VerificationEmail(username string, verifyCode string) string {
// 	return fmt.Sprintf(`
//         <html>
//             <body>
//                 <p>Hello %s,</p>
//                 <p>Your verification code is: <strong>%s</strong></p>
//                 <p>Thank you!</p>
//             </body>
//         </html>
//     `, username, verifyCode)
// }

func main() {
	// Example usage
	//err := godotenv.Load()

	// if err != nil {
	// 	log.Fatal("Error Loading Api key")
	// }

	http.HandleFunc("/send-verification-email", handleSendEmail)
	port := "8010"

	log.Printf("Server starting on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
