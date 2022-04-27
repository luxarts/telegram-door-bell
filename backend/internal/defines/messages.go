package defines

const (
	MessageStart = "Hola!\nPara recibir un mensaje debes enviar una petición POST a %s usando el header `Authorization` con el token:\n`%s`\nMantén tu token *seguro* y no lo compartas con *nadie*.\nPuedes generar un nuevo token en cualquier momento con el comando /token."
	MessageHelp  = "Generá un token con /token y luego haz una petición POST a %s usando el header `Authorization` con tu token para recibir el mensaje."
	MessageToken = "Tu nuevo token es:\n`%s`\nMantén tu token *seguro* y no lo compartas con *nadie*."
	MessageError = "Oh no! Hubo un error, vuelve a intentarlo nuevamente."
)
