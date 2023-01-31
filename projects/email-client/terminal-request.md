curl -H 'api-key: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' \ 
-X POST -d '{ 
"name":"Campaign sent via the API",
"subject":"TestEmail",
"sender": { "name": "From name", "email":"thakur.cs.tarun@gmail.com" },
"type": "classic",

"htmlContent": "Congratulations! You successfully sent this example campaign via the Sendinblue API.",

"recipients": { "mnirmalkar745@gmail.com" },
# Schedule the sending in one hour
"scheduledAt": "2023-01-29 11:45:01",
}'\
'https://api.sendinblue.com/v3/emailCampaigns'

