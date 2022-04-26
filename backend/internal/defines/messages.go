package defines

const (
	MessageStart = "Hola!\nPara recibir un mensaje primero debes generar un /token. Luego envía una petición POST a %s usando el header `Authorization` con el token generado."
	MessageHelp  = "Generá un token con /token y luego haz una petición POST a %s usando el header `Authorization` con tu token para recibir el mensaje."
	MessageToken = "Tu nuevo token es:\n`%s`\nMantén tu token *seguro* y no lo compartas con *nadie*."
)
